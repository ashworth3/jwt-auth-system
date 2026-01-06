package utils

import (
	"os"
	"testing"
)

func TestGenerateAndParseJWT(t *testing.T) {
	// Set a test JWT_SECRET if not already set
	originalSecret := os.Getenv("JWT_SECRET")
	testSecret := "test-secret-key-must-be-at-least-32-characters-long-for-security"
	
	if originalSecret == "" {
		os.Setenv("JWT_SECRET", testSecret)
		defer os.Unsetenv("JWT_SECRET")
		// Re-initialize jwtKey since init() already ran
		InitJWTKey()
	} else if len(originalSecret) < 32 {
		// If existing secret is too short, use test secret
		os.Setenv("JWT_SECRET", testSecret)
		defer func() {
			os.Setenv("JWT_SECRET", originalSecret)
			InitJWTKey()
		}()
		InitJWTKey()
	}

	token, err := GenerateJWT(1)
	if err != nil {
		t.Fatalf("Failed to generate JWT: %v", err)
	}

	claims, err := ParseJWT(token)
	if err != nil {
		t.Fatalf("Failed to parse JWT: %v", err)
	}

	if claims.Subject != "1" {
		t.Errorf("Expected Subject to be '1', got '%s'", claims.Subject)
	}
}
