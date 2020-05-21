package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestT(t *testing.T) {
	size := 100
	assert.EqualValues(t, 100, size, "Len error")
	assert.Equal(t, true, true, "IsEmpty error")
	assert.Nil(t, nil, "Dequeue error")
}
