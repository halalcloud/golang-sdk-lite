package apiclient

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// TokenResponse 表示刷新令牌响应
type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"` // 过期时间(秒)
}

// TokenManager 负责管理访问令牌
type TokenManager interface {
	// GetToken 返回当前有效的访问令牌
	GetToken() (string, error)

	// RefreshToken 刷新访问令牌
	RefreshToken() (string, error)

	// SetTokens 设置新的访问令牌和刷新令牌
	SetTokens(accessToken, refreshToken string, expiresIn int64)
}

// DefaultTokenManager 提供令牌管理的默认实现
type DefaultTokenManager struct {
	client       *Client                                                                // API客户端引用
	accessToken  string                                                                 // 当前访问令牌
	refreshToken string                                                                 // 用于刷新的令牌
	expiry       time.Time                                                              // 令牌过期时间
	mutex        sync.RWMutex                                                           // 保证线程安全
	refreshFunc  func(ctx context.Context, refreshToken string) (*TokenResponse, error) // 刷新令牌的函数
}

// NewDefaultTokenManager 创建一个新的默认令牌管理器
func NewDefaultTokenManager(client *Client, refreshFunc func(ctx context.Context, refreshToken string) (*TokenResponse, error)) *DefaultTokenManager {
	return &DefaultTokenManager{
		client:      client,
		refreshFunc: refreshFunc,
		mutex:       sync.RWMutex{},
	}
}

// GetToken 获取有效的访问令牌，如果令牌已过期则尝试刷新
func (tm *DefaultTokenManager) GetToken() (string, error) {
	tm.mutex.RLock()

	// 检查是否有token且未过期
	if tm.accessToken != "" && time.Now().Before(tm.expiry) {
		token := tm.accessToken
		tm.mutex.RUnlock()
		return token, nil
	}

	tm.mutex.RUnlock()

	// 如果token为空或已过期，则刷新
	return tm.RefreshToken()
}

// RefreshToken 刷新访问令牌
func (tm *DefaultTokenManager) RefreshToken() (string, error) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	// 再次检查，防止在获取锁过程中已被其他goroutine刷新
	if tm.accessToken != "" && time.Now().Before(tm.expiry) {
		return tm.accessToken, nil
	}

	if tm.refreshToken == "" {
		return "", errors.New("no refresh token available")
	}

	// 调用刷新令牌的函数
	tokenResp, err := tm.refreshFunc(context.Background(), tm.refreshToken)
	if err != nil {
		return "", fmt.Errorf("failed to refresh token: %w", err)
	}

	// 更新令牌信息
	tm.accessToken = tokenResp.AccessToken
	tm.refreshToken = tokenResp.RefreshToken
	tm.expiry = time.Now().Add(time.Duration(tokenResp.ExpiresIn) * time.Second)

	return tm.accessToken, nil
}

// SetTokens 设置新的访问令牌和刷新令牌
func (tm *DefaultTokenManager) SetTokens(accessToken, refreshToken string, expiresIn int64) {
	tm.mutex.Lock()
	defer tm.mutex.Unlock()

	tm.accessToken = accessToken
	tm.refreshToken = refreshToken
	tm.expiry = time.Now().Add(time.Duration(expiresIn) * time.Second)
}
