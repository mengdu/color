package color

import (
	"fmt"
	"os"

	"github.com/mattn/go-isatty"
)

var enable int = -1

func EnableColor(cache bool) (result bool) {
	if cache {
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

func New(enable bool) ColorFn {
	return func(s string) (string, string, string, bool) { return "", "", s, enable }
}

var Std = New(EnableColor(false))

type ColorFn func(s string) (string, string, string, bool)

func (f ColorFn) color(start, end string) ColorFn {
	return func(s string) (string, string, string, bool) {
		a, b, c, e := f(s)
		return fmt.Sprintf("%s\u001b[%sm", a, start), fmt.Sprintf("%s\u001b[%sm", b, end), c, e
	}
}

func (f ColorFn) String(s string) string {
	start, end, v, enable := f(s)
	if !enable {
		return v
	}
	return start + v + end
}

func (c ColorFn) Black() ColorFn {
	return c.color("30", "39")
}

func (c ColorFn) Red() ColorFn {
	return c.color("31", "39")
}

func (c ColorFn) Green() ColorFn {
	return c.color("32", "39")
}

func (c ColorFn) Blue() ColorFn {
	return c.color("34", "39")
}

func (c ColorFn) Yellow() ColorFn {
	return c.color("33", "39")
}

func (c ColorFn) Magenta() ColorFn {
	return c.color("35", "39")
}

func (c ColorFn) Cyan() ColorFn {
	return c.color("36", "39")
}

func (c ColorFn) White() ColorFn {
	return c.color("37", "39")
}

func (c ColorFn) Gray() ColorFn {
	return c.color("90", "39")
}

func (c ColorFn) RedBright() ColorFn {
	return c.color("91", "49")
}

func (c ColorFn) GreenBright() ColorFn {
	return c.color("92", "49")
}

func (c ColorFn) BlueBright() ColorFn {
	return c.color("94", "49")
}

func (c ColorFn) YellowBright() ColorFn {
	return c.color("93", "49")
}

func (c ColorFn) MagentaBright() ColorFn {
	return c.color("95", "49")
}

func (c ColorFn) CyanBright() ColorFn {
	return c.color("96", "49")
}

func (c ColorFn) WhiteBright() ColorFn {
	return c.color("97", "49")
}

func (c ColorFn) BgBlack() ColorFn {
	return c.color("40", "49")
}

func (c ColorFn) BgRed() ColorFn {
	return c.color("41", "49")
}

func (c ColorFn) BgGreen() ColorFn {
	return c.color("42", "49")
}

func (c ColorFn) BgBlue() ColorFn {
	return c.color("44", "49")
}

func (c ColorFn) BgYellow() ColorFn {
	return c.color("43", "49")
}

func (c ColorFn) BgMagenta() ColorFn {
	return c.color("45", "49")
}

func (c ColorFn) BgCyan() ColorFn {
	return c.color("46", "49")
}

func (c ColorFn) BgWhite() ColorFn {
	return c.color("47", "49")
}

func (c ColorFn) BgGray() ColorFn {
	return c.color("100", "49")
}

func (c ColorFn) BgRedBright() ColorFn {
	return c.color("101", "49")
}

func (c ColorFn) BgGreenBright() ColorFn {
	return c.color("102", "49")
}

func (c ColorFn) BgBlueBright() ColorFn {
	return c.color("104", "49")
}

func (c ColorFn) BgYellowBright() ColorFn {
	return c.color("103", "49")
}

func (c ColorFn) BgMagentaBright() ColorFn {
	return c.color("105", "49")
}

func (c ColorFn) BgCyanBright() ColorFn {
	return c.color("106", "49")
}

func (c ColorFn) BgWhiteBright() ColorFn {
	return c.color("107", "49")
}

func (c ColorFn) Dim() ColorFn {
	return c.color("2", "22")
}

func (c ColorFn) Bold() ColorFn {
	return c.color("1", "22")
}

func (c ColorFn) Underline() ColorFn {
	return c.color("4", "24")
}

func (c ColorFn) Reverse() ColorFn {
	return c.color("7", "27")
}

func (c ColorFn) Italic() ColorFn {
	return c.color("3", "23")
}

func (c ColorFn) Invisible() ColorFn {
	return c.color("8", "28")
}

func (c ColorFn) StrikeThrough() ColorFn {
	return c.color("9", "29")
}

// See: https://en.wikipedia.org/wiki/ANSI_escape_code#CSI_codes
func (c ColorFn) Other(start string, end string) ColorFn {
	return c.color(start, end)
}
