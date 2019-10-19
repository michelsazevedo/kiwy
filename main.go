package main

import (
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/michelsazevedo/kiwy/internal/secure"
	"github.com/michelsazevedo/kiwy/pkg/file"
	"github.com/michelsazevedo/kiwy/pkg/request"
)

func main() {
	concurrentWorkers, err := strconv.Atoi(os.Getenv("CONCURRENCY"))
	if err != nil {
		log.Fatal("Error to load number of workers")
	}

	collection := make(chan request.CollectionResult)
	filename := strings.Join([]string{secure.RandomHex(20), "csv"}, ".")

	go request.MakeParallelsRequests(concurrentWorkers, collection)

	file := file.NewCsv(filename, os.Getenv("BUCKET"))
	var line []string

	for result := range collection {
		for _, res := range result.Result {
			line = []string{
				res.DateTime,
				strconv.Itoa(res.Time),
				strconv.Itoa(res.Count),
			}
			file.WriteLine(line)
		}
	}
	file.Flush()
	//file.SendToGcp()
}
