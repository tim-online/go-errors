package errors_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	errors "github.com/tim-online/go-errors"
)

func TestGroupLen(t *testing.T) {
	var multi error
	errs := errors.Ungroup(multi)
	assert.Len(t, errs, 0, "Nil multierror should have 0 errors")

	err1 := fmt.Errorf("Error #1")
	err2 := fmt.Errorf("Error #2")
	err3 := fmt.Errorf("Error #3")
	multi = errors.Group(err1, err2, err3)

	errs = errors.Ungroup(multi)
	assert.Len(t, errs, 3, "multierror should have 3 errors")
}

func TestGroupNil(t *testing.T) {
	var multi error

	multi = errors.Group()
	assert.Nil(t, multi)

	multi = errors.Group(nil)
	assert.Nil(t, multi)

	var nilErr error
	multi = errors.Group(nilErr, nil)
	assert.Nil(t, multi)
}
