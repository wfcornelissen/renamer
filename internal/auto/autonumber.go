package auto

import (
	"fmt"
	"path"
	"strings"

	"github.com/wfcornelissen/renamer/internal/builder"
	"github.com/wfcornelissen/renamer/internal/cmd"
	"github.com/wfcornelissen/renamer/internal/models"
	"github.com/wfcornelissen/renamer/internal/walker"
)

func AutoEntry() []string {
	fmt.Println("Auto Entry")
	renamedFiles := []string{}
	files := walker.WalkDir(models.WindowsPath)
	if len(files) == 0 {
		fmt.Println("No files found")
		return renamedFiles
	}
	fmt.Println("How many times do you want to repeat?")
	repeat := cmd.GetInput()
	fmt.Println(repeat) //debug
	companySelection, startingPOD, startingLC, startingINV := cmd.GetAllInputs()
	var docName string
	var i int
	for _, file := range files {
		if i >= repeat {
			break
		}
		if strings.HasSuffix(file, "/") || strings.HasPrefix(path.Base(file), ".") {
			continue
		}
		docName = builder.BuildString(models.Companies[companySelection], startingPOD, startingLC, startingINV)
		renamedFiles = append(renamedFiles, docName)
		if startingPOD > 0 {
			startingPOD++
		}
		if startingLC > 0 {
			startingLC++
		}
		if startingINV > 0 {
			startingINV++
		}
		i++
	}
	return renamedFiles
}
