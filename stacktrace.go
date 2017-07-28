package errors

import (
	stack "github.com/pkg/errors"
)

var (
	WithStack = stack.WithStack
)

// Create interface for captureErrorAndWait to check for
type StackTracer interface {
	StackTrace() stack.StackTrace
}
