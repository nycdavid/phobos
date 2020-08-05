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

	ctxVal := dbcreateCmd.Context().Value("foo")
	if ctxVal.(string) != "bar" {
		t.Error("Expected bar")
	}
}
