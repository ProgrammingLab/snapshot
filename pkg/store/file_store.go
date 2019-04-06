package store

import (
	"context"
	"io"
	"time"
)

// FileStore is interface to access files
type FileStore interface {
	ListFiles(ctx context.Context, folderUUID string) ([]*File, error)
	Get(ctx context.Context, fileUUID string) (io.ReadCloser, error)
	Create(ctx context.Context, name string, folderUUID string, r io.Reader) (*File, error)
	Delete(ctx context.Context, fileUUID string) error
}

// File represents file metadata
type File struct {
	UUID      string
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
