package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTestKey(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := GetTestKey(i)
		assert.NotNil(t, string(res))
	}
}

func TestRandomValue(t *testing.T) {
	for i := 0; i < 10; i++ {
		res := RandomValue(10)
		assert.NotNil(t, string(res))
	}
}
