package nilerr_test

import (
	"fmt"
	"testing"

	"github.com/chrisdoherty4/nilerr"
	"github.com/stretchr/testify/assert"
)

func TestWrappedNilError(t *testing.T) {
	err := fmt.Errorf("foo: %w", nilerr.New("bar"))
	assert.True(t, nilerr.Is(err))
}
