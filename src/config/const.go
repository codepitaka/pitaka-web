package config

import (
	"path/filepath"
	"runtime"
)

var _, currentFilePath, _, _ = runtime.Caller(0) 
var RootDIR = filepath.Dir(filepath.Dir(filepath.Dir(currentFilePath)))

var DevServerURL = "https://pitaka-server-dev.herokuapp.com"
var PrdServerURL = "https://pitaka-server.herokuapp.com"
