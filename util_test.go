package main

import (
	"errors"
	"github.com/plombardi89/gozeug/randomzeug"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func Test_getServerID(t *testing.T) {
	id := getServerID(nil, randomzeug.NewRandom())
	assert.Equal(t, func() string {
		v, _ := os.Hostname()
		return v
	}(), id)

	id = getServerID(os.Hostname, randomzeug.NewRandom())
	assert.Equal(t, func() string {
		v, _ := os.Hostname()
		return v
	}(), id)

	id = getServerID(func() (string, error) { return "foo", nil }, randomzeug.NewRandom())
	assert.Equal(t, "foo", id)

	id = getServerID(func() (string, error) { return "", errors.New("error") }, randomzeug.NewRandom())
	assert.NotEmpty(t, id)
}
