package middleware

import (
	"testing"

	"github.com/zepyrshut/pet-clinic/internal/config"
)

func TestNewMiddleware(t *testing.T) {
	a := &config.Application{}
	NewMiddleware(a)

	if app == nil {
		t.Error("app should not be nil")
	}
}

func TestSessions(t *testing.T) {
	sessions := Sessions("test")

	if sessions == nil {
		t.Error("Sessions() should not be nil")
	}
}

func TestCORSMiddleware(t *testing.T) {
	corsMiddleware := CORSMiddleware()

	if corsMiddleware == nil {
		t.Error("CORSMiddleware() should not be nil")
	}
}

func TestLocalize(t *testing.T) {
	localize := Localize()

	if localize == nil {
		t.Error("Localize() should not be nil")
	}
}
