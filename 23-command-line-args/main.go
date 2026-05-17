package main

import (
	"fmt"
	"os"
)

/*
For professional-grade Go CLI apps, the industry standard is a third-party framework called Cobra
usually paired with Viper for configuration management.
Why developers use Cobra:

Subcommands: Easily create complex command trees (app server start, app user add).

POSIX Compliance: Supports standard double-dash long flags (--port) and single-dash short flags (-p).

Hooks: Run code before or after commands execute.

Auto-completion: Can automatically generate shell completion scripts for Bash, Zsh, and PowerShell.
*/

func main() {
	args := os.Args
	fmt.Println(len(args))
	for i, arg := range args {
		fmt.Printf("Argument %d: %s\n", i, arg)
	}

	// Argument 0: path\to\main.exe
}
