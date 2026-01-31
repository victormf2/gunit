package internal

type MockT struct {
	Failed bool
}

func (m *MockT) Fatalf(format string, args ...any) {
	m.Failed = true
}
func (m *MockT) Fatal(args ...any) {
	m.Failed = true
}
func (m *MockT) Helper() {}
