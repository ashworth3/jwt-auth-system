package utils

import "testing"

func TestHashAndComparePassword(t *testing.T) {
	password := "password123"

	// Hash the password
	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Failed to hash password: %v", err)
	}

	// Correct password check
	if !CheckPassword(hash, password) {
		t.Errorf("Password verification failed, should have succeeded")
	}

	// Incorrect password check
	if CheckPassword(hash, "wrongpassword") {
		t.Errorf("Password verification passed for wrong password")
	}
}
