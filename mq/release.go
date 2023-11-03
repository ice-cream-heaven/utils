//go:build release

package mq

import (
	"github.com/ice-cream-heaven/utils/app"
	"os"
	"path/filepath"
)

func DataPath() string {
	return filepath.Join(os.TempDir(), "ice", app.Name, "mq")
}
