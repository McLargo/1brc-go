package version3

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const mb = 1024 * 1024

type station struct {
	min, max, sum float64
	total         int32
}

type VersionV3 struct {
	stations map[string]station
	cities   []string
}

func (v *VersionV3) splitLine(line string) (string, float64) {
	start, end, _ := strings.Cut(line, ";")

	tempFl, _ := strconv.ParseFloat(end, 64)
	return start, tempFl
}

func (v *VersionV3) handleLine(line string) {
	city, tempFl := v.splitLine(line)
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
		v.cities = append(v.cities, city)
	}
}

func (v *VersionV3) postProcess() {
	fmt.Println(v.cities)
	sort.Strings(v.cities)
	for _, c := range v.cities {
		st := v.stations[c]
		avg := st.sum / float64(st.total)
		fmt.Printf("City: %s (%.1f/%.1f/%.1f)\n", c, st.min, avg, st.max)
	}
}

func (v *VersionV3) Execute(file string) {
	v.stations = make(map[string]station)
	v.cities = make([]string, 0)

	// This channel is used to send every line in various go-routines.
	channel := make(chan (string))

	// Done is a channel to signal the main thread that all the lines has been processed.
	done := make(chan (bool), 1)

	// Read all incoming lines from the channel and handle them.
	go func() {
		for s := range channel {
			v.handleLine(s)
		}

		// Signal the main thread that all the words have entered the dictionary.
		done <- true
	}()

	// Current signifies the counter for bytes of the file.
	var current int64

	// Limit signifies the chunk size of file to be processed by every thread.
	var limit int64 = 100 * mb

	// A waitgroup to wait for all go-routines to finish.
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		v.readFileInBatch(current, limit, file, channel)

		// Increment the current by 1+(last byte read by previous thread).
		current += limit + 1
		wg.Done()
	}

	// Wait for all go routines to complete.
	wg.Wait()
	close(channel)

	// Wait for dictionary to process all the words.
	<-done
	close(done)

	v.postProcess()
}

func (v *VersionV3) readFileInBatch(offset int64, limit int64, fileName string, channel chan (string)) {
	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Move the pointer of the file to the start of designated chunk.
	file.Seek(offset, 0)
	reader := bufio.NewReader(file)

	// This block of code ensures that the start of chunk is the start of a line. If
	// a character is encountered at the given position it moves a few bytes till
	// the end of the line.
	if offset != 0 {
		_, err = reader.ReadBytes('\n')
		if err == io.EOF {
			return
		}

		if err != nil {
			panic(err)
		}
	}

	var cummulativeSize int64
	for {
		// Break if read size has exceed the chunk size.
		if cummulativeSize > limit {
			break
		}

		b, err := reader.ReadBytes('\n')

		// Break if end of file is encountered.
		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}

		cummulativeSize += int64(len(b))
		s := strings.TrimSpace(string(b))
		if s != "" {
			// Send the read line in the channel to process.
			channel <- s
		}
	}
}

func NewV3() *VersionV3 {
	return &VersionV3{}
}
