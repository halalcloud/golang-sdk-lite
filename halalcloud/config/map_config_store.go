package config

import "sync"

type MapConfigStore struct {
	configs sync.Map
}

// GetAccessToken implements ConfigStore.
func (m *MapConfigStore) GetAccessToken() (string, error) {
	value, exists := m.configs.Load("access_token")
	if !exists {
		return "", nil // 如果不存在，返回空字符串
	}
	return value.(string), nil // 返回配置项的值
}

// GetRefreshToken implements ConfigStore.
func (m *MapConfigStore) GetRefreshToken() (string, error) {
	value, exists := m.configs.Load("refresh_token")
	if !exists {
		return "", nil // 如果不存在，返回空字符串
	}
	return value.(string), nil // 返回配置项的值
}

// SetAccessToken implements ConfigStore.
func (m *MapConfigStore) SetAccessToken(token string) error {
	m.configs.Store("access_token", token)
	return nil
}

// SetRefreshToken implements ConfigStore.
func (m *MapConfigStore) SetRefreshToken(token string) error {
	m.configs.Store("refresh_token", token)
	return nil
}

// SetToken implements ConfigStore.
func (m *MapConfigStore) SetToken(accessToken string, refreshToken string, expiresIn int64) error {
	m.configs.Store("access_token", accessToken)
	m.configs.Store("refresh_token", refreshToken)
	m.configs.Store("expires_in", expiresIn)
	return nil
}

// ClearConfigs implements ConfigStore.
func (m *MapConfigStore) ClearConfigs() error {
	m.configs = sync.Map{} // 清空map
	return nil
}

// DeleteConfig implements ConfigStore.
func (m *MapConfigStore) DeleteConfig(key string) error {
	_, exists := m.configs.Load(key)
	if !exists {
		return nil // 如果不存在，直接返回
	}
	m.configs.Delete(key) // 删除指定的配置项
	return nil
}

// GetConfig implements ConfigStore.
func (m *MapConfigStore) GetConfig(key string) (string, error) {
	value, exists := m.configs.Load(key)
	if !exists {
		return "", nil // 如果不存在，返回空字符串
	}
	return value.(string), nil // 返回配置项的值
}

// ListConfigs implements ConfigStore.
func (m *MapConfigStore) ListConfigs() (map[string]string, error) {
	configs := make(map[string]string)
	m.configs.Range(func(key, value interface{}) bool {
		configs[key.(string)] = value.(string) // 将每个配置项添加到map中
		return true                            // 继续遍历
	})
	return configs, nil // 返回所有配置项
}

// SetConfig implements ConfigStore.
func (m *MapConfigStore) SetConfig(key string, value string) error {
	m.configs.Store(key, value) // 使用Store方法设置或更新配置项
	return nil                  // 成功设置配置项后返回nil
}

func NewMapConfigStore() ConfigStore {
	return &MapConfigStore{
		configs: sync.Map{},
	}
}
