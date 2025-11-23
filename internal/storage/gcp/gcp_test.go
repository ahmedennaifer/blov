package gcp

import (
	"context"
	"os"
	"testing"

	"cloud.google.com/go/storage"
	"github.com/ahmedennaifer/blov/internal/config/gcp"
)

// all tests run on : docker run -d -p 4443:4443 --name fake-gcs fsouza/fake-gcs-server -scheme http

func TestListAll(t *testing.T) {
	// must set host before running
	if os.Getenv("STORAGE_EMULATOR_HOST") == "" {
		t.Skip("No storage emulator available")
	}

	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	if err != nil {
		t.Fatalf("Failed to create storage client: %v", err)
	}
	defer client.Close()

	projectID := "test-project"

	bucket1 := client.Bucket("test-bucket-europe")
	bucket2 := client.Bucket("test-bucket-us")

	if err := bucket1.Create(ctx, projectID, &storage.BucketAttrs{
		Location: "EUROPE-WEST1",
	}); err != nil {
		t.Logf("Bucket might already exist: %v", err)
	}

	if err := bucket2.Create(ctx, projectID, &storage.BucketAttrs{
		Location: "US-CENTRAL1",
	}); err != nil {
		t.Logf("Bucket might already exist: %v", err)
	}

	gcpStorage := &GCPStorage{
		Client: client,
		GcpConfig: gcp.GoogleCloudConfig{
			ProjectId: projectID,
			Region:    "europe-west1",
		},
	}

	blobs, err := gcpStorage.ListAll(ctx)
	if err != nil {
		t.Fatalf("ListAll failed: %v", err)
	}

	if len(blobs) == 0 {
		t.Error("Expected at least 1 blob, got 0")
	}
}
