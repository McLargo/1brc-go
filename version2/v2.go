package version2

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type station struct {
	min, max, sum float64
	total         int32
}

type VersionV2 struct {
	stations map[string]station
}

func (v *VersionV2) splitLine(line string) (string, float64) {
	start, end, _ := strings.Cut(line, ";")

	tempFl, _ := strconv.ParseFloat(end, 64)
	return start, tempFl
}

func (v *VersionV2) scanFile(f *os.File) {
	reader := bufio.NewReader(f)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		city, tempFl := v.splitLine(string(line[:]))
		st, found := v.stations[city]
		if found {
			if st.min > tempFl {
				st.min = tempFl
			}
			if st.max < tempFl {
				st.max = tempFl
			}
			st.total++
			st.sum = st.sum + tempFl
			v.stations[city] = st

		} else {
			st := station{
				min:   tempFl,
				max:   tempFl,
				sum:   tempFl,
				total: 1,
			}
			v.stations[city] = st
		}
	}
}

func (v *VersionV2) postProcess() {
	// avoid sorting cities
	for c, st := range v.stations {
		avg := st.sum / float64(st.total)
		fmt.Printf("City: %s (%.1f/%.1f/%.1f)\n", c, st.min, avg, st.max)
	}
}

func (v *VersionV2) Execute(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	v.stations = make(map[string]station)
	v.scanFile(f)
	v.postProcess()
}

func NewV2() *VersionV2 {
	return &VersionV2{}
}
