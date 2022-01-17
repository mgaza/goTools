package goTools

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
)

func CheckErrorFatal(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

func filePathWalker(filePath string, extfile string) []string {
	re := regexp.MustCompile(extfile)
	var paths []string

	err := filepath.WalkDir(filePath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		if re.MatchString(d.Name()) {
			paths = append(paths, path)
		}

		return nil
	})

	CheckErrorFatal("error walking the path: ", err)
	return paths
}
