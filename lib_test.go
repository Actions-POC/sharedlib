package sharedlib_test

import (
	"github.com/Actions-POC/sharedlib"
	"testing"
)

func TestName(t *testing.T) {
	const expectedName = "sharedlib"
	if name:=sharedlib.Name(); name != expectedName {
		t.Errorf("Expected name '%s' not found, instead: %s",
			expectedName, name)
	}
}


