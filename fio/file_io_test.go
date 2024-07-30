package fio

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func destroyFile(name string) {
	if err := os.RemoveAll(name); err != nil {
		panic(err)
	}
}

func TestNewFileIOManager(t *testing.T) {
	path := filepath.Join("D:\\golang_pro\\kv_project\\tmp", "a1.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	fio.Close()
}

func TestFileIO_Write(t *testing.T) {
	path := filepath.Join("D:\\golang_pro\\kv_project\\tmp", "a2.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	n1, err := fio.Write([]byte("hello"))
	t.Log(n1, err)
	n2, err := fio.Write([]byte("world"))
	t.Log(n2, err)

	n3, err := fio.Write([]byte(""))
	assert.Equal(t, 0, n3)
	assert.Nil(t, err)

	fio.Close()
}

func TestFileIO_Read(t *testing.T) {
	path := filepath.Join("D:\\golang_pro\\kv_project\\tmp", "a3.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	_, err = fio.Write([]byte("key-a"))
	assert.Nil(t, err)

	_, err = fio.Write([]byte("key-b"))
	assert.Nil(t, err)

	b := make([]byte, 5)
	n, err := fio.Read(b, 0)
	t.Log(b, n)

	fio.Close()
}

func TestFileIO_Sync(t *testing.T) {
	path := filepath.Join("D:\\golang_pro\\kv_project\\tmp", "a4.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Sync()
	assert.Nil(t, err)

	fio.Close()
}

func TestFileIO_Close(t *testing.T) {
	path := filepath.Join("D:\\golang_pro\\kv_project\\tmp", "a5.data")
	fio, err := NewFileIOManager(path)
	defer destroyFile(path)
	assert.Nil(t, err)
	assert.NotNil(t, fio)

	err = fio.Close()
	assert.Nil(t, err)
}
