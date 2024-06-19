package version

import (
	"bufio"
	"fmt"
	"os"
)

func ExecuteV1(file string) {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}
}
