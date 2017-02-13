package storage

import "context"

type Storage interface {
	Setup() error
	Store(ctx context.Context, filename string, data []byte, metadata map[string]string) error
	Delete(ctx context.Context, filename string) error
}
