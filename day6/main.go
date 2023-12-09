package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/romietis/aocutil"
)

func main() {
	timeDistance := make(map[int]int)

	// Load file into map
	file, _ := aocutil.ReadFile("example.txt")
	timePart := strings.Split(file[0], ":")
	distancePart := strings.Split(file[1], ":")

	times := strings.Split(timePart[1], " ")
	distances := strings.Split(distancePart[1], " ")

	for i := 0; i < len(times); i++ {
		timeInt, _ := strconv.Atoi(times[i])
		distanceInt, _ := strconv.Atoi(distances[i])
		timeDistance[timeInt] = distanceInt
	}

	var multi []int
	for key, value := range timeDistance {
		minVal, maxVal := getQuadraticRaceValues(key, value)
		naturalMin, naturalMax := getNaturalMinMax(minVal, maxVal)
		ways := waysOfWinning(naturalMin, naturalMax)
		multi = append(multi, len(ways))
	}

	fmt.Println(multi)

	multiplier := 1
	for _, i := range multi {
		multiplier *= i
	}
	fmt.Println(multiplier)
}

func getQuadraticRaceValues(time int, distanceToBeat int) (float64, float64) {
	delta := math.Pow(float64(time), 2) - 4*float64(distanceToBeat)

	x1 := (float64(time) - math.Sqrt(delta)) / 2
	x2 := (float64(time) + math.Sqrt(delta)) / 2
	return x1, x2
}

func getNaturalMinMax(minValue float64, maxValue float64) (int, int) {
	naturalMin := math.Ceil(minValue)
	naturalMax := math.Floor(maxValue)

	if naturalMin == minValue {
		naturalMin++
	}

	if naturalMax == maxValue {
		naturalMax--
	}
	return int(naturalMin), int(naturalMax)
}

func waysOfWinning(minTime int, maxTime int) []int {
	var temp []int
	for i := minTime; i <= maxTime; i++ {
		temp = append(temp, i)
	}
	return temp
}
