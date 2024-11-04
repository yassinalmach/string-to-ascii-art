package utils

import "strings"

func checkFlags(arg string, flags *Flags, lenArgs int) {
	flags.flag = false

	// output flag
	if strings.HasPrefix(arg, "--output=") {
		flags.flag = true
		flags.outputFile = arg[9:]
		if !strings.HasSuffix(arg, ".txt") || flags.outputFile == ".txt" {
			PrintError("output")
		}
	}

	// Align flag
	if strings.HasPrefix(arg, "--align=") {
		flags.flag = true
		flags.alignOption = arg[8:]
		option := []string{"center", "right", "left", "justify"}
		if !contains(flags.alignOption, option) {
			PrintError("justify")
		}
	}

	// color flag
	if strings.HasPrefix(arg, "--color=") {
		flags.flag = true
		flags.color = arg[8:]
		colors := []string{"white", "black", "red", "pink", "blue", "green", "brown", "orange", "yellow", "purple"}
		if !contains(flags.color, colors) {
			PrintError("color")
		}
	}

	// check if there is an conflict
	if (lenArgs == 5 && flags.color == "") || (lenArgs == 2 && flags.flag) || (lenArgs == 4 && !flags.flag) {
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
