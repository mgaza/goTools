/*
	These functions are for generally coding practices using the
	Go Language. All of these are used frequently in our data engineering
	day to day.

	Author: Mark
*/

package goTools

import (
	"fmt"
	"io/fs"
	"log"
	"path/filepath"
	"regexp"
)

// Use to check for errors that should
// not logically allow the program to continue
func CheckErrorFatal(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}

// Use for testing errors that could allow the
// program to pass or during testing stages
func CheckErrorNonFatal(message string, err error) {
	if err != nil {
		fmt.Println(message, err)
	}
}

// Allows for a slice return of file searches
// down whatever path the user wants to go down
func FilePathWalker(filePath string, extfile string) []string {
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
