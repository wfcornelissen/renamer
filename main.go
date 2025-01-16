package main

import (
	"fmt"
	"os"

	"github.com/wfcornelissen/renamer/internal/auto"
	"github.com/wfcornelissen/renamer/internal/cmd"
	"github.com/wfcornelissen/renamer/internal/dev"
	"github.com/wfcornelissen/renamer/internal/manual"
	"github.com/wfcornelissen/renamer/internal/models"
)

func main() {
	dev.SpoofFiles()
	fmt.Println("Please enter the entry method: ")
	for i, method := range models.EntryMethods {
		fmt.Println(i, method)
	}
	entryMethod := cmd.GetInput()
	var newNames []string
	if entryMethod == 1 {
		newNames = manual.ManualEntry()
	} else if entryMethod == 2 {
		newNames = auto.AutoEntry()
	} else {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	fmt.Println(newNames)

	os.Exit(0)
}
