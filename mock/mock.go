package mock

import (
	"reflect"
	"slices"
	"sync"
)

// Call represents a single function call with its arguments and return values
type Call struct {
	Args    []any
	Returns []any
}

// MockFunction wraps a function and records all calls to it
type MockFunction struct {
	mu                    sync.RWMutex
	defaultImplementation reflect.Value
	calls                 []Call
	name                  string
}

// NewMockFunction creates a new mock function wrapper
func NewMockFunction(name string, fn any) *MockFunction {
	fnValue := getFunction(fn)

	return &MockFunction{
		name:                  name,
		defaultImplementation: fnValue,
		calls:                 []Call{},
	}
}

func (m *MockFunction) Name() string {
	return m.name
}

// Call invokes the mock function with the provided arguments
func (m *MockFunction) Call(args ...any) []any {
	// Convert args to reflect.Value slice
	reflectArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectArgs[i] = reflect.ValueOf(arg)
	}

	// Call the function
	results := m.defaultImplementation.Call(reflectArgs)

	// Convert results back to any slice
	returns := make([]any, len(results))
	for i, result := range results {
		returns[i] = result.Interface()
	}

	if args == nil {
		args = []any{}
	}
	// Record the call
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = append(m.calls, Call{
		Args:    args,
		Returns: returns,
	})

	return returns
}

func (m *MockFunction) Calls() []Call {
	m.mu.RLock()
	defer m.mu.RUnlock()
	return slices.Clone(m.calls)
}

func (m *MockFunction) ResetCalls() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = []Call{}
}

func (m *MockFunction) SetDefaultImplementation(fn any) {
	fnValue := getFunction(fn)
	m.defaultImplementation = fnValue
}

func getFunction(value any) reflect.Value {
	fnValue := reflect.ValueOf(value)
	if fnValue.Kind() != reflect.Func {
		panic("provided value is not a function")
	}
	return fnValue
}
