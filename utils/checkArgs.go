package utils

import (
	"os"
	"strings"
)

type Flags struct {
	str         string
	bannerFile  string
	flag        bool
	outputFile  string
	color       string
	coloredStr  string
	alignOption string
}

func CheckArgs(args []string) (Flags) {
	var flags Flags
	if len(args) < 2 || len(args) > 5 {
		PrintError("color")
	} else if len(args) == 2 && strings.HasPrefix(args[1], "--help") {
		PrintError("color")
	} else {
		checkFlags(args[1], &flags, len(args))
	}

	// pick banner file and letters to be colored
	if len(args) == 5 {
		flags.bannerFile = banner(args[4])
		flags.coloredStr = args[2]
	} else if len(args) == 4 && flags.color != "" {
		if args[3] == "standard" || args[3] == "shadow" || args[3] == "thinkertoy" {
			flags.bannerFile = banner(args[3])
		} else {
			flags.bannerFile = banner("standard")
			flags.coloredStr = args[2]
		}
	} else if len(args) == 4 {
		flags.bannerFile = banner(args[3])
	} else if len(args) == 3 && !flags.flag {
		flags.bannerFile = banner(args[2])
	} else {
		flags.bannerFile = banner("standard")
	}

	// pick string position
	if flags.coloredStr != "" {
		flags.str = args[3]
	} else if flags.flag {
		flags.str = args[2]
	} else {
		flags.str = args[1]
	}

	// Empty String
	if flags.str == "" {
		os.Exit(0)
	}
	return flags

}
