package repository

import (
	"context"

	"github.com/victormf2/testo/example/domain"
)

type Repository interface {
	GetUser(ctx context.Context, userId string) (*domain.User, error)
	SaveUser(ctx context.Context, user *domain.User) error
	SomethingReturnsInt() int
}
