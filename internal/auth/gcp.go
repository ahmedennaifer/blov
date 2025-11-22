package auth

import (
	"context"
	"fmt"
	"os"
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

func checkGcloudInstalled(ctx context.Context) error {
	if _, err := exec.LookPath("gcloud"); err != nil {
		return fmt.Errorf("gcloud CLI is not installed")
	}
	cmd := exec.CommandContext(ctx, "gcloud", "version")
	version, err := cmd.Output()
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}
	gcloudVersion := parseGcloudVersion(string(version))
	fmt.Printf("Found gcloud: %v", gcloudVersion)
	return nil
}

func launchGcloudCmdWithCaptureStd(ctx context.Context, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, "gcloud", args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd
}

func (g *GCPAuthenticator) Login(ctx context.Context) error {
	if err := checkGcloudInstalled(ctx); err != nil {
		return fmt.Errorf("error: %v", err)
	}

	fmt.Println("Launching gcloud...")
	cmd := launchGcloudCmdWithCaptureStd(ctx, "auth", "login")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to run gcloud: %w", err)
	}
	fmt.Println("Logged to gcp with success!")
	return nil
}
