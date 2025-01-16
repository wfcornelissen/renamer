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
	files := walker.WalkDir(models.MacPath)
	if len(files) == 0 {
		fmt.Println("No files found")
		return renamedFiles
	}
	fmt.Println("How many times do you want to repeat?")
	repeat := cmd.GetInput()
	companySelection, startingPOD, startingLC, startingINV := cmd.GetAllInputs()
	var docName string
	for i := 0; i < repeat; i++ {
		for _, file := range files {
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
		}
	}
	return renamedFiles
}
