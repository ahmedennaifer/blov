package storage

import (
	"context"

	"github.com/ahmedennaifer/blov/internal/structs"
)

type Storer interface {
	List(ctx context.Context, bucket string, prefix string) ([]structs.BlobInfo, error)
}
