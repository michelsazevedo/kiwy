package main

import (
	"os"

	"github.com/michelsazevedo/kiwy/pkg/file"
	"github.com/michelsazevedo/kiwy/pkg/request"
)

func main() {
	concurrentWorkers := 2

	result := make(chan request.Result)
	filename := "req.csv"

	go request.MakeParallelsRequests(concurrentWorkers, result)

	file := file.NewCsv(filename, os.Getenv("BUCKET"))
	var line []string

	for res := range result {
		line = []string{
			res.TableId,
			res.Key,
		}
		file.WriteLine(line)
	}

	file.SendToGcp()
}
