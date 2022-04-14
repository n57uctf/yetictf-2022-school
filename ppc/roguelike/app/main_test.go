package main

import (
	"fmt"
	"testing"
)

func TestErrorFormat(t *testing.T) {
	got := ErrorFormat(fmt.Errorf("123"))
	want := "\x1b[31m123\x1b[0m"
	if got != want {
		t.Errorf("Got: %q, want: %q", got, want)
	}
}
