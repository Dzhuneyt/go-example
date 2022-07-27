package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("J")
	want := "Hello, J"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
