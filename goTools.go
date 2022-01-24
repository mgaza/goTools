/*
	These functions are for generally coding practices using the
	Go Language. All of these are used frequently in our data engineering
	day to day.

	Author: Mark
*/

package goTools

import (
	"encoding/csv"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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

// The correct format to be read in for this function is:
//       countyname_dddd-dd-dd_dddd-dd-dd.csv - d for digit (start and end dates)
//
// Should return dddd-dd-dd_dddd-dd-dd to variable
func GetExportYearMonth(fullPath string) string {
	fullPathSplit := strings.Split(fullPath, "\\")
	return fullPathSplit[len(fullPathSplit)-1][strings.Index(fullPathSplit[len(fullPathSplit)-1], "_")+1 : strings.LastIndex(fullPathSplit[len(fullPathSplit)-1], ".")]
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

// Creates, Opens, and writes to a new file for csv readers
func OpenAndWriteCSVFile(file string, outdirectory string, content [][]string) {
	writefilepath := outdirectory + "\\" + file
	writefile, err := os.Create(writefilepath)
	CheckErrorFatal("Could not create writefile: ", err)
	defer CloseFile(writefile)

	w := csv.NewWriter(writefile)
	// defer w.Flush()

	w.WriteAll(content)
}

// Call to close and Open File Path
func CloseFile(f *os.File) {
	err := f.Close()
	CheckErrorFatal("could not close file: ", err)
}
