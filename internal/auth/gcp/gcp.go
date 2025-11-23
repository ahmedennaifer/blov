package gcp

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/ahmedennaifer/blov/internal/config/gcp"
	storage "github.com/ahmedennaifer/blov/internal/storage/gcp"
)

type GCPAuthenticator struct {
	Config gcp.GoogleCloudConfig
}

// maybe from fime, or sa.json ?

func NewGCPAuthenticator() *GCPAuthenticator {
	return &GCPAuthenticator{
		Config: gcp.GoogleCloudConfig{},
	}
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
	fmt.Printf("\033[32mFound gcloud: %v\033[0m", gcloudVersion)
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
	return nil
}

func (g *GCPAuthenticator) Verify(ctx context.Context) error {
	gcpStorage, err := storage.NewGCPStorageFromConfig(ctx, g.Config)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	numberOfBuckets, err := gcpStorage.CountBuckets(ctx)
	if err != nil {
		return fmt.Errorf("%v\n", err)
	}
	fmt.Printf("\033[32mLogged in with success and found %v buckets\033[0m\n", numberOfBuckets)
	return nil
}
