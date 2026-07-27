package auth

import (
	"errors"
	"net/http"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	goodHeaderGoodApiKey := http.Header{"Authorization": {"ApiKey 123"}}
	goodHeaderBadApiKey := http.Header{"Authorization": {"pikey"}}
	goodHeaderEmptyApiKey := http.Header{"Authorization": {"Apikey "}}
	badHeader := http.Header{"Authorizat": {"Apikey"}}
	tests := map[string]struct {
		header     http.Header
		wantString string
		wantErr    error
	}{
		"Good Header and ApiKey":       {header: goodHeaderGoodApiKey, wantString: "123", wantErr: nil},
		"Good Header and Bad ApiKey":   {header: goodHeaderBadApiKey, wantString: "", wantErr: errors.New("malformed authorization header")},
		"Good Header and Empty ApiKey": {header: goodHeaderEmptyApiKey, wantString: "", wantErr: nil},
		"Bad Header":                   {header: badHeader, wantString: "", wantErr: ErrNoAuthHeaderIncluded},
	}
	for name, tc := range tests {
		got, gotErr := GetAPIKey(tc.header)
		t.Run(name, func(t *testing.T) {
			if gotErr != tc.wantErr && got != tc.wantString {
				t.Fatalf("expected: %v, got: %v\n expectedErr: %v, gotErr: %v",
					tc.wantString, got, tc.wantErr, gotErr)
			}

		})
	}
}
