package dev

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/wfcornelissen/renamer/internal/models"
)

func SpoofFiles() {
	fmt.Println("Spoofing files")
	for i := 1; i <= 2; i++ {
		filename := fmt.Sprintf("test%d.txt", i)
		content := fmt.Sprintf("test%d", i)
		os.WriteFile(filepath.Join(models.MacPath, filename), []byte(content), 0644)
	}
}
