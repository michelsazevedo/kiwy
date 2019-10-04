package file

import (
	"context"
	"encoding/csv"
	"io"
	"log"
	"os"
	"path"

	"cloud.google.com/go/storage"
)

type Csv struct {
	Filename string
	File     *os.File
	Writer   *csv.Writer
	Bucket   string
}

func NewCsv(filename, bucket string) *Csv {
	file, err := os.Create(filename)
	checkError("Error to create file", err)

	writer := csv.NewWriter(file)

	return &Csv{
		Filename: filename,
		File:     file,
		Writer:   writer,
		Bucket:   bucket,
	}
}

func (c *Csv) WriteLine(line []string) {
	err := c.Writer.Write(line)
	checkError("Cannot write to file", err)
}

func (c *Csv) SendToGcp() {
	ctx := context.Background()

	client, err := storage.NewClient(ctx)
	checkError("Unable to create Gcp client", err)

	object := path.Base(c.Filename)

	writerContext := client.Bucket(c.Bucket).Object(object).NewWriter(ctx)

	defer writerContext.Close()

	_, err = io.Copy(writerContext, c.File)
	checkError("Error to send file", err)
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
