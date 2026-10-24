package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestRunVersion(t *testing.T) {
	var buf bytes.Buffer
	// Run writes to stdout via fmt; capture not needed for error check
	if err := Run([]string{"version"}); err != nil {
		t.Fatal(err)
	}
	_ = buf
	_ = strings.Contains
}

func TestRunUnknown(t *testing.T) {
	if err := Run([]string{"nope"}); err == nil {
		t.Fatal("expected error")
	}
}

func TestRunNoArgs(t *testing.T) {
	if err := Run(nil); err == nil {
		t.Fatal("expected error")
	}
}
