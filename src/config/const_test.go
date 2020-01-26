package config

import (
	"path/filepath"
	"runtime"
    "testing"
    "github.com/stretchr/testify/assert"
)

func Test_URLs(t *testing.T) {
	assert.Equal(t, DevServerURL, "https://pitaka-server-dev.herokuapp.com")
	assert.Equal(t, PrdServerURL, "https://pitaka-server.herokuapp.com")
}

func Test_RootDIR(t *testing.T) {
	var _, currentFilePath, _, _ = runtime.Caller(0) 
	var actualRootDIR = filepath.Dir(filepath.Dir(filepath.Dir(currentFilePath)))
	assert.Equal(t, RootDIR, actualRootDIR)
}