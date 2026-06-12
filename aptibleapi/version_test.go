package aptibleapi

import "testing"

func TestAptibleClientVersion(t *testing.T) {
	v := aptibleClientVersion()
	if v == "unknown" {
		t.Errorf("aptibleClientVersion() returned %q", v)
	}
}

func TestUserAgent(t *testing.T) {
	cfg := NewAPIConfiguration()
	expected := "aptible/aptible-api-go/" + aptibleClientVersion()
	if cfg.UserAgent != expected {
		t.Errorf("UserAgent = %q, want %q", cfg.UserAgent, expected)
	}
}
