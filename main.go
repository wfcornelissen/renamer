package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Hello, World!")
	/*files := []string{}
	filepath.WalkDir("/mnt/f/Scans", func(path string, d fs.DirEntry, err error) error {
		files = append(files, path)
		return nil
	})
	fmt.Println(files)*/
	for i, company := range companies {
		fmt.Println(i, company.name)
	}

	scan := bufio.NewScanner(os.Stdin)
	scan.Scan()
	input, err := strconv.Atoi(scan.Text())
	if err != nil {
		fmt.Println("Invalid input")
		os.Exit(1)
	}
	fmt.Println(companies[input].name)

	os.Exit(0)
}
