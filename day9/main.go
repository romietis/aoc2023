package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/romietis/aocutil"
)

func main() {
	file, _ := aocutil.ReadFile("example.txt")
	var sequenceInt []int
	for _, line := range file {
		sequence := strings.Split(line, " ")
		for _, seqInt := range sequence {
			seqInt, _ := strconv.Atoi(seqInt)
			sequenceInt = append(sequenceInt, seqInt)
		}

		fmt.Println(getDiff(sequenceInt))
	}

}

func getDiff(sequence []int) (diff int) {

	for i := 0; i < len(sequence)-1; i++ {
		diff = sequence[i+1] - sequence[i]
	}
	return diff
}
