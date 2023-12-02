package mypackage1

import (
	"strings"
	"testing"
)

func TestMyVar(t *testing.T) {
	if !strings.Contains(MyVar.Name, "mypackage1") {
		t.Fatalf("caller file should contain 'mypackage1', but was '%s'", MyVar.Name)
	}
}
