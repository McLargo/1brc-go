package version1

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type station struct {
	min, max, sum float64
	total         int32
}

type VersionV1 struct {
	stations map[string]station
	cities   []string
}

func (v *VersionV1) splitLine(line string) (string, float64) {
	start, end, _ := strings.Cut(line, ";")

	tempFl, _ := strconv.ParseFloat(end, 64)
	return string(start), tempFl
}

func (v *VersionV1) scanFile(f *os.File) {
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		city, tempFl := v.splitLine(scanner.Text())
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
			v.cities = append(v.cities, city)
			st := station{
				min:   tempFl,
				max:   tempFl,
				sum:   tempFl,
				total: 1,
			}
			v.stations[city] = st
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}

func (v *VersionV1) postProcess() {
	sort.Strings(v.cities)
	for _, c := range v.cities {
		st := v.stations[c]
		avg := st.sum / float64(st.total)
		fmt.Printf("City: %s (%.1f/%.1f/%.1f)\n", c, st.min, avg, st.max)
	}
}

func (v *VersionV1) Execute(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	v.stations = make(map[string]station)
	v.cities = make([]string, 0)
	v.scanFile(f)
	v.postProcess()
}

func NewV1() *VersionV1 {
	return &VersionV1{}
}
