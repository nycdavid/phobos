package cmd

import (
	"testing"
)

func Test_NewConfigFile(t *testing.T) {
	configFile := NewConfigFile("test/db/config.json")

	h := configFile.
		Environments["development"]["host"].(string)

	expected := "localhost"
	if h != expected {
		t.Errorf("Expected url to be %s, got %s", expected, h)
	}
}
