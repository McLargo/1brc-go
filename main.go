package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mclargo/1brc/version"
)

func main() {
	defer func(s time.Time) {
		fmt.Println("Time taken to complete the challenge: ", time.Since(s))
	}(time.Now())

	var filePath = flag.String("file", "data/measurements.txt", "file to read files from")
	var v = flag.String("version", "v1", "version to execute")
	flag.Parse()

	switch *v {
	case "v1":
		fmt.Println("running v1")
		version.ExecuteV1(*filePath)
	default:
		panic("version not found")
	}
}
