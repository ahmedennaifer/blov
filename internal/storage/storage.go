package storage

import (
	"context"

	"github.com/ahmedennaifer/blov/internal/blob"
)

type Storer interface {
	ListAll(ctx context.Context, args ...any) ([]blob.Blob, error)
	List(ctx context.Context, bucket string, prefix string) ([]blob.Blob, error)
}
