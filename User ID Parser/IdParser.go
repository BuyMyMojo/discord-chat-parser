package main

import (
	"bytes"
	"encoding/csv"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Variables used for command line parameters
var (
	CsvInput    string
	CsvOutput   string
	unique_old  []string
	unique_name []string
	unique_ids  []string
)

func init() {

	flag.StringVar(&CsvInput, "i", "", "Input csv or folder [required]")
	flag.StringVar(&CsvOutput, "o", "out.csv", "Output CSV")
	flag.Parse()
}

// struct for getting just user ID out of csv file
type DiscordCSV struct {
	UserID   string
	Username string
}

func main() {

	// Check if there is no input and exits if true
	if CsvInput == "" {
		println("No input given, use -h for help on how to use this program")
		os.Exit(0)
	}

	// Opens input and gets info
	fileInfo, err := os.Stat(CsvInput)
	if err != nil {
		// error handling
		println(err.Error())
	}

	// checks if input is directory or file
	if fileInfo.IsDir() {
		// file is a directory
		files, err := ioutil.ReadDir(CsvInput)
		if err != nil {
			log.Fatal(err)
		}

		// For each file in folder
		for _, file := range files {
			if !file.IsDir() {
				// file is not a directory so process it's full path
				// check if windows or linux
				if strings.Contains(CsvInput, "\\") {
					// windows
					process(CsvInput + "\\" + file.Name())
				} else {
					// linux
					process(CsvInput + "/" + file.Name())
				}

			}
		}

	} else {
		// file is not a directory
		process(CsvInput)

	}

}

func process(CsvPath string) {

	if _, err := os.Stat(CsvOutput); err == nil {
		// path/to/whatever exists

		// Read CSV file
		lines, err := ReadCsv(CsvOutput)
		if err != nil {
			return
		}

		// Loop through lines & add to DiscordIDs list
		for i, _ := range lines {

			data := DiscordCSV{
				UserID:   lines[i][0],
				Username: lines[i][1],
			}
			if !contains(unique_old, data.UserID) {
				unique_old = append(unique_old, data.UserID)
				unique_name = append(unique_name, data.Username)
			}
		}
		WriteCSV(CsvPath)

	} else if errors.Is(err, os.ErrNotExist) {
		// path/to/whatever does *not* exist
		OutFile, err := os.Create(CsvOutput)
		if err != nil {
			fmt.Println(err)
			return
		}
		OutFile.Close()

		WriteCSV(CsvPath)

	} else {
		// Schrodinger: file may or may not exist. See err for details.

		// Therefore, do *NOT* use !os.IsNotExist(err) to test for file existence
		println(err)
	}

}

func WriteCSV(InFile string) {
	// Create DiscordIDs list
	var DiscordIDs []string
	var DiscordNames []string

	// Read CSV file
	lines, err := ReadCsv(InFile)
	if err != nil {
		panic(err)
	}

	OutFile, err := os.OpenFile(CsvOutput, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	// Loop through lines & add to DiscordIDs list
	for i, _ := range lines {

		data := DiscordCSV{
			UserID:   lines[i][0],
			Username: lines[i][1],
		}
		if !contains(DiscordIDs, data.UserID) {
			DiscordIDs = append(DiscordIDs, data.UserID)
			DiscordNames = append(DiscordNames, data.Username)
		}
	}

	// Checks if ID is already in list
	// if it isnt then it gets writen to list
	for i, ID := range DiscordIDs {
		if !contains(unique_old, ID) {
			OutFile.WriteString(ID + "," + DiscordNames[i] + "\n")
		}
	}
}

func ReadCsv(filename string) ([][]string, error) {

	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return [][]string{}, err
	}
	defer f.Close()

	var buf bytes.Buffer
	io.Copy(&buf, f)

	x := string(buf.Bytes())

	strings.ReplaceAll(x, "\r\n", "\n")

	strings.ReplaceAll(x, "\r\n", " ")

	reader := csv.NewReader(strings.NewReader(x))

	reader.ReuseRecord = true

	// Read File into a Variable
	lines, err := reader.ReadAll()
	if err != nil {
		return [][]string{}, err
	}

	return lines, nil
}

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}
