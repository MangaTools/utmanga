package executor

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var acceptedImageExtensions = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}

func getImagesPaths(folder string, extensions map[string]bool, recursive bool) ([]string, error) {
	dirs := []string{folder}

	result := make([]string, 0)

	folderCut := folder + string(filepath.Separator)

	for len(dirs) != 0 {
		currentFolder := dirs[0]

		files, err := os.ReadDir(currentFolder)
		if err != nil {
			return nil, fmt.Errorf("read dir \"%s\" to find images: %w", currentFolder, err)
		}

		dirs = dirs[1:]

		for _, file := range files {
			if file.IsDir() {
				if recursive {
					dirs = append(dirs, filepath.Join(currentFolder, file.Name()))
				}

				continue
			}

			filePath := filepath.Join(currentFolder, file.Name())
			fileExt := filepath.Ext(filePath)
			if _, ok := acceptedImageExtensions[fileExt]; !ok {
				continue
			}

			relativePath, _ := strings.CutPrefix(filePath, folderCut)
			result = append(result, relativePath)
		}
	}

	return result, nil
}
