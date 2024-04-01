package utils

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("Error hashing password: %v", err)
	}

	if hash == "" {
		t.Error("Hash is empty")
	}

	if !CheckPasswordHash(password, hash) {
		t.Error("Failed to verify hashed password")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "password"
	hash := "$2a$12$zIjvZ.Z43s9YOAcNWsAEc.cTLa.yPzzWRfNXCARloOdKVet9D93L2"

	if !CheckPasswordHash(password, hash) {
		t.Error("Failed to verify hashed password")
	}

	if CheckPasswordHash("wrongpassword", hash) {
		t.Error("Incorrect password verification")
	}

	if CheckPasswordHash(password, "wronghash") {
		t.Error("Incorrect password verification")
	}
}
