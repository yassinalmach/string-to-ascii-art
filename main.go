package main

import (
	"AsciiArt/utils"
	"log"
	"os"
)

func main() {
	flags := utils.CheckArgs(os.Args)

	asciiMap, err := utils.MakeMap(flags.BannerFile)
	if err != nil {
		utils.PrintError("banner")
	}

	// Split the string string into lines
	splittedStr := utils.SplitString(flags.Str, flags.AlignOption)

	// Make the Ascii Art
	result := utils.GenerateArt(asciiMap, splittedStr, flags)

	// Print the Ascii Art
	err = utils.WriteOutput(result, flags.OutputFile)
	if err != nil {
		log.Fatalln(err)
	}
}
