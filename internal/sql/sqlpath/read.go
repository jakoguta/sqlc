package sqlpath

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/kyleconroy/sqlc/internal/migrations"
)

// Glob return all valid SQL files found in the listed paths.
//
// Valid files include files ending in {.sql}. It omits hidden files, directories and migrations.
func Glob(paths []string) ([]string, error) {
	var files []string
	for _, path := range paths {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return nil, fmt.Errorf("path %s does not exist", path)
		}
		if fileInfo.IsDir() {
			dirFiles, err := os.ReadDir(path)
			if err != nil {
				return nil, err
			}
			for _, file := range dirFiles {
				newFilePath := filepath.Join(path, file.Name())
				if isValidSqlFile(newFilePath) {
					files = append(files, newFilePath)
				}
			}
			continue
		}
		if isValidSqlFile(path) {
			files = append(files, path)
		}
	}
	return files, nil
}

func isValidSqlFile(path string) bool {
	if !strings.HasSuffix(path, ".sql") {
		return false
	}
	if strings.HasPrefix(filepath.Base(path), ".") {
		return false
	}
	if migrations.IsDown(filepath.Base(path)) {
		return false
	}
	return true
}
