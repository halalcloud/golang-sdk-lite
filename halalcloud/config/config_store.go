package config

type ConfigStore interface {
	// GetConfig retrieves the configuration for a given key.
	GetConfig(key string) (string, error)
	// SetConfig sets the configuration for a given key.
	SetConfig(key, value string) error
	// DeleteConfig deletes the configuration for a given key.
	DeleteConfig(key string) error
	// ListConfigs lists all configurations.
	ListConfigs() (map[string]string, error)
	// ClearConfigs clears all configurations.
	ClearConfigs() error
}
