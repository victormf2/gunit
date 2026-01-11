package repository

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/victormf2/testo/example/domain"
	"github.com/victormf2/testo/expect"
	"github.com/victormf2/testo/mock"
)

func TestTestinho(t *testing.T) {
	ctx := t.Context()

	mockRepo := NewMockRepository()
	mockRepo.GetUser(ctx, "userId")

	require.Len(t, mockRepo.MockGetUser.Calls(), 1)

	actualCall := mockRepo.MockGetUser.Calls()[0]
	expectedCall := mock.Call{
		Args:    []any{ctx, "userId"},
		Returns: []any{(*domain.User)(nil), error(nil)},
	}
	require.Equal(t, expectedCall, actualCall)

	mockRepo.SomethingReturnsInt()

	require.Len(t, mockRepo.MockSomethingReturnsInt.Calls(), 1)

	actualCall = mockRepo.MockSomethingReturnsInt.Calls()[0]
	expectedCall = mock.Call{
		Args:    []any{},
		Returns: []any{0},
	}
	require.Equal(t, expectedCall, actualCall)

	expect.It(mockRepo.MockSomethingReturnsInt).
		ToHaveBeenCalled(t, expect.Times(1))

	mockRepo.MockSomethingReturnsInt.Reset()
	require.Len(t, mockRepo.MockSomethingReturnsInt.Calls(), 0)
}
