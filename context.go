package errors

import (
	"fmt"
	"sort"
	"strings"
)

type errWithContext interface {
	Context() context
}

type context map[string]interface {
}

func Context(err error) context {
	for err != nil {
		err, ok := err.(errWithContext)
		if !ok {
			break
		}
		return err.Context()
	}
	return nil
}

func WithContext(err error, context context) error {
	if err == nil {
		return nil
	}
	return &withContext{
		error:   WithStack(err),
		context: context,
	}
}

type withContext struct {
	error
	context context
}

func (e *withContext) Context() context {
	return e.context
}

func (e *withContext) Error() string {
	keys := make([]string, len(e.context))
	i := 0
	for k, _ := range e.context {
		keys[i] = k
		i++
	}
	sort.Strings(keys)

	pieces := []string{}
	for _, k := range keys {
		v := e.context[k]
		pieces = append(pieces, fmt.Sprintf("%s:%s", k, v))
	}

	return e.error.Error() + " [" + strings.Join(pieces, ", ") + "]"
}
