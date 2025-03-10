package inmemory

import "github.com/AnnDutova/in-memory-db/internal/engine/db"

type (
	storageType map[string]string

	storageImpl struct {
		storage storageType
	}
)

func New() (db.Storage, error) {
	return &storageImpl{storage: make(storageType, 0)}, nil
}

func (s *storageImpl) Set(key, value string) {
	s.storage[key] = value
}

func (s *storageImpl) Get(key string) (string, bool) {
	v, ok := s.storage[key]
	return v, ok
}

func (s *storageImpl) Delite(key string) {
	delete(s.storage, key)
}

func (s *storageImpl) Length() int {
	return len(s.storage)
}
