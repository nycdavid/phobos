package cmd

import (
	"path"
	"testing"
)

func Test_NewConfigFile(t *testing.T) {
	configPath := path.Join("test", "db", "config.json")

	configFile := NewConfigFile(configPath)

	developmentUrl := configFile.
		Environments["development"]["url"].(string)

	expected := "postgres://localhost/postgres?sslmode=disable"
	if developmentUrl != expected {
		t.Errorf("Expected url to be %s, got %s", expected, developmentUrl)
	}
}
