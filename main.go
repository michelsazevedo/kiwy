package main

import (
	"github.com/michelsazevedo/kiwy/pkg/file"
	"github.com/michelsazevedo/kiwy/pkg/request"
)

func main() {
	concurrentWorkers := 2

	result := make(chan request.Result)
	filename := "req.csv"

	go request.MakeParallelsRequests(concurrentWorkers, result)

	file := file.NewCsv(filename)
	var line []string

	for res := range result {
		line = []string{
			res.TableId,
			res.Key,
		}
		file.WriteLine(line)
	}
}
