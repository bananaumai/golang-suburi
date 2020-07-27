package main

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestError(t *testing.T) {
	e1 := errors.New("test")
	e2 := fmt.Errorf("test")

	assert.Equal(t, e1, e2)
}
