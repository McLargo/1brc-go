package version

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

func splitLine(line string) (string, float64) {
	start, end, _ := strings.Cut(line, ";")

	tempFl, _ := strconv.ParseFloat(end, 64)
	return string(start), tempFl
}

func scanFile(f *os.File) (map[string]station, []string) {
	var stations = make(map[string]station)
	var cities []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		city, tempFl := splitLine(scanner.Text())
		st, found := stations[city]
		if found {
			if st.min > tempFl {
				st.min = tempFl
			}
			if st.max < tempFl {
				st.max = tempFl
			}
			st.total++
			st.sum = st.sum + tempFl
			stations[city] = st

		} else {
			cities = append(cities, city)
			st := station{
				min:   tempFl,
				max:   tempFl,
				sum:   tempFl,
				total: 1,
			}
			stations[city] = st
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
	return stations, cities
}
func postProcess(stations map[string]station, cities []string) {
	sort.Strings(cities)
	for _, c := range cities {
		st := stations[c]
		avg := st.sum / float64(st.total)
		fmt.Printf("City: %s (%.1f/%.1f/%.1f)\n", c, st.min, avg, st.max)
	}
}
func ExecuteV1(file string) {

	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	stations, cities := scanFile(f)
	postProcess(stations, cities)
}
