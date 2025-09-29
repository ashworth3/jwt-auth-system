package utils

import (
	"testing"
)

func TestGenerateAndParseJWT(t *testing.T) {
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
