package service

import (
	"context"
	"fmt"
	"os"

	"go.uber.org/zap"

	"github.com/AnnDutova/in-memory-db/internal/compute"
)

type serviceImpl struct {
	logger *zap.Logger
	parser compute.Compute
}

func NewService() (Servicer, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	parser, err := compute.New(logger)
	if err != nil {
		return nil, err
	}

	return &serviceImpl{
		logger: logger,
		parser: parser,
	}, nil
}

func (s *serviceImpl) Run(ctx context.Context) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "recovered from panic", r)
		}
	}()
	err := s.parser.Parse(ctx, os.Stdin, os.Stdout)
	if err != nil {
		return err
	}
	return nil
}
