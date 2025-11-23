package blob

import (
	"time"

	"cloud.google.com/go/storage"
)

type Blob struct {
	Name         string
	Size         int64
	LastModified time.Time
	IsFolder     bool
	Provider     string // prefix
}

func NewGCPBlobFromAttrs(attrs *storage.BucketAttrs) Blob {
	return Blob{
		Name:         attrs.Name,
		Size:         0,
		LastModified: attrs.Updated,
		IsFolder:     false,
		Provider:     "gcp",
	}
}
