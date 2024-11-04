package main

import (
	"AsciiArt/utils"
	"fmt"
	"os"
)

func main() {
	flags := utils.CheckArgs(os.Args)
	fmt.Println(flags)
}
