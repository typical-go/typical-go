package typgo

import (
	"fmt"
)

type (
	// MockCommandRunner mocking bash
	MockCommandRunner struct {
		Mocks []*MockCommand
		Ptr   int
	}
	// MockCommand is the test expectation
	MockCommand struct {
		CommandLine string
		OutputBytes []byte
		ErrorBytes  []byte
		ReturnError error
	}
)

// Close mocker
func (r *MockCommandRunner) Close() error {
	if expectation := r.Expectation(); expectation != nil {
		return fmt.Errorf("missing bash call: \"%s\"", expectation.CommandLine)
	}
	r.Ptr = 0
	return nil
}

// Expectation for bash
func (r *MockCommandRunner) Expectation() *MockCommand {
	if r.Ptr < len(r.Mocks) {
		expect := r.Mocks[r.Ptr]
		r.Ptr++
		return expect
	}
	return nil
}

// Run bash
func (r *MockCommandRunner) Run(bash *Command) error {
	expc := r.Expectation()
	if expc == nil {
		return fmt.Errorf("typgo-mock: no run expectation for \"%s\"", bash.String())
	}

	if bash.String() != expc.CommandLine {
		return fmt.Errorf("typgo-mock: \"%s\" should be \"%s\"", bash.String(), expc.CommandLine)
	}

	if expc.OutputBytes != nil && bash.Stdout != nil {
		bash.Stdout.Write(expc.OutputBytes)
	}
	if expc.ErrorBytes != nil && bash.Stderr != nil {
		bash.Stderr.Write(expc.ErrorBytes)
	}

	return expc.ReturnError
}
