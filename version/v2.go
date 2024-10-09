package version

import (
	"fmt"
	"os"
)

type VersionV2 struct {
	// stations map[string]station
	// cities   []string
}

func (v *VersionV2) splitLine(line string) (string, float64) {
	return "", 0
}

func (v *VersionV2) scanFile(f *os.File) {
	fmt.Println("Executing scan file")
}

func (v *VersionV2) postProcess() {
	fmt.Println("Executing Post process")
}

func (v *VersionV2) Execute(file string) {
	fmt.Printf("Running version 2\n")
}

func NewV2() Version {
	return &VersionV2{}
}
