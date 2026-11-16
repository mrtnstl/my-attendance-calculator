// attendance calculator
package main

import (
	"fmt"
	"os"

	"example.com/munkaidoszamlalo/cmd"
)

const (
	ANSI_TERM       = "\x1b[0m"
	ANSI_BOLD_RED   = "\x1b[1;31m"
	ANSI_BOLD_GREEN = "\x1b[1;32m"
)

func main() {
	if len(os.Args) < 2 {
		ErrExit(fmt.Sprintf("err: %s\n", "insufficient amount of arguments"))
	}

	filePath := os.Args[1]
	if filePath == "" {
		ErrExit(fmt.Sprintf("err: %s\n", "empty file path"))
	}

	fContent, err := cmd.ReadFile(filePath)
	if err != nil {
		ErrExit(err.Error())
	}

	dayRecords, err := cmd.ParseFile(fContent)
	if err != nil {
		ErrExit(err.Error())
	}

	sum := int64(0)
	for _, pr := range dayRecords {
		sum += pr.Fraction
	}

	fmt.Print("attendance:\n")
	if sum >= 0 {
		fmt.Printf("%syou got %d minutes. %s%s\n", ANSI_BOLD_GREEN, sum, "good to go!", ANSI_TERM)
		return
	}
	fmt.Printf("%syou got %d minutes. %s%s\n", ANSI_BOLD_RED, sum, "slave on...", ANSI_TERM)
}

// prints error message, then exits the application with code 1
func ErrExit(message string) {
	fmt.Printf("%s", message)
	os.Exit(1)
}
