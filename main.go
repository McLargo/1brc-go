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
		v1 := version.NewV1()
		v1.Execute(*filePath)
	case "v2":
		fmt.Println("running v2")
		v2 := version.NewV2()
		v2.Execute(*filePath)
	default:
		panic("version not found")
	}
}
