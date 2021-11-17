package main

import (
	"encoding/csv"
	"flag"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
)

var (
	CsvInput  string
	UID       string
	CsvOutput string
)

// struct for getting just user ID out of csv file
type DiscordCSV struct {
	UserID      string
	Author      string
	Date        string
	Content     string
	Attachments string
}

func init() {

	flag.StringVar(&UID, "id", "", "User ID to locate")
	flag.StringVar(&CsvInput, "i", "", "The csv or folder to search")
	flag.StringVar(&CsvOutput, "o", "", "Output CSV to save mesages from user")
	flag.Parse()
}

func main() {
	// Opens input and gets info
	fileInfo, err := os.Stat(CsvInput)
	if err != nil {
		// error handling
		println(err.Error())
	}

	if CsvOutput != "" {
		OutFile, err := os.OpenFile(CsvOutput, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		OutFile.WriteString("UID, Author, Date, Content, Attachments, Found in\n")
		OutFile.Close()
	}

	// checks if input is directory or file
	if fileInfo.IsDir() {
		// file is a directory
		files, err := ioutil.ReadDir(CsvInput) // CsvInput
		if err != nil {
			log.Fatal(err)
		}

		// For each file in folder
		for _, file := range files {
			if !file.IsDir() {
				file_path := CsvInput + file.Name()
				process(file_path, file)
			}

		}

	} else {
		// file is not a directory
		files, err := os.Stat(CsvInput) // CsvInput
		if err != nil {
			log.Fatal(err)
		}
		process(CsvInput, files)

	}
}

func process(inFile string, fileInfo fs.FileInfo) {
	var messages []string

	// Read CSV file
	lines, err := ReadCsv(inFile)
	if err != nil {
		panic(err)
	}

	for _, line := range lines {
		data := DiscordCSV{
			UserID:      line[0],
			Author:      line[1],
			Date:        line[2],
			Content:     line[3],
			Attachments: line[4],
		}
		if UID == data.UserID {
			messages = append(messages, data.UserID+","+data.Author+","+data.Date+","+data.Content+","+data.Attachments+","+fileInfo.Name())
			println(data.UserID + "," + data.Author + "," + data.Date + "," + data.Content + "," + data.Attachments + "," + fileInfo.Name())
		}
	}

	if CsvOutput != "" {
		OutFile, err := os.OpenFile(CsvOutput, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			panic(err)
		}
		for _, message := range messages {
			OutFile.WriteString(message + "\n")
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

	// Read File into a Variable
	lines, err := csv.NewReader(f).ReadAll()
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
