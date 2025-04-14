package main

import (
	"fmt"

	"github.com/mengdu/color"
)

func cell(s string) string {
	return fmt.Sprintf("%-15s", s)
}

func main() {
	fmt.Printf("%s, %s\n", color.Std.Bold().String("Hi"), color.Std.Italic().Blue().String("Go Color!"))
	c := color.New(color.EnableColor(true))
	fmt.Printf("%s %s %s\n", c.Black().String(cell("Black")), c.Gray().String(cell("Gray")), c.BgBlack().String(cell("BgBlack")))
	fmt.Printf("%s %s %s\n", c.Red().String(cell("Red")), c.RedBright().String(cell("RedBright")), c.BgRed().String(cell("BgRed")))
	fmt.Printf("%s %s %s\n", c.Green().String(cell("Green")), c.GreenBright().String(cell("GreenBright")), c.BgGreen().String(cell("BgGreen")))
	fmt.Printf("%s %s %s\n", c.Blue().String(cell("Blue")), c.BlueBright().String(cell("BlueBright")), c.BgBlue().String(cell("BgBlue")))
	fmt.Printf("%s %s %s\n", c.Yellow().String(cell("Yellow")), c.YellowBright().String(cell("YellowBright")), c.BgYellow().String(cell("BgYellow")))
	fmt.Printf("%s %s %s\n", c.Magenta().String(cell("Magenta")), c.MagentaBright().String(cell("MagentaBright")), c.BgMagenta().String(cell("BgMagenta")))
	fmt.Printf("%s %s %s\n", c.Cyan().String(cell("Cyan")), c.CyanBright().String(cell("CyanBright")), c.BgCyan().String(cell("BgCyan")))
	fmt.Printf("%s %s %s\n", c.White().String(cell("White")), c.WhiteBright().String(cell("WhiteBright")), c.BgWhite().String(cell("BgWhite")))

	fmt.Println(
		c.Bold().String("Bold"),
		c.Dim().String("Dim"),
		c.Italic().String("Italic"),
		c.Underline().String("Underline"),
		c.Reverse().String("Reverse"),
		c.StrikeThrough().String("StrikeThrough"),
	)
	fmt.Println(c.StrikeThrough().Underline().Bold().Italic().Red().String("StrikeThrough Underline Bold Italic Green"))
	fmt.Println(c.Other("38;2;117;45;245", "0").String("Other color"))
	fmt.Println(c.Other("48;2;117;45;245", "0").String("Other color"))
}
