package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/romietis/aocutil"
)

func main() {
	cards, _ := aocutil.ReadFile("input.txt")

	cardCopies := make(map[int]int)
	count := 0

	for _, card := range cards {
		count++
		id, winningNumber, myNumbers := parseCard(card)
		fmt.Println(id, winningNumber, myNumbers)

		multiplier := calcMultiplier(id, winningNumber, myNumbers)
		cardCopies[id]++

		for i := 1; i <= multiplier; i++ {
			cardCopies[id+i] += 1 * cardCopies[id]
		}
	}
	fmt.Println(cardCopies)

	var sum int
	for _, value := range cardCopies {
		sum += value
	}
	fmt.Println(sum)
	fmt.Println(len(cardCopies))
	fmt.Println(len(cards))
	fmt.Println(count)
	fmt.Println(cardCopies[0])
}

func calcMultiplier(id int, winningNumbers []int, myNumbers []int) int {
	multiplier := 0
	for _, winningNumber := range winningNumbers {
		for _, myNumber := range myNumbers {
			if winningNumber == myNumber {
				multiplier++
			}
		}
	}
	return multiplier
}

func parseCard(card string) (id int, winningNumbers []int, myNumbers []int) {
	parts := strings.Split(card, ":")
	idPart := parts[0]
	idString := strings.Split(idPart, "d")[1]
	idStringTrimmed := strings.TrimSpace(idString)
	id, _ = strconv.Atoi(idStringTrimmed)

	numbersPart := parts[1]
	numbers := strings.Split(numbersPart, "|")

	winningNumbersStrArray := strings.Split(numbers[0], " ")
	myNumbersStrArray := strings.Split(numbers[1], " ")

	for _, winningNumber := range winningNumbersStrArray {
		winningNumberInt, err := strconv.Atoi(winningNumber)
		if err != nil {
			continue
		}
		winningNumbers = append(winningNumbers, winningNumberInt)
	}

	for _, myNumber := range myNumbersStrArray {
		myNumberInt, err := strconv.Atoi(myNumber)
		if err != nil {
			continue
		}
		myNumbers = append(myNumbers, myNumberInt)
	}

	return id, winningNumbers, myNumbers
}
