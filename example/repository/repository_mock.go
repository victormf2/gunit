package repository

import (
	"context"

	"github.com/victormf2/gunit/example/domain"
	"github.com/victormf2/gunit/gunit"
	"github.com/victormf2/gunit/mock"
)

type MockRepository struct {
	MockGetUser             *mock.MockFunction
	MockSaveUser            *mock.MockFunction
	MockSomethingReturnsInt *mock.MockFunction
}

func (m *MockRepository) GetUser(ctx context.Context, userId string) (*domain.User, error) {
	returns := m.MockGetUser.Call(ctx, userId)
	return gunit.As[*domain.User](returns[0]),
		gunit.As[error](returns[1])
}

func (m *MockRepository) SaveUser(ctx context.Context, user *domain.User) error {
	returns := m.MockGetUser.Call(ctx, user)
	return gunit.As[error](returns[0])
}

func (m *MockRepository) SomethingReturnsInt() int {
	returns := m.MockSomethingReturnsInt.Call()
	return gunit.As[int](returns[0])
}

func NewMockRepository() *MockRepository {
	return &MockRepository{
		MockGetUser: mock.NewMockFunction(func(ctx context.Context, userId string) (*domain.User, error) {
			var zero0 *domain.User
			var zero1 error
			return zero0, zero1
		}),
		MockSaveUser: mock.NewMockFunction(func(ctx context.Context, user *domain.User) error {
			var zero0 error
			return zero0
		}),
		MockSomethingReturnsInt: mock.NewMockFunction(func() int {
			var zero0 int
			return zero0
		}),
	}
}

var _ Repository = &MockRepository{}
