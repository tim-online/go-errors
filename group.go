package errors

import multierror "github.com/hashicorp/go-multierror"

var (
// ErrorOrNil    = multierror.ErrorOrNil
// WrappedErrors = multierror.WrappedErrors
)

type Grouper interface {
	Errors() []error
}

type group struct {
	err    *multierror.Error
	tracer error
}

func (g group) Error() string {
	return g.err.Error()
}

// func (g group) StackTrace() StackTrace {
// 	tracer, ok := g.tracer.(StackTracer)
// 	if !ok {
// 		return nil
// 	}
// 	return tracer.StackTrace()
// }

// @TODO: rename?
// - group
// - multi
// - collect(ion)
// - combine
// - collect / split
// - many
// - store
// - list
func Group(errors ...error) error {
	if len(errors) == 0 {
		return nil
	}

	multi := multierror.Append(errors[0], errors[1:]...)
	if len(multi.WrappedErrors()) == 0 {
		return nil
	}

	return group{
		err:    multi,
		tracer: WithStack(multi),
	}
}

func Ungroup(err error) []error {
	for err == nil {
		return []error{}
	}

	multi, ok := err.(group)
	if !ok {
		return []error{err}
	}

	return multi.err.Errors
}
