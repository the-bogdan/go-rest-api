package users

import (
	"context"

	"github.com/the-bogdan/go-rest-api/app/pkg/logging"
)

type service struct {
	storage Storage
	logger  logging.Logger
}

func (s *service) Create(ctx context.Context, dto CreateUserDTO) (*User, error) {

	panic("implement me")
}
