package auth

import (
	"testing"
)

func TestParseGcloudVersion(t *testing.T) {
	input := `Google Cloud SDK 541.0.0
beta 2025.09.29
bq 2.1.24
core 2025.09.29
gcloud-crc32c 1.0.0
gke-gcloud-auth-plugin 0.5.10
gsutil 5.35
Updates are available for some Google Cloud CLI components.  To install them,
please run:`
	expected := "Google Cloud SDK 541.0.0" + "\n"
	given := parseGcloudVersion(input)
	if given != expected {
		t.Errorf("Expected: %v, got: %v", expected, given)
	}
}
