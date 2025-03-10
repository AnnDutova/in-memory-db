package engine

import (
	"go.uber.org/zap"

	"github.com/AnnDutova/in-memory-db/internal/engine/db"
	inmemory "github.com/AnnDutova/in-memory-db/internal/engine/db/in_memory"
)

type (
	engineImpl struct {
		logger  *zap.Logger
		storage db.Storage
	}
)

func New(logger *zap.Logger) (Engine, error) {
	storage, err := inmemory.New()
	if err != nil {
		return nil, err
	}

	return &engineImpl{logger: logger, storage: storage}, nil
}

func (s *engineImpl) Set(key, value string) error {
	s.storage.Set(key, value)
	return nil
}

func (s *engineImpl) Get(key string) (string, error) {
	if v, ok := s.storage.Get(key); ok {
		return v, nil
	}
	return "", ErrNotFound
}

func (s *engineImpl) Delite(key string) error {
	s.storage.Delite(key)
	return nil
}
