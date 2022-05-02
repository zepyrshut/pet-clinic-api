package main

import "testing"

func TestTun(t *testing.T) {
	_, err := run()
	if err != nil {
		t.Error("failed run")
	}
}
