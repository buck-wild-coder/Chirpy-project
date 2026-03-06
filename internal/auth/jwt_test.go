package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestJwt(t *testing.T) {
	type input struct {
		userID      uuid.UUID
		tokenSecret string
		expiresIn   time.Duration
	}

	cases := []struct {
		name                string
		in                  input
		expectValidateError bool
	}{
		{
			name: "valid token",
			in: input{
				userID:      uuid.New(),
				tokenSecret: "secret-key",
				expiresIn:   1 * time.Hour,
			},
			expectValidateError: false,
		},
		{
			name: "expired token",
			in: input{
				userID:      uuid.New(),
				tokenSecret: "secret-key",
				expiresIn:   -1 * time.Hour, // already expired
			},
			expectValidateError: true,
		},
		{
			name: "Another secret",
			in: input{
				userID:      uuid.New(),
				tokenSecret: "correct-secret",
				expiresIn:   1 * time.Hour,
			},
			expectValidateError: false,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			tokenString, err := MakeJWT(tc.in.userID, tc.in.tokenSecret, tc.in.expiresIn)
			if err != nil {
				t.Fatalf("unexpected error from MakeJWT: %v", err)
			}

			userID, err := ValidateJWT(tokenString, tc.in.tokenSecret)
			if tc.expectValidateError {
				if err == nil {
					t.Fatalf("expected validation error, got none")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected validation error: %v", err)
			}
			if tc.in.userID != userID {
				t.Error("User id do not match")
			}
		})
	}
}
