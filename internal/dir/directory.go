package dir

import (
	"log"
	"os"
	"path/filepath"
)

func GetFilePaths(root string) ([]string, error) {
	var filepaths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			filepaths = append(filepaths, path)
			if err != nil {
				log.Println("Error opening file:", err)
			}
		}

		return nil
	})

	if err != nil {
		log.Println("Error:", err)
	}

	return filepaths, err
}
