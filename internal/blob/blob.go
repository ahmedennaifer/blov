package blob

import (
	"fmt"
	"os"
	"time"

	"cloud.google.com/go/storage"
	"github.com/olekukonko/tablewriter"
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

func FormatList(blobs []Blob) {
	if len(blobs) == 0 {
		fmt.Println("No blobs found")
		return
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.Header([]string{"Name", "Provider", "Last Modified"})

	for _, blob := range blobs {
		table.Append([]string{
			blob.Name,
			blob.Provider,
			blob.LastModified.Format("2006-01-02 15:04"),
		})
	}

	table.Render()
}
