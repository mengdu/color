package color

import (
	"os"
	"strconv"

	"github.com/mattn/go-isatty"
)

var enable int = -1 // -1: not initialized, 0: disabled, 1: enabled

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
	envForceColor, _ := strconv.ParseBool(os.Getenv("FORCE_COLOR"))
	fd := os.Stdout.Fd()
	return envForceColor || (isatty.IsTerminal(fd) || isatty.IsCygwinTerminal(fd))
}
