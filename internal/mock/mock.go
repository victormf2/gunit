package mock

import (
	"reflect"
	"slices"
	"sync"

	"github.com/victormf2/gunit/internal/expect"
	"github.com/victormf2/gunit/internal/mock/matchers"
	"github.com/victormf2/gunit/internal/utils"
)

// call represents a single function call with its arguments and return values
type call struct {
	args           []any
	returnedValues []any
}

func (c *call) Args() []any {
	return slices.Clone(c.args)
}

func (c *call) ReturnedValues() []any {
	return slices.Clone(c.returnedValues)
}

// mockFunction wraps a function and records all calls to it
type mockFunction struct {
	mu                    sync.RWMutex
	defaultImplementation reflect.Value
	setups                []*mockFunctionSetup
	calls                 []*call
	name                  string
}

// NewMockFunction creates a new mock function wrapper
func NewMockFunction(name string, fn any) MockFunction {
	fnValue := getFunction(fn)

	return &mockFunction{
		name:                  name,
		defaultImplementation: fnValue,
		calls:                 []*call{},
		setups:                []*mockFunctionSetup{},
	}
}

func (m *mockFunction) Name() string {
	return m.name
}

// Call invokes the mock function with the provided arguments
func (m *mockFunction) Call(args ...any) []any {
	fn := m.defaultImplementation
	// Call the function
	for _, setup := range m.setups {
		setup.mu.RLock()

		if !setup.canCall() {
			setup.mu.RUnlock()
			continue
		}
		matchedAll := true
		for i, argMatcher := range setup.expectedArgs {
			matchResult := argMatcher.Match(args[i])
			if !matchResult.Matches() {
				matchedAll = false
				break
			}
		}
		setup.mu.RUnlock()

		if matchedAll {
			setup.mu.Lock()

			if !setup.canCall() {
				setup.mu.Unlock()
				continue
			}

			setup.calls += 1
			fn = setup.implementation
			setup.mu.Unlock()
			break
		}
	}

	if args == nil {
		args = []any{}
	}

	returnedValues := callFn(fn, args)

	// Record the call
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = append(m.calls, &call{
		args:           args,
		returnedValues: returnedValues,
	})

	return returnedValues
}

func (m *mockFunction) Calls() []Call {
	m.mu.RLock()
	defer m.mu.RUnlock()

	clonedCalls := utils.SliceAs[*call, Call](m.calls)
	return clonedCalls
}

func (m *mockFunction) ResetCalls() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.calls = []*call{}
}

func (m *mockFunction) Always(fn any) {
	fnValue := getFunction(fn)
	m.defaultImplementation = fnValue
}

type mockFunctionSetup struct {
	mu             sync.RWMutex
	mockFunction   *mockFunction
	expectedArgs   []expect.Matcher
	implementation reflect.Value
	maxCalls       *int
	calls          int
}

func (m *mockFunctionSetup) clone() *mockFunctionSetup {
	newSetup := &mockFunctionSetup{
		mockFunction:   m.mockFunction,
		expectedArgs:   m.expectedArgs,
		implementation: m.implementation,
		maxCalls:       m.maxCalls,
		calls:          m.calls,
	}
	return newSetup
}

func (m *mockFunctionSetup) canCall() bool {
	return m.maxCalls == nil || m.calls < *m.maxCalls
}

func (m *mockFunction) On(args ...any) MockFunctionSetup {
	argMatchers := []expect.Matcher{}
	for _, arg := range args {
		argMatcher := getArgMatcher(arg)
		argMatchers = append(argMatchers, argMatcher)
	}
	setup := &mockFunctionSetup{
		mockFunction: m,
		expectedArgs: argMatchers,
	}
	return setup
}

func (s *mockFunctionSetup) Return(values ...any) MockFunctionSetup {
	implementation := func(v ...any) *aggregateResults {
		return &aggregateResults{
			values: values,
		}
	}

	newSetup := s.clone()
	newSetup.implementation = reflect.ValueOf(implementation)

	return newSetup
}

func (s *mockFunctionSetup) Once() MockFunctionSetup {
	newSetup := s.clone()
	maxCalls := 1
	newSetup.maxCalls = &maxCalls
	return newSetup
}
func (s *mockFunctionSetup) Times(times int) MockFunctionSetup {
	newSetup := s.clone()
	newSetup.maxCalls = &times
	return newSetup
}

func getFunction(value any) reflect.Value {
	fnValue := reflect.ValueOf(value)
	if fnValue.Kind() != reflect.Func {
		panic("provided value is not a function")
	}
	return fnValue
}

func getArgMatcher(arg any) expect.Matcher {
	switch V := arg.(type) {
	case expect.Matcher:
		return V
	default:
		// This uses a symbolic link to avoid import cycle
		return matchers.NewEqualMatcher(arg)
	}
}

type aggregateResults struct {
	values []any
}

func callFn(fn reflect.Value, args []any) []any {
	// Convert args to reflect.Value slice
	reflectArgs := utils.Map(args, func(_ int, arg any) reflect.Value {
		return reflect.ValueOf(arg)
	})

	results := fn.Call(reflectArgs)
	if len(results) == 1 {
		aggregateResults, ok := results[0].Interface().(*aggregateResults)
		if ok {
			return aggregateResults.values
		}
	}

	returnedValues := utils.Map(results, func(_ int, result reflect.Value) any {
		return result.Interface()
	})
	return returnedValues
}
