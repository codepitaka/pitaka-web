package config

import (
	"path/filepath"
	"runtime"
)

var (
	_, currentFilePath, _, _ = runtime.Caller(0)
	// RootDIR = filepath.Join(filepath.Dir(currentFilePath), "../..")
	RootDIR = filepath.Dir(currentFilePath)
)
var DevServerURL = "https://pitaka-server-dev.herokuapp.com"
var PrdServerURL = "https://pitaka-server.herokuapp.com"