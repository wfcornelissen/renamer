package manual

import (
	"fmt"
	"path"
	"strings"

	"github.com/wfcornelissen/renamer/internal/builder"
	"github.com/wfcornelissen/renamer/internal/cmd"
	"github.com/wfcornelissen/renamer/internal/models"
	"github.com/wfcornelissen/renamer/internal/walker"
)

func ManualEntry() []string {
	fmt.Println("Manual Entry")
	renamedFiles := []string{}
	files := walker.WalkDir(models.MacPath)
	if len(files) == 0 {
		fmt.Println("No files found")
		return renamedFiles
	}
	for _, file := range files {
		if strings.HasSuffix(file, "/") || strings.HasPrefix(path.Base(file), ".") {
			continue
		}
		fmt.Println(file)
		if cmd.Skip() {
			continue
		}
		companySelection, PODNumber, lcOrExternalPOD, invoiceNumber := cmd.GetAllInputs()
		docName := builder.BuildString(models.Companies[companySelection], PODNumber, lcOrExternalPOD, invoiceNumber)
		fmt.Println(docName) //debug
		renamedFiles = append(renamedFiles, docName)
	}
	return renamedFiles
}
