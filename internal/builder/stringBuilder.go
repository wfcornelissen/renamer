package builder

import (
	"fmt"
	"strconv"

	"github.com/wfcornelissen/renamer/internal/models"
)

func BuildString(company models.Company, PODNumber int, lcOrExternalPOD int, invoiceNumber int) string {
	fmt.Println("Building string")
	var POD string
	var LC string
	var inv string

	if PODNumber == 0 {
		POD = ""
	} else {
		POD = "POD" + strconv.Itoa(PODNumber) + "_"
	}

	if lcOrExternalPOD == 0 {
		LC = ""
	} else {
		LC = "LC" + strconv.Itoa(lcOrExternalPOD) + "_"
	}

	if invoiceNumber == 0 {
		inv = ""
	} else {
		inv = "INV" + strconv.Itoa(invoiceNumber)
	}

	fileName := company.FileStart + POD + LC + inv
	return fileName
}
