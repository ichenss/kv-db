package utils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestDirSize(t *testing.T) {
	dir, _ := os.Getwd()
	dirSize, err := DirSize(dir)
	assert.Nil(t, err)
	assert.True(t, dirSize > 0)
}

func TestAvailableDiskSize(t *testing.T) {
	size, err := AvailableDiskSize()
	// 换算成 gb
	t.Log(size / 1024 / 1024 / 1024)
	assert.Nil(t, err)
	assert.True(t, size > 0)
}
