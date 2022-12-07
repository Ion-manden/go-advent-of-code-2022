package day7

import (
	"fmt"
	"log"

	"golang.org/x/exp/slices"
)

func getFolderSize(folderPath string) int {
	idx := slices.IndexFunc(folders, func(f folder) bool {
		return f.path == folderPath
	})

	if idx < 0 {
		log.Fatal("folder does not exist: ", folderPath)
	}

	folder := folders[idx]

	totalSize := 0
	// Add size of all files in the folder
	for _, f := range files {
		if f.folderPath != folder.path {
			continue
		}

		totalSize += f.size
	}

	// Find size of all files in sub folders
	for _, f := range folders {
		subFolderPath := fmt.Sprint(folder.path, "/", f.name)
		if folderPath == "/" {
			subFolderPath = fmt.Sprint(folder.path, f.name)
		}

		if f.path != subFolderPath {
			continue
		}

		totalSize += getFolderSize(f.path)
	}

	return totalSize
}
