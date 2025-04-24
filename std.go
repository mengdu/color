package color

var Std = New(EnableColor(false))

func Black() ColorFn {
	return Std.Black()
}

func Red() ColorFn {
	return Std.Red()
}

func Green() ColorFn {
	return Std.Green()
}

func Blue() ColorFn {
	return Std.Blue()
}

func Yellow() ColorFn {
	return Std.Yellow()
}

func Magenta() ColorFn {
	return Std.Magenta()
}

func Cyan() ColorFn {
	return Std.Cyan()
}

func White() ColorFn {
	return Std.White()
}

func Gray() ColorFn {
	return Std.Gray()
}

func RedBright() ColorFn {
	return Std.RedBright()
}

func GreenBright() ColorFn {
	return Std.GreenBright()
}

func BlueBright() ColorFn {
	return Std.BlueBright()
}

func YellowBright() ColorFn {
	return Std.YellowBright()
}

func MagentaBright() ColorFn {
	return Std.MagentaBright()
}

func CyanBright() ColorFn {
	return Std.CyanBright()
}

func WhiteBright() ColorFn {
	return Std.WhiteBright()
}

func BgBlack() ColorFn {
	return Std.BgBlack()
}

func BgRed() ColorFn {
	return Std.BgRed()
}

func BgGreen() ColorFn {
	return Std.BgGreen()
}

func BgBlue() ColorFn {
	return Std.BgBlue()
}

func BgYellow() ColorFn {
	return Std.BgYellow()
}

func BgMagenta() ColorFn {
	return Std.BgMagenta()
}

func BgCyan() ColorFn {
	return Std.BgCyan()
}

func BgWhite() ColorFn {
	return Std.BgWhite()
}

func BgGray() ColorFn {
	return Std.BgGray()
}

func BgRedBright() ColorFn {
	return Std.BgRedBright()
}

func BgGreenBright() ColorFn {
	return Std.BgGreenBright()
}

func BgBlueBright() ColorFn {
	return Std.BgBlueBright()
}

func BgYellowBright() ColorFn {
	return Std.BgYellowBright()
}

func BgMagentaBright() ColorFn {
	return Std.BgMagentaBright()
}

func BgCyanBright() ColorFn {
	return Std.BgCyanBright()
}

func BgWhiteBright() ColorFn {
	return Std.BgWhiteBright()
}

func Dim() ColorFn {
	return Std.Dim()
}

func Bold() ColorFn {
	return Std.Bold()
}

func Underline() ColorFn {
	return Std.Underline()
}

func Reverse() ColorFn {
	return Std.Reverse()
}

func Italic() ColorFn {
	return Std.Italic()
}

func Invisible() ColorFn {
	return Std.Invisible()
}

func StrikeThrough() ColorFn {
	return Std.StrikeThrough()
}

func Other(start string, end string) ColorFn {
	return Std.Other(start, end)
}
