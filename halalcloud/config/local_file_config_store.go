package config

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

type localFileConfigStore struct {
	filePath string
	locker   sync.Mutex
	dataMap  map[string]string
}

// ClearConfigs implements ConfigStore.
func (l *localFileConfigStore) ClearConfigs() error {
	l.locker.Lock()
	defer l.locker.Unlock()
	os.RemoveAll(l.filePath)
	return nil
}

// DeleteConfig implements ConfigStore.
func (l *localFileConfigStore) DeleteConfig(key string) error {
	l.locker.Lock()
	defer l.locker.Unlock()
	delete(l.dataMap, key)
	// sync to file
	return l.syncToFile()
}

func (l *localFileConfigStore) syncToFile() error {
	jsonData, err := json.Marshal(l.dataMap)
	if err != nil {
		return err
	}
	err = os.WriteFile(l.filePath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

// GetAccessToken implements ConfigStore.
func (l *localFileConfigStore) GetAccessToken() (string, error) {
	l.locker.Lock()
	defer l.locker.Unlock()
	return l.dataMap["access_token"], nil
}

// GetConfig implements ConfigStore.
func (l *localFileConfigStore) GetConfig(key string) (string, error) {
	l.locker.Lock()
	defer l.locker.Unlock()
	return l.dataMap[key], nil
}

// GetRefreshToken implements ConfigStore.
func (l *localFileConfigStore) GetRefreshToken() (string, error) {
	l.locker.Lock()
	defer l.locker.Unlock()
	return l.dataMap["refresh_token"], nil
}

// ListConfigs implements ConfigStore.
func (l *localFileConfigStore) ListConfigs() (map[string]string, error) {
	l.locker.Lock()
	defer l.locker.Unlock()
	return l.dataMap, nil
}

// SetAccessToken implements ConfigStore.
func (l *localFileConfigStore) SetAccessToken(token string) error {
	l.locker.Lock()
	defer l.locker.Unlock()
	l.dataMap["access_token"] = token
	return l.syncToFile()
}

// SetConfig implements ConfigStore.
func (l *localFileConfigStore) SetConfig(key string, value string) error {
	l.locker.Lock()
	defer l.locker.Unlock()
	l.dataMap[key] = value
	return l.syncToFile()
}

// SetRefreshToken implements ConfigStore.
func (l *localFileConfigStore) SetRefreshToken(token string) error {
	l.locker.Lock()
	defer l.locker.Unlock()
	l.dataMap["refresh_token"] = token
	return l.syncToFile()
}

// SetToken implements ConfigStore.
func (l *localFileConfigStore) SetToken(accessToken string, refreshToken string, expiresIn int64) error {
	l.locker.Lock()
	defer l.locker.Unlock()
	l.dataMap["access_token"] = accessToken
	l.dataMap["refresh_token"] = refreshToken
	l.dataMap["expires_in"] = fmt.Sprintf("%d", expiresIn)
	return l.syncToFile()
}

func NewLocalFileConfigStore(filePath string) ConfigStore {
	cf := &localFileConfigStore{
		filePath: filePath,
		locker:   sync.Mutex{},
	}
	cf.dataMap = make(map[string]string)
	// load from file
	if _, err := os.Stat(filePath); err == nil {
		jsonData, err := os.ReadFile(filePath)
		if err != nil {
			return nil
		}
		if err := json.Unmarshal(jsonData, &cf.dataMap); err != nil {
			return nil
		}
	}
	return cf
}
