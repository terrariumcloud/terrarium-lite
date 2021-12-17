package filesystem

import (
	"context"
	"os"
	"path"
)

type TerrariumFilesystemStorage struct {
	path string
}

func (s *TerrariumFilesystemStorage) FetchModuleSource(ctx context.Context, key string) ([]byte, error) {
	fullPath := path.Clean(path.Join(s.path, key))
	return os.ReadFile(fullPath)
}

func (s *TerrariumFilesystemStorage) GetBackingStoreName() string {
	return "filesystem"
}

func New(storageRootPath string) (*TerrariumFilesystemStorage, error) {
	s := &TerrariumFilesystemStorage{
		path: path.Clean(storageRootPath),
	}
	return s, nil
}
