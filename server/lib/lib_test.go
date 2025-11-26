package mylib

import "testing"

func TestLib(t *testing.T) {
	got := getval()
	want := "foo"
	if got != want {
		t.Errorf("getval() = %q, want %q", got, want)
	}
}
