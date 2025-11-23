package structs

import "time"

type BlobInfo struct {
	Name         string
	Size         int64
	LastModified time.Time
	IsFolder     bool
	Provider     string // prefix
}
