package main

import (
	"fmt"

	"github.com/mengdu/color"
)

func cell(s string) string {
	return fmt.Sprintf("%-15s", s)
}

func main() {
	fmt.Printf("%s%s\n", color.BgGray().String(" Github "), color.BgGreen().White().String(" github.com/mengdu/color "))
	// c := color.New(false)
	c := color.New(color.EnableColor(true))
	fmt.Printf("%s %s %s %s\n", c.Black().String(cell("Black")), c.Gray().String(cell("Gray")), c.BgBlack().String(cell("BgBlack")), c.BgGray().String(cell("BgGray")))
	fmt.Printf("%s %s %s %s\n", c.Red().String(cell("Red")), c.RedBright().String(cell("RedBright")), c.BgRed().String(cell("BgRed")), c.BgRedBright().String(cell("BgRedBright")))
	fmt.Printf("%s %s %s %s\n", c.Green().String(cell("Green")), c.GreenBright().String(cell("GreenBright")), c.BgGreen().String(cell("BgGreen")), c.BgGreenBright().String(cell("BgGreenBright")))
	fmt.Printf("%s %s %s %s\n", c.Blue().String(cell("Blue")), c.BlueBright().String(cell("BlueBright")), c.BgBlue().String(cell("BgBlue")), c.BgBlueBright().String(cell("BgBlueBright")))
	fmt.Printf("%s %s %s %s\n", c.Yellow().String(cell("Yellow")), c.YellowBright().String(cell("YellowBright")), c.BgYellow().String(cell("BgYellow")), c.BgYellowBright().String(cell("BgYellowBright")))
	fmt.Printf("%s %s %s %s\n", c.Magenta().String(cell("Magenta")), c.MagentaBright().String(cell("MagentaBright")), c.BgMagenta().String(cell("BgMagenta")), c.BgMagentaBright().String(cell("BgMagentaBright")))
	fmt.Printf("%s %s %s %s\n", c.Cyan().String(cell("Cyan")), c.CyanBright().String(cell("CyanBright")), c.BgCyan().String(cell("BgCyan")), c.BgCyanBright().String(cell("BgCyanBright")))
	fmt.Printf("%s %s %s %s\n", c.White().String(cell("White")), c.WhiteBright().String(cell("WhiteBright")), c.BgWhite().String(cell("BgWhite")), c.BgWhiteBright().String(cell("BgWhiteBright")))

	fmt.Println(
		c.Bold().String("Bold"),
		c.Dim().String("Dim"),
		c.Italic().String("Italic"),
		c.Underline().String("Underline"),
		c.Reverse().String("Reverse"),
		c.StrikeThrough().String("StrikeThrough"),
	)
	// support nested color
	fmt.Println(c.Italic().Underline().Red().String(fmt.Sprintf("red %s red %s red",
		c.Green().String("green"),
		c.Blue().String("blue"),
	)))
	fmt.Println(c.StrikeThrough().Underline().Bold().Italic().Red().String("StrikeThrough Underline Bold Italic Red"))
	fmt.Println(c.Other("38;2;117;45;245", "0").String("Other color"))
	fmt.Println(c.Other("48;2;117;45;245", "0").String("Other bg color"))
}
