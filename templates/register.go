package templates

import (
	"path/filepath"
	"runtime"
)

var (
	_, filePath, _, _ = runtime.Caller(0)
	Directory         = filepath.Dir(filePath)
)
