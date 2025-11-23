package tui

import (
	"context"
	"fmt"
	"time"

	gcpconfig "github.com/ahmedennaifer/blov/internal/config/gcp"
	storage "github.com/ahmedennaifer/blov/internal/storage/gcp"
	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	width       int
	height      int
	ctx         context.Context
	cancel      context.CancelFunc
	buckets     []BucketItem
	objects     []ObjectItem
	menuIndex   int
	bucketIndex int
	objectIndex int
	activePane  int
	currentPath string
	loading     bool
	status      string
	provider    string
	region      string
}

type BucketItem struct {
	Name     string
	Provider string
	Region   string
}

type ObjectItem struct {
	Name         string
	Size         string
	LastModified string
}

var menuItems = []string{
	"Browse",
	"Login",
	"Config",
	"Dash",
}

func NewModel() Model {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	return Model{
		ctx:         ctx,
		cancel:      cancel,
		buckets:     []BucketItem{},
		objects:     []ObjectItem{},
		menuIndex:   0,
		bucketIndex: 0,
		objectIndex: 0,
		activePane:  0,
		currentPath: "/",
		loading:     true,
		status:      "Ready",
		provider:    "GCP",
		region:      "europe-west1",
	}
}

func (m Model) Init() tea.Cmd {
	return m.loadBuckets()
}

func (m Model) loadBuckets() tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		config := gcpconfig.NewGCPConfig()
		if err := config.Read(); err != nil {
			return ErrorMsg{err: err}
		}

		gcpStorage, err := storage.NewGCPStorageFromConfig(ctx, *config)
		if err != nil {
			return ErrorMsg{err: err}
		}

		blobs, err := gcpStorage.ListAll(ctx)
		if err != nil {
			return ErrorMsg{err: err}
		}

		var buckets []BucketItem
		for _, blob := range blobs {
			buckets = append(buckets, BucketItem{
				Name:     blob.Name,
				Provider: blob.Provider,
				Region:   "europe-west1",
			})
		}

		return BucketsLoadedMsg{buckets: buckets}
	}
}

func (m Model) loadObjects(bucketName string) tea.Cmd {
	return func() tea.Msg {
		ctx := context.Background()
		config := gcpconfig.NewGCPConfig()
		if err := config.Read(); err != nil {
			return ErrorMsg{err: err}
		}

		gcpStorage, err := storage.NewGCPStorageFromConfig(ctx, *config)
		if err != nil {
			return ErrorMsg{err: err}
		}

		blobs, err := gcpStorage.List(ctx, bucketName, "")
		if err != nil {
			return ErrorMsg{err: err}
		}

		var objects []ObjectItem
		for _, blob := range blobs {
			size := fmt.Sprintf("%d bytes", blob.Size)
			if blob.Size > 1024*1024 {
				size = fmt.Sprintf("%.1fMB", float64(blob.Size)/(1024*1024))
			} else if blob.Size > 1024 {
				size = fmt.Sprintf("%.1fKB", float64(blob.Size)/1024)
			}

			objects = append(objects, ObjectItem{
				Name:         blob.Name,
				Size:         size,
				LastModified: blob.LastModified.Format("2006-01-02 15:04"),
			})
		}

		return ObjectsLoadedMsg{objects: objects}
	}
}

type BucketsLoadedMsg struct {
	buckets []BucketItem
}

type ObjectsLoadedMsg struct {
	objects []ObjectItem
}

type ErrorMsg struct {
	err error
}

