package main

import (
	"testing"
)

func TestLanguage(t *testing.T) {
	got := language("Flutter")
	want := "My fave programming language is Flutter"
	if got != want {
		t.Errorf("error - want: %s but got: %s", want, got)
	}
}
