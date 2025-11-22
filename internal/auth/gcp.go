package auth

import (
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type GCPAuthenticator struct{}

// maybe from fime, or sa.json ?

func NewGCPAuthenticator() *GCPAuthenticator {
	return &GCPAuthenticator{}
}

func parseGcloudVersion(versionInfo string) string {
	gcloudVersionLine := strings.Split(versionInfo, "\n")[0]
	return gcloudVersionLine + "\n"
}

func getGcloudVersion(ctx context.Context) (string, error) {
	cmd := exec.CommandContext(ctx, "gcloud", "version")
	version, err := cmd.Output()
	if err != nil {
		fmt.Printf("Error: %v", err)
		return "", err
	}
	gcloudVersion := parseGcloudVersion(string(version))
	return gcloudVersion, nil
}

func (g *GCPAuthenticator) Login(ctx context.Context) error {
	version, err := getGcloudVersion(ctx)
	if err != nil {
		fmt.Printf("Error: %v ", err)
	}
	fmt.Printf("Found gcloud: %v", version)
	return nil
}
