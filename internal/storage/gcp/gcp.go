package gcp

import (
	"context"
	"fmt"
	"strings"

	"cloud.google.com/go/storage"
	"github.com/ahmedennaifer/blov/internal/blob"
	"github.com/ahmedennaifer/blov/internal/config/gcp"
	"google.golang.org/api/iterator"
)

type GCPStorage struct {
	Client    *storage.Client
	GcpConfig gcp.GoogleCloudConfig
}

func NewGCPStorageFromConfig(ctx context.Context, config gcp.GoogleCloudConfig) (*GCPStorage, error) {
	config.Read()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("error creating a new cloud storage client: %v", err)
	}

	return &GCPStorage{
		Client:    client,
		GcpConfig: config,
	}, nil
}

func (s *GCPStorage) CountBuckets(ctx context.Context) (int, error) {
	s.GcpConfig.Read()
	buckets, err := s.ListAll(ctx)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}
	return len(buckets), nil
}

func (s *GCPStorage) ListAll(ctx context.Context, args ...any) ([]blob.Blob, error) {
	s.GcpConfig.Read()
	it := s.Client.Buckets(ctx, s.GcpConfig.ProjectId)
	blobs := make([]blob.Blob, 0)
	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return blobs, fmt.Errorf("error listing buckets: %v", err)
		}

		if s.GcpConfig.Region != "" && !strings.EqualFold(attrs.Location, s.GcpConfig.Region) {
			blob := blob.NewGCPBlobFromAttrs(attrs)
			blobs = append(blobs, blob)
		}

	}
	return blobs, nil
}

func (s *GCPStorage) List(ctx context.Context, bucket string, prefix string) ([]blob.Blob, error) {
	it := s.Client.Bucket(bucket).Objects(ctx, &storage.Query{Prefix: prefix})
	blobs := make([]blob.Blob, 0)

	for {
		attrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return blobs, fmt.Errorf("error listing objects: %v", err)
		}

		blob := blob.Blob{
			Name:         attrs.Name,
			Size:         attrs.Size,
			LastModified: attrs.Updated,
			IsFolder:     strings.HasSuffix(attrs.Name, "/"),
			Provider:     "gcp",
		}
		blobs = append(blobs, blob)
	}

	return blobs, nil
}
