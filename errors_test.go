package errors_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/tim-online/go-errors"
)

func TestIntegration(t *testing.T) {
	baseError := errors.New("Base Error")
	wrappedError := errors.Wrap(baseError, "Wrapped Error")
	contextError := errors.WithContext(wrappedError, map[string]interface{}{
		"key": "value",
	})

	errGroup := errors.Group(baseError, wrappedError, contextError)
	wrappedErrGroup := errors.Wrap(errGroup, "An error occured")
	log.Println(wrappedErrGroup)

	err, ok := errGroup.(errors.StackTracer)
	if !ok {
		panic("oops, err does not implement stackTracer")
	}

	st := err.StackTrace()
	fmt.Printf("%+v\n", st[:]) // top two frames
}
