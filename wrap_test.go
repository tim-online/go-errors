package errors_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tim-online/go-errors"
)

func TestWrap(t *testing.T) {
	err := errors.New("Original error")

	wrapped := errors.Wrap(err, "Wrapped error")
	assert.NotNil(t, wrapped, "Wrapped error should not be nil")
	assert.EqualError(t, wrapped, "Wrapped error: Original error")

	cause := errors.Cause(wrapped)
	assert.NotNil(t, cause, "Cause should not be nil")
	assert.EqualError(t, cause, "Original error")

}
