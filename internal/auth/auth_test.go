package auth

import "testing"

func TestHashPassword(t *testing.T) {
	cases := []struct {
		name     string
		password string
	}{
		{name: "weak pass", password: "1234"},
		{name: "strong pass", password: "Pass12fdsa#182?"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			hash, err := HashPassword(c.password)
			if err != nil {
				t.Fatalf("HashPassword error: %v", err)
			}

			match, err := CheckPasswordHash(c.password, hash)
			if err != nil {
				t.Fatalf("CheckPasswordHash error: %v", err)
			}
			if !match {
				t.Errorf("password %q does not match its hash", c.password)
			}
		})
	}
}
