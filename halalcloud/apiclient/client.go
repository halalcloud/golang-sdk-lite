package apiclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/halalcloud/golang-sdk-lite/halalcloud/config"
	"github.com/halalcloud/golang-sdk-lite/halalcloud/signer"
)

// Client 表示API客户端
type Client struct {
	Host       string
	HTTPClient *http.Client
	// Signer       Signer
	AccessToken  string
	ClientID     string
	ClientSecret string
	// TokenManager TokenManager // 令牌管理器
	configStore config.ConfigStore
}

// ClientOption 是一个函数类型，用于配置Client
type ClientOption func(*Client)

// NewClient 创建一个新的API客户端
func NewClient(httpClient *http.Client, host string, secretID, secretKey string, configStore config.ConfigStore, options ...ClientOption) *Client {
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 10 * time.Second,
		}
	}
	client := &Client{
		Host:         host,
		HTTPClient:   httpClient,
		ClientSecret: secretKey,
		ClientID:     secretID,
		configStore:  configStore,
	}

	// 应用所有选项
	for _, option := range options {
		option(client)
	}

	token, err := configStore.GetAccessToken()
	if err == nil && len(token) > 0 {
		client.AccessToken = token
	}

	return client
}

// Request 发送请求并处理响应
func (c *Client) Request(ctx context.Context, method string, path string,
	paramsMap map[string]string, body any, result any, isRefreshTokenRequest bool) error {

	// 尝试最多2次请求（首次请求和刷新令牌后的重试）
	for attempt := range 2 {
		// 构建完整URL
		//fullURL, err := url.JoinPath(c.BaseURL, path)
		//if err != nil {
		//	return NewAPIError("invalid_url", err.Error(), 0)
		//}

		// 准备请求体
		var bodyRaw []byte
		headers := make(map[string]string)
		headersToSign := []string{}
		if body != nil {
			jsonData, err := json.Marshal(body)
			if err != nil {
				return NewAPIError("marshal_error", "Failed to marshal request body", 0)
			}
			bodyRaw = jsonData
			headers["content-type"] = "application/json; charset=utf-8"
			headersToSign = append(headersToSign, "content-type")
		}

		signConfig := signer.NewConfig(c.Host, c.ClientID, c.ClientSecret, c.AccessToken, bodyRaw, method, path, paramsMap, headers, headersToSign)
		signerData := signer.NewSigner(signConfig)
		// 创建请求
		req, err := http.NewRequestWithContext(ctx, method, signerData.GetRequestURL(true), bytes.NewReader(bodyRaw))
		if err != nil {
			return NewAPIError("request_creation_error", err.Error(), 0)
		}

		for k, v := range signerData.GetHeaders() {
			req.Header.Set(k, v)
		}

		// 设置请求头

		// 签名请求

		// 发送请求
		resp, err := c.HTTPClient.Do(req)
		if err != nil {
			return NewAPIError("request_error", err.Error(), 0)
		}
		defer resp.Body.Close()

		// 读取响应体
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return NewAPIError("response_read_error", err.Error(), resp.StatusCode)
		}

		// 检查是否是认证错误（通常是401 Unauthorized）
		if resp.StatusCode == http.StatusUnauthorized && attempt == 0 && !c.isUseClientToken() && !isAuthEndpoint(path) && !isRefreshTokenRequest {
			// 尝试刷新令牌
			err := c.RefreshToken(ctx)
			if err != nil {
				// 如果刷新失败，返回原始错误
				apiErr := APIError{}
				if jsonErr := json.Unmarshal(respBody, &apiErr); jsonErr == nil {
					apiErr.StatusCode = resp.StatusCode
					return &apiErr
				}
				return NewAPIError("unauthorized", "Authentication failed and token refresh failed", resp.StatusCode)
			}
			// 刷新成功，进行下一次尝试
			continue
		}

		// 检查状态码
		if resp.StatusCode < 200 || resp.StatusCode >= 300 {
			apiErr := APIError{}
			if err := json.Unmarshal(respBody, &apiErr); err != nil {
				// 如果无法解析错误响应，则创建一个通用错误
				return NewAPIError("unexpected_status", fmt.Sprintf("HTTP %d: %s", resp.StatusCode, string(respBody)), resp.StatusCode)
			}
			apiErr.StatusCode = resp.StatusCode
			return &apiErr
		}

		// 解析结果
		if result != nil {
			if err := json.Unmarshal(respBody, result); err != nil {
				return NewAPIError("unmarshal_error", err.Error(), resp.StatusCode)
			}
		}

		// 请求成功
		return nil
	}

	// 如果多次尝试后仍然失败
	return NewAPIError("max_attempts_reached", "Failed after maximum number of attempts", 0)
}

// isAuthEndpoint 检查是否是认证相关的端点
func isAuthEndpoint(path string) bool {
	// 根据具体API调整，这里假设认证相关的端点包含"auth", "login", "token"等关键词
	authPaths := []string{"/auth", "/login", "/token", "/oauth"}
	for _, authPath := range authPaths {
		if strings.Contains(path, authPath) {
			return true
		}
	}
	return false
}

// Get 发送GET请求
func (c *Client) Get(ctx context.Context, path string, params map[string]string, result any) error {
	return c.Request(ctx, http.MethodGet, path, params, nil, result, false)
}

// Post 发送POST请求
func (c *Client) Post(ctx context.Context, path string, params map[string]string, body any, result any) error {
	return c.Request(ctx, http.MethodPost, path, params, body, result, false)
}

// Put 发送PUT请求
func (c *Client) Put(ctx context.Context, path string, params map[string]string, body any, result any) error {
	return c.Request(ctx, http.MethodPut, path, params, body, result, false)
}

// Delete 发送DELETE请求
func (c *Client) Delete(ctx context.Context, path string, params map[string]string, result any) error {
	return c.Request(ctx, http.MethodDelete, path, params, nil, result, false)
}

func (c *Client) isUseClientToken() bool {
	return false
}

func (c *Client) SetToken(accessToken, refreshToken string, expiresIn int32) {
	c.AccessToken = accessToken

	err := c.configStore.SetToken(accessToken, refreshToken, int64(expiresIn))
	if err != nil {
		log.Printf("Failed to store token: %v", err)
	}
	// 这里可以设置过期时间戳，如果需要的话
}

func (c *Client) RefreshToken(ctx context.Context) error {
	// c.AccessToken = token

	refreshToken, err := c.configStore.GetRefreshToken()
	if err != nil {
		return fmt.Errorf("failed to get refresh token: %w", err)
	}
	if refreshToken == "" {
		return fmt.Errorf("refresh token is empty")
	}
	log.Printf("Refreshing token for old [%s]...", refreshToken)
	res := &TokenResponse{}
	err = c.Request(ctx, http.MethodPost, "/v6/oauth/refresh_token", nil, map[string]string{
		"refresh_token": refreshToken,
		"grant_type":    "refresh_token",
		"client_id":     c.ClientID,
	}, res, true)
	if err != nil {
		log.Printf("Error refreshing token: %v", err)
		return fmt.Errorf("failed to refresh token: %w", err)
	}
	c.AccessToken = res.AccessToken
	err = c.configStore.SetToken(res.AccessToken, res.RefreshToken, int64(res.ExpiresIn))
	if err != nil {
		return fmt.Errorf("failed to store refreshed token: %w", err)
	}
	log.Printf("Token refreshed successfully: Access Token: %s, Refresh Token: %s, Expires In: %d seconds",
		res.AccessToken, res.RefreshToken, res.ExpiresIn)

	return nil
}
