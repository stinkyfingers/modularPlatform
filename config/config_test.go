package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestGetConfig(t *testing.T) {
	configPath := filepath.Join("test_data", "config.conf")
	defer os.RemoveAll(filepath.Dir(configPath))

	err := os.Setenv(CONFIG_LOCATION, configPath)
	if err != nil {
		t.Error(err)
	}
	v := struct {
		ID   int    `yaml:"id"`
		Name string `yaml:"name"`
	}{}
	mockConfig(configPath, `id: 1
name: foo`)

	err = GetConfig(&v)
	if err != nil {
		t.Error(err)
	}
	if v.ID != 1 {
		t.Error("expected ID to be 1")
	}
	if v.Name != "foo" {
		t.Error("expected Name to be foo")
	}
}

// UTIL

// mockConfig populates a config at path with data for testing
func mockConfig(path, data string) error {
	os.Mkdir(filepath.Dir(path), os.ModePerm)
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	_, err = f.WriteString(data)
	return err
}
