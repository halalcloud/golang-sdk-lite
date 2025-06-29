package apiclient

import (
	"net/http"
	"time"
)

// WithTimeout 设置超时时间的选项
func WithTimeout(timeout time.Duration) ClientOption {
	return func(c *Client) {
		c.HTTPClient.Timeout = timeout
	}
}

// WithTokenAuth 添加令牌认证，包括初始令牌和刷新功能
/*
func WithTokenAuth(accessToken, refreshToken string, expiresIn int64, tokenEndpoint string) ClientOption {
	return func(c *Client) {
		refreshFunc := func(ctx context.Context, refreshToken string) (*TokenResponse, error) {
			// 创建令牌刷新请求
			reqBody := map[string]string{
				"grant_type":    "refresh_token",
				"refresh_token": refreshToken,
			}

			// 序列化请求体
			jsonData, err := json.Marshal(reqBody)
			if err != nil {
				return nil, fmt.Errorf("failed to marshal refresh token request: %w", err)
			}

			// 构建完整URL
			fullURL := c.BaseURL
			if !strings.HasSuffix(fullURL, "/") && !strings.HasPrefix(tokenEndpoint, "/") {
				fullURL += "/"
			}
			fullURL += tokenEndpoint

			// 创建请求
			req, err := http.NewRequestWithContext(
				ctx,
				http.MethodPost,
				fullURL,
				strings.NewReader(string(jsonData)),
			)
			if err != nil {
				return nil, fmt.Errorf("failed to create refresh token request: %w", err)
			}

			// 设置请求头
			req.Header.Set("Content-Type", "application/json")
			req.Header.Set("Accept", "application/json")

			// 签名请求（如需要）
			if err := c.Signer.Sign(req, c.AccessToken, c.SecretKey); err != nil {
				return nil, fmt.Errorf("failed to sign refresh token request: %w", err)
			}

			// 发送请求
			resp, err := c.HTTPClient.Do(req)
			if err != nil {
				return nil, fmt.Errorf("failed to execute refresh token request: %w", err)
			}
			defer resp.Body.Close()

			// 检查响应
			if resp.StatusCode != http.StatusOK {
				return nil, fmt.Errorf("failed to refresh token: HTTP %d", resp.StatusCode)
			}

			// 解析响应
			var tokenResp TokenResponse
			if err := json.NewDecoder(resp.Body).Decode(&tokenResp); err != nil {
				return nil, fmt.Errorf("failed to parse refresh token response: %w", err)
			}

			return &tokenResp, nil
		}

		// 创建令牌管理器
		tm := NewDefaultTokenManager(c, refreshFunc)
		tm.SetTokens(accessToken, refreshToken, expiresIn)
		c.TokenManager = tm
	}
}
*/
// WithCustomHTTPClient 设置自定义HTTP客户端
func WithCustomHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *Client) {
		c.HTTPClient = httpClient
	}
}
