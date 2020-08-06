package cmd

import (
	"os"
	"path"
	"testing"
)

func Test_NewConfigFile(t *testing.T) {
	root := os.Getenv("PROJECT_ROOT")
	configPath := path.Join(root, "db", "config.json")

	configFile := NewConfigFile(configPath)

	developmentUrl := configFile.
		Environments["development"]["url"].(string)

	expected := "postgres://localhost/postgres?sslmode=disable"
	if developmentUrl != expected {
		t.Errorf("Expected url to be %s, got %s", expected, developmentUrl)
	}
}
