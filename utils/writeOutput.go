package utils

import (
	"fmt"
	"os"
)

// this Function choose between printing the output in the terminal or create the file output
func WriteOutput(result, outputFile string) error {
	if outputFile != "" {
		file, err := os.Create(outputFile)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(result)
		if err != nil {
			return err
		}
	} else {
		fmt.Print(result)
	}
	return nil
}
