package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/amodemoli/zc/commands"
	"github.com/amodemoli/zc/menus"
)

func main() {

	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "info":
			commands.Info()
			return
		case "back":
			commands.Back()
			return
		}
	}

	help := flag.Bool("help", false, "Show help message")
	version := flag.Bool("version", false, "Show correct version")
	flag.Parse()

	if *help {
		menus.HelpMenu()
		return
	}
	if *version {
		showVersion()
		return
	}

	if flag.NFlag() == 0 {
		menus.Default()
	}
}

func showVersion() {
	fmt.Println("Zima Command (zc) - Version 1.12@lastles")
}
