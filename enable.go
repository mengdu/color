package color

import (
	"os"

	"github.com/mattn/go-isatty"
)

var enable int = -1

// EnableColor returns true if color is enabled
func EnableColor(fromCache bool) (result bool) {
	if fromCache {
		if enable != -1 {
			return enable == 1
		}
		defer func() {
			if result {
				enable = 1
			} else {
				enable = 0
			}
		}()
	}
	envVal := os.Getenv("FORCE_COLOR")
	envForceColor := envVal != "" && envVal != "0"
	fd := os.Stdout.Fd()
	return envForceColor || (isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd))
}
