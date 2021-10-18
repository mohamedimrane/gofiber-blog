package utils

import "github.com/fatih/color"

func PrintRed(text string) {
	color.New(color.BgRed).Add(color.FgWhite).Add(color.Bold).Println(text)
}

func PrintGreen(text string) {
	color.New(color.BgGreen).Add(color.FgWhite).Add(color.Bold).Println(text)
}

func PrintYellow(text string) {
	color.New(color.BgYellow).Add(color.FgBlack).Add(color.Bold).Println(text)
}
