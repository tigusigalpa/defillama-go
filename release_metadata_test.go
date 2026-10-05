package defillama

import (
	"os"
	"strings"
	"testing"
)

func TestReleaseMetadataMatchesCurrentRelease(t *testing.T) {
	const releasedVersion = "1.2.1"
	if Version != releasedVersion {
		t.Fatalf("Version = %q, want %q", Version, releasedVersion)
	}

	changelog, err := os.ReadFile("CHANGELOG.md")
	if err != nil {
		t.Fatal(err)
	}
	heading := "## [" + Version + "]"
	if !strings.Contains(string(changelog), heading) {
		t.Fatalf("CHANGELOG.md does not contain %q", heading)
	}
}
