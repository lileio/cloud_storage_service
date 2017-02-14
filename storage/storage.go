package storage

import "context"

type Storage interface {
	Setup() error
	PublicURL(filename string) string
	Store(ctx context.Context, filename string, data []byte, metadata map[string]string) error
	Delete(ctx context.Context, filename string) error
}
