package errors_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tim-online/go-errors"
)

func TestContext(t *testing.T) {
	errWithoutContext := errors.New("Without Context")
	errWithContext := errors.WithContext(errWithoutContext, map[string]interface{}{
		"key": "value",
		"k":   "v",
	})

	assert.Contains(t, errors.Context(errWithContext), "key")
	assert.Contains(t, errors.Context(errWithContext), "k")
	assert.Equal(t, errors.Context(errWithContext)["key"], "value")
	assert.Equal(t, errors.Context(errWithContext)["k"], "v")

	assert.EqualError(t, errWithContext, "Without Context [k:v, key:value]")
}

type customContextError struct {
	error
}

func (e customContextError) Context() map[string]interface{} {
	return map[string]interface{}{
		"key": "value",
	}
}

func TestCustomContextError(t *testing.T) {
	contextError := customContextError{
		error: fmt.Errorf("Base error"),
	}

	log.Println(contextError)
}
