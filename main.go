package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func main() {
	files := []string{}
	filepath.WalkDir("/mnt/f/Scans", func(path string, d fs.DirEntry, err error) error {
		files = append(files, path)
		return nil
	})

	for _, file := range files {
		fmt.Println(file)
	}
	// Print all company names with numbers
	for i, company := range companies {
		fmt.Println(i, company.name)
	}

	// Get all inputs
	companySelection, startingNumber, lcOrExternalPOD, invoiceNumber := getAllInputs()

	// Create the document name
	docName := fmt.Sprintf("%sPOD%d_LC%d_INV%d", companies[companySelection].fileStart, startingNumber, lcOrExternalPOD, invoiceNumber)
	fmt.Println(docName) //debug

	// Confirm the document name - debugging
	if userConfirmation(docName) {
		fmt.Println("Confirmed")
	} else {
		os.Exit(1)
	}

	os.Exit(0)
}

func getAllInputs() (int, int, int, int) {
	fmt.Println("Please enter the company number: ")
	companySelection := getInput()
	fmt.Println("Please enter internal POD number: ")
	startingNumberSelection := getInput()
	fmt.Println("Please enter the Loadcon or external POD number: ")
	lcOrExternalPOD := getInput()
	fmt.Println("Please enter the invoice number: ")
	invoiceNumber := getInput()
	return companySelection, startingNumberSelection, lcOrExternalPOD, invoiceNumber
}

func getInput() int {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input, err := strconv.Atoi(scan.Text())
	if err != nil {
		fmt.Println("Invalid input: Unknown command")
		os.Exit(1)
	}
	return input
}

func userConfirmation(userInput string) bool {
	fmt.Println("Please confirm the input: ", userInput)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := scan.Text()
	if strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
		return true
	}
	return false
}
