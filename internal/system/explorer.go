package system

import (
	"os"
	"path/filepath"
	"strings"
)

type testable struct {
	fileName         string
	testFile         string
	belongsToPackage string
}

func Explore(path string) ([]string, error) {
	var goFiles = make(map[string]testable)
	var unTested []string
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		var currentDir string
		if err != nil {
			return err
		}
		if info.IsDir() {
			currentDir = path
			if currentDir == "vendor" {
				return filepath.SkipDir
			}
		}
		if strings.Contains(path, "_test.go") {
			s := strings.Split(path, "_test.go")
			if _, exists := goFiles[s[0]]; exists {
				//TODO: sort this with pointers.
				tested := goFiles[s[0]]
				tested.testFile = path
				goFiles[s[0]] = tested
			} else {
				goFiles[s[0]] = testable{
					testFile:         info.Name(),
					belongsToPackage: currentDir,
				}
			}
		} else if strings.Contains(path, ".go") {
			s := strings.Split(path, ".go")
			if _, exists := goFiles[s[0]]; exists {
				//TODO: sort this with pointers.
				file := goFiles[s[0]]
				file.fileName = path
				goFiles[s[0]] = file
			} else {
				goFiles[s[0]] = testable{
					fileName:         info.Name(),
					belongsToPackage: currentDir,
				}
			}

		}

		return err
	})

	if err != nil {
		return nil, err
	}
	for _, v := range goFiles {
		if v.testFile == "" {
			unTested = append(unTested, v.fileName)
		}
	}

	return unTested, nil
}
