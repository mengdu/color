package color

import (
	"fmt"
	"strings"
)

func stringReplaceAll(str, substring, replacer string) string {
	if !strings.Contains(str, substring) {
		return str
	}

	var result strings.Builder
	subLen := len(substring)
	start := 0

	for {
		index := strings.Index(str[start:], substring)
		if index == -1 {
			break
		}

		index += start
		result.WriteString(str[start:index])
		result.WriteString(substring)
		result.WriteString(replacer)

		start = index + subLen
	}

	result.WriteString(str[start:])
	return result.String()
}

func New(enable bool) ColorFn {
	return func(s string) (open []string, clase []string, out string, color bool) {
		return []string{}, []string{}, s, enable
	}
}

type ColorFn func(s string) (open []string, clase []string, out string, enableColor bool)

func (f ColorFn) color(start, end string) ColorFn {
	return func(s string) (open []string, clase []string, out string, color bool) {
		a, b, c, e := f(s)
		a = append(a, start)
		b = append(b, end)
		return a, b, c, e
	}
}

func (f ColorFn) String(s string) string {
	open, close, v, enable := f(s)
	if !enable {
		return v
	}

	strs := make([]string, len(open)+len(close)+1)
	l := len(strs)
	vi := l / 2
	strs[vi] = v
	for i := range open {
		strs[i] = fmt.Sprintf("\x1b[%sm", open[i])
		strs[l-i-1] = fmt.Sprintf("\x1b[%sm", close[i])
	}

	// support nested color
	for i := len(open) - 1; i >= 0; i-- {
		if strings.Contains(v, "\x1b") {
			strs[vi] = stringReplaceAll(strs[vi], strs[l-i-1], strs[i])
		}
	}
	return strings.Join(strs, "")
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
