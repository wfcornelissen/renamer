package manual

import (
	"fmt"

	"github.com/wfcornelissen/renamer/internal/cmd"
	"github.com/wfcornelissen/renamer/internal/models"
	"github.com/wfcornelissen/renamer/internal/walker"
)

func ManualEntry() {
	fmt.Println("Manual Entry")
	renamedFiles := []string{}
	files := walker.WalkDir("/mnt/f/codeTesting")
	if len(files) == 0 {
		fmt.Println("No files found")
		return
	}
	for _, file := range files {
		fmt.Println(file)
		if cmd.Skip() {
			continue
		}
		companySelection, startingNumber, lcOrExternalPOD, invoiceNumber := cmd.GetAllInputs()
		docName := fmt.Sprintf("%sPOD%d_LC%d_INV%d", models.Companies[companySelection].FileStart, startingNumber, lcOrExternalPOD, invoiceNumber)
		fmt.Println(docName) //debug
		renamedFiles = append(renamedFiles, docName)
	}
}
