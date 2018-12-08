package main

import (
	"github.com/fatih/color"
	"time"
)

func main() {
	// Print with default helper functions
	color.Cyan("Prints text in cyan.")

	// A newline will be appended automatically
	color.Blue("Prints %s in blue.", "text")

	// These are using the default foreground colors
	color.Red("We have red")
	color.Magenta("And many others ..")
	time.Sleep(time.Hour)
}