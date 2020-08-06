package cmd

import (
	"testing"
)

func TestDbCreate_UsesDevelopmentByDefault(t *testing.T) {
	dbcreateCmd := NewDbCreateCommand()

	fVal := dbcreateCmd.Flag("environment").Value.String()
	if fVal != "development" {
		t.Errorf("Expected %s got %s", "development", fVal)
	}
}

func TestDbCreate_ConfigFileInCtx(t *testing.T) {
	dbcreateCmd := NewDbCreateCommand()
	configFile := dbcreateCmd.Context().Value(ContextKey("configFile"))

	if configFile == nil {
		t.Errorf("Expected non-nil config file")
	}
}
