package version

import "os"

type station struct {
	min, max, sum float64
	total         int32
}

type Version interface {
	splitLine(line string) (string, float64)
	scanFile(f *os.File)
	postProcess()
	Execute(file string)
}
