package walker

import (
	"io/fs"
	"path/filepath"
)

func WalkDir(path string) []string {
	files := []string{}
	filepath.WalkDir(path, func(path string, d fs.DirEntry, err error) error {
		files = append(files, path)
		return nil
	})
	return files
}
