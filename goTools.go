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
// Should return dddddddd_dddddddd to variable
func GetExportYearMonth(fullPath string) (string, string) {
	regVar := `(?P<county>[a-z]+|[a-z]+_[a-z]+)_(?P<startyear>\d{4})-(?P<startmonth>\d{2})-(?P<startday>\d{2})_(?P<endyear>\d{4})-(?P<endmonth>\d{2})-(?P<endday>\d{2})\.csv`
	re := regexp.MustCompile(regVar)
	matched, err := regexp.MatchString(regVar, fullPath)

	CheckErrorFatal("There's a problem with the date: ", err)

	if matched {
		matches := re.FindStringSubmatch(fullPath)
		countyName := matches[1]
		newDateFileName := matches[2] + matches[3] + matches[4] + "_" + matches[5] + matches[6] + matches[7]

		return countyName, newDateFileName
	} else {
		return "county", "nodate"
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
