package utils

import (
	"os"
	"strings"
)

type Flags struct {
	Str         string
	BannerFile  string
	Flag        bool
	OutputFile  string
	Color       string
	ColoredStr  string
	AlignOption string
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
		flags.BannerFile = banner(args[4])
		flags.ColoredStr = args[2]
	} else if len(args) == 4 && flags.Color != "" {
		if args[3] == "standard" || args[3] == "shadow" || args[3] == "thinkertoy" {
			flags.BannerFile = banner(args[3])
		} else {
			flags.BannerFile = banner("standard")
			flags.ColoredStr = args[2]
		}
	} else if len(args) == 4 {
		flags.BannerFile = banner(args[3])
	} else if len(args) == 3 && !flags.Flag {
		flags.BannerFile = banner(args[2])
	} else {
		flags.BannerFile = banner("standard")
	}

	// pick string position
	if flags.ColoredStr != "" {
		flags.Str = args[3]
	} else if flags.Flag {
		flags.Str = args[2]
	} else {
		flags.Str = args[1]
	}

	// Empty String
	if flags.Str == "" {
		os.Exit(0)
	}
	return flags

}