package auth

import (
	"regexp"
	"testing"
)

// HashPassword hashes plaintext password
func TestHashPassword(t *testing.T) {
	password := "pa$$word"
	passwordHashed, err := HashPassword(password)
	if err != nil {
		t.Errorf("expected no errors, but got %s", err)
	}
	match := regexp.MustCompile(`\b` + password + `\b`)
	if match.MatchString(passwordHashed) {
		t.Errorf(`HashPassword("password") = %q, %v, want no match for %#q, nil`, passwordHashed, err, match)
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "pa$$w0rd"
	passwordHashed, err := HashPassword(password)
	if err != nil {
		t.Errorf("expected no errors, but got %s", err)
	}
	match, err := CheckPasswordHash(password, passwordHashed)
	if err != nil {
		t.Errorf("expected no errors, but got %s", err)
	}
	if !match {
		t.Errorf(`TestCheckPasswordHash(password, hashedPassword) = %t, want match for true`, match)
	}
}
