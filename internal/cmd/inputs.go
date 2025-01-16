package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/wfcornelissen/renamer/internal/models"
)

func GetAllInputs() (int, int, int, int) {
	for i, company := range models.Companies {
		fmt.Println(i, company.Name)
	}
	fmt.Println("Please enter the company number: ")
	companySelection := GetInput()
	fmt.Println("Please enter internal POD number: ")
	startingNumberSelection := GetInput()
	fmt.Println("Please enter the Loadcon or external POD number: ")
	lcOrExternalPOD := GetInput()
	fmt.Println("Please enter the invoice number: ")
	invoiceNumber := GetInput()
	return companySelection, startingNumberSelection, lcOrExternalPOD, invoiceNumber
}

func GetInput() int {
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	if scan.Text() == "" {
		return 0
	}
	input, err := strconv.Atoi(scan.Text())
	if err != nil {
		fmt.Println("Invalid input: Unknown command")
		os.Exit(1)
	}

	return input
}

func UserConfirmation(userInput string) bool {
	fmt.Println("Please confirm the input: ", userInput)
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := scan.Text()
	if strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
		return true
	}
	return false
}

func Skip() bool {
	fmt.Println("Do you want to skip this file? (y/n)")
	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input := scan.Text()
	if strings.ToLower(input) == "y" || strings.ToLower(input) == "yes" {
		return true
	}
	return false
}
