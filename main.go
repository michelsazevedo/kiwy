package main

import (
	"os"
	"strconv"
	"strings"

	"github.com/michelsazevedo/kiwy/internal/secure"
	"github.com/michelsazevedo/kiwy/pkg/file"
	"github.com/michelsazevedo/kiwy/pkg/request"
)

func main() {
	concurrentWorkers := 1000

	result := make(chan request.Result)
	filename := strings.Join([]string{secure.RandomHex(20), "csv"}, ".")

	go request.MakeParallelsRequests(concurrentWorkers, result)

	file := file.NewCsv(filename, os.Getenv("BUCKET"))
	var line []string

	for res := range result {
		line = []string{
			res.TableId,
			res.Key,
			res.StartDate.String(),
			res.EndDate.String(),
			res.SysDate.String(),
			strconv.FormatFloat(res.SysTime, 'f', 4, 32),
		}
		file.WriteLine(line)
	}
	file.Flush()
	file.SendToGcp()
}
