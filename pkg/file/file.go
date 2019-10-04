package file

import (
	"encoding/csv"
	"log"
	"os"
)

type Csv struct {
	File   *os.File
	Writer *csv.Writer
}

func NewCsv(filename string) *Csv {
	file, err := os.Create(filename)
	checkError("Error to create file", err)
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return &Csv{
		File:   file,
		Writer: writer,
	}
}

func (c *Csv) WriteLine(line []string) {
	err := c.Writer.Write(line)
	checkError("Cannot write to file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
