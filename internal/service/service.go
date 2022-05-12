package service

import (
	"context"

	"github.com/rs/zerolog"
)

type Service struct {
	logger *zerolog.Logger
	repo   Repository
}

type Repository interface {
	Load(string) []byte
}

func New(logger *zerolog.Logger, repo Repository) *Service {
	return &Service{
		logger: logger,
		repo:   repo,
	}
}

func (s *Service) Download(ctx context.Context, filename string) ([]byte, error) {
	//TODO deal with context,
	// auth or smth
	return s.repo.Load(filename), nil
}
