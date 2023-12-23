package main

import (
	"fmt"
	"strings"

	"github.com/romietis/aocutil"
)

func main() {
	file, _ := aocutil.ReadFile("input.txt")
	var steps []string

	for _, line := range file {
		steps = strings.Split(line, ",")
	}

	// Part 1
	sequenceValue := 0
	for _, step := range steps {
		step_value := 0
		for _, char := range step {
			step_value += int(char)
			step_value = multiplyBy17(step_value)
			step_value = reminderOf256(step_value)
		}
		sequenceValue += step_value
	}

	fmt.Println(sequenceValue)
}

func multiplyBy17(value int) int {
	return value * 17
}

func reminderOf256(value int) int {
	return value % 256
}
