package helpers

import (
	"fmt"
	"strings"
)

// GetEndpointType converts a user-facing endpoint type to the API value.
func GetEndpointType(t string) (string, error) {
	switch strings.ToLower(t) {
	case "https":
		return "http", nil
	case "tcp":
		return "tcp", nil
	case "tls":
		return "tls", nil
	case "grpc":
		return "grpc", nil
	default:
		return "", fmt.Errorf("invalid endpoint type, please use HTTPS, TLS, GRPC, or TCP")
	}
}

// GetHumanReadableEndpointType converts an API endpoint type to a user-facing value.
func GetHumanReadableEndpointType(t string) (string, error) {
	switch t {
	case "http_proxy_protocol", "http":
		return "https", nil
	case "tcp":
		return "tcp", nil
	case "tls":
		return "tls", nil
	case "grpc":
		return "grpc", nil
	default:
		return "", fmt.Errorf("invalid endpoint type - %s", t)
	}
}
