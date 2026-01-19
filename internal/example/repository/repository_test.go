package repository

import (
	"testing"

	"github.com/victormf2/gunit/expect"
)

func TestTestinho(t *testing.T) {
	ctx := t.Context()

	mockRepo := NewMockRepository()
	mockRepo.GetUser(ctx, "userId")

	expect.It(mockRepo.MockGetUser).
		ToHaveBeenCalled(t, expect.Call().WithArgs(ctx, "userId").Times(1))

	mockRepo.SomethingReturnsInt()
	mockRepo.SomethingReturnsInt()
	mockRepo.SomethingReturnsInt()

	expect.It(mockRepo.MockSomethingReturnsInt).
		ToHaveBeenCalled(t)

	expect.It(mockRepo.MockSomethingReturnsInt).
		ToHaveBeenCalled(t, expect.Call().AtMost(4))

	expect.It(mockRepo.MockSomethingReturnsInt).
		ToHaveBeenCalled(t, expect.Call().AtLeast(1))

	expect.It(mockRepo.MockSomethingReturnsInt).
		ToHaveBeenCalled(t, expect.Call().Times(3))

	mockRepo.MockSomethingReturnsInt.ResetCalls()

	expect.It(mockRepo.MockSomethingReturnsInt).
		ToHaveBeenCalled(t, expect.Call().Never())
}
