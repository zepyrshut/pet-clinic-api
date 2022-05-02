package middleware

import "testing"

func TestLocalize(t *testing.T) {
	localize := Localize()

	if localize == nil {
		t.Error("Localize() should not be nil")
	}
}
