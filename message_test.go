package errors_test

// import (
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/tim-online/go-errors"
// )

// func TestMessage(t *testing.T) {
// 	errWithoutContext := errors.New("Without Context")
// 	errWithContext := errors.WithContext(errWithoutContext, map[string]interface{}{
// 		"key": "value",
// 	})

// 	wrappedErr := errors.Wrap(errWithContext, "Wrapped error")

// 	assert.Equal(t, errors.Message(errWithoutContext), "Without Context")
// 	assert.Equal(t, errors.Message(wrappedErr), "Wrapped error")
// }
