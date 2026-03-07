package auth

import (
	"net/http"
	"testing"
)

func TestGetBearerToken(t *testing.T) {
	cases := []struct {
		name   string
		input  string
		output string
	}{
		{
			name:   "GetBearer Test 1",
			input:  "AtJO5accYd1kqs2G730IYA2MtEuUL",
			output: "AtJO5accYd1kqs2G730IYA2MtEuUL",
		},
		{
			name:   "GetBearer Test 2",
			input:  "RXbcZ4b18SkOkJ+MiwJuzB8",
			output: "RXbcZ4b18SkOkJ+MiwJuzB8",
		},
		{
			name:   "GetBearer Test 3",
			input:  "VSgajbwOrXcO+tzEVzbB4KqfD13zotTCH7Tuv2",
			output: "VSgajbwOrXcO+tzEVzbB4KqfD13zotTCH7Tuv2",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			header := http.Header{}
			header.Set("Authorization", tc.input)
			str, err := GetBearerToken(header)
			if err != nil {
				t.Fatal("Execution Error:", err)
			}

			if str != tc.output {
				t.Fatal("Not matched")
			}
		})
	}
}
