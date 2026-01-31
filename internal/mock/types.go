package mock

type MockFunction interface {
	Always(fn any)
	Name() string
	Call(args ...any) []any
	Calls() []Call
	On(args ...any) MockFunctionSetup
	ResetCalls()
}

type Call interface {
	Args() []any
	ReturnedValues() []any
}

type MockFunctionSetup interface {
	Return(values ...any) MockFunctionSetup
	Once() MockFunctionSetup
	Times(times int) MockFunctionSetup
}
