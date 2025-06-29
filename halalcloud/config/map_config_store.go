package config

import "sync"

type MapConfigStore struct {
	configs sync.Map
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
