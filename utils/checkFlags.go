package utils

import "strings"

func checkFlags(arg string, flags *Flags, lenArgs int) {
	flags.Flag = false

	// output Flag
	if strings.HasPrefix(arg, "--output=") {
		flags.Flag = true
		flags.OutputFile = arg[9:]
		if !strings.HasSuffix(arg, ".txt") || flags.OutputFile == ".txt" {
			PrintError("output")
		}
	}

	// Align Flag
	if strings.HasPrefix(arg, "--align=") {
		flags.Flag = true
		flags.AlignOption = arg[8:]
		option := []string{"center", "right", "left", "justify"}
		if !contains(flags.AlignOption, option) {
			PrintError("justify")
		}
	}

	// color Flag
	if strings.HasPrefix(arg, "--color=") {
		flags.Flag = true
		flags.Color = arg[8:]
		colors := []string{"white", "black", "red", "pink", "blue", "green", "brown", "orange", "yellow", "purple"}
		if !contains(flags.Color, colors) {
			PrintError("color")
		}
	}

	// check if there is an conflict
	if (lenArgs == 5 && flags.Color == "") || (lenArgs == 2 && flags.Flag) || (lenArgs == 4 && !flags.Flag) {
		PrintError("color")
	}
}

// contains checks if a given string matches any string in a slice of options.
func contains(str string, option []string) bool {
	for _, sub := range option {
		if str == sub {
			return true
		}
	}
	return false
}
