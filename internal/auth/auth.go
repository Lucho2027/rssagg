package auth

import (
	"errors"
	"net/http"
	"strings"
)

// GetAPIKey extracts API Key from
// the headers of a HTTP request
// Example:
// Authorization: ApiKey { instert apikey here}
func GetAPIKey(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("o authentication info found")
	}

	vals := strings.Split(val, " ")

	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed  first part of auth header")
	}
	return vals[1], nil
}
