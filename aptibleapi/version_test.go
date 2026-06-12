package aptibleapi

import "testing"

func TestAptibleClientVersion(t *testing.T) {
	v := aptibleClientVersion()
	if v == "aptible/aptible-api-go/unknown" {
		t.Errorf("aptibleClientVersion() returned %q", v)
	}
}
