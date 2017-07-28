package errors

import wrap "github.com/pkg/errors"

var (
	Cause  = wrap.Cause
	New    = wrap.New
	Errorf = wrap.Errorf
	Wrap   = wrap.Wrap
)

type Causer interface {
	Cause() error
}
