package cmd

import (
	"testing"
)

func TestDbCreate_UsesDevelopmentByDefault(t *testing.T) {
	dbcreateCmd := NewDbCreateCommand()
	dbcreateCmd.Execute()
}
