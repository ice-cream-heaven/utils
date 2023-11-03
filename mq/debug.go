//go:build !release

package mq

import (
	"github.com/ice-cream-heaven/utils/runtime"
	"path/filepath"
)

func DataPath() string {
	return filepath.Join(runtime.ExecDir(), "temp", "mq")
}
