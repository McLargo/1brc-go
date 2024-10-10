package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/mclargo/1brc/version1"
	"github.com/mclargo/1brc/version2"
	"github.com/mclargo/1brc/version3"
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
		v1 := version1.NewV1()
		v1.Execute(*filePath)
	case "v2":
		fmt.Println("running v2")
		v2 := version2.NewV2()
		v2.Execute(*filePath)
	case "v3":
		fmt.Println("running v3")
		v3 := version3.NewV3()
		v3.Execute(*filePath)
	default:
		panic("version not found")
	}
}
