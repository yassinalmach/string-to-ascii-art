package utils

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

// checkFlags checks if there is a valid flag in the first argument
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
		reg := regexp.MustCompile(`^rgb\((\d+),\s?(\d+),\s?(\d+)\)$`)
		if reg.MatchString(flags.Color) {
			flags.RgbColor = true
			submatched := reg.FindStringSubmatch(flags.Color)
			for i := 1; i < len(submatched); i++ {
				nbr, err := strconv.Atoi(submatched[i])
				if err != nil {
					log.Fatalln("error converting RGB number")
				}
				if nbr > 255 || nbr < 0 {
					log.Fatalln("error RGB number is incorrect")
				}
			}
			flags.RGB.R = submatched[1]
			flags.RGB.G = submatched[2]
			flags.RGB.B = submatched[3]
		} else {
			colors := []string{"white", "black", "red", "pink", "blue", "green", "brown", "orange", "yellow", "purple"}
			if !contains(flags.Color, colors) {
				PrintError("color")
			}
		}
	}

	// reverse flag
	if strings.HasPrefix(arg, "--reverse=") {
        flags.ReverseFile = arg[10:]
        if !strings.HasSuffix(flags.ReverseFile, ".txt") || flags.ReverseFile == ".txt" {
            PrintError("reverse")
        }
    }

	// check if there is an conflict between args
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
