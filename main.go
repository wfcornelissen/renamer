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
	if entryMethod == 1 {
		manual.ManualEntry()
	} else if entryMethod == 2 {
		auto.AutoEntry()
	} else {
		fmt.Println("Invalid input")
		os.Exit(1)
	}

	os.Exit(0)
}
