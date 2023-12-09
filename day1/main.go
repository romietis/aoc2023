package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// document := []string{
	// 	"two1nine",
	// 	"eightwothree",
	// 	"abcone2threexyz",
	// 	"xtwone3four",
	// 	"4nineeightseven2",
	// 	"zoneight234",
	// 	"7pqrstsixteen"}

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	pairs := []string{}
	for scanner.Scan() {
		// for _, word := range document {
		var pair string

		replaced_word := StringsToInts(scanner.Text())
		pair += findFirstDigit(replaced_word)
		pair += findLastDigit(replaced_word)
		pairs = append(pairs, pair)
	}
	fmt.Println(pairs)

	var sum int
	for _, pair := range pairs {
		pair_int, _ := strconv.Atoi(pair)
		sum += pair_int
	}
	fmt.Println(sum)
}

func findFirstDigit(word string) string {
	for _, letter := range word {
		if letter >= '0' && letter <= '9' {
			return string(letter)
		}
	}
	return ""
}

func findLastDigit(word string) string {
	for i := len(word) - 1; i >= 0; i-- {
		if word[i] >= '0' && word[i] <= '9' {
			return string(word[i])
		}
	}
	return ""
}

func StringsToInts(word string) string {
	r := strings.NewReplacer(

		"twone", "21",
		"oneight", "18",
		"sevenine", "79",
		"eightwo", "82",
		"fiveight", "58",
		"eighthree", "83",
		"nineight", "98",
		"threeight", "38",

		"one", "1",
		"two", "2",
		"three", "3",
		"four", "4",
		"five", "5",
		"six", "6",
		"seven", "7",
		"eight", "8",
		"nine", "9",
	)
	return r.Replace(word)
}
