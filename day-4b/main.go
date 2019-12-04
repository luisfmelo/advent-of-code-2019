package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

const MaxInt = int(^uint(0) >> 1)

func ruleIsSixDigitNumber(password int) bool {
	return password > 99999 && password < 1000000
}

func ruleNeverDecrease(password int) bool {
	//never decrease is equal to never increase (reverse order)
	for lastRemaining := MaxInt; password > 0; password /= 10 {
		remaining := password % 10
		if remaining > lastRemaining {
			return false
		}
		lastRemaining = remaining
	}
	return true
}

func ruleHasTwoAdjacentNumbers(password int) bool {
	for lastRemaining := MaxInt; password/10 > 0; password /= 10 {
		remaining := password % 10
		nextValue := password / 10 % 10
		nextNextValue := password / 100 % 10
		if remaining == nextValue && remaining != nextNextValue && remaining != lastRemaining {
			return true
		}
		lastRemaining = remaining
	}
	return false
}

func isPasswordValid(password int) bool {
	return ruleIsSixDigitNumber(password) && ruleNeverDecrease(password) && ruleHasTwoAdjacentNumbers(password)
}

func main() {
	start := time.Now()

	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	numberOfValidPassword := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		splitted := strings.Split(sc.Text(), "-")

		lowerBound, err := strconv.Atoi(splitted[0])
		if err != nil {
			panic(err)
		}
		upperBound, err := strconv.Atoi(splitted[1])
		if err != nil {
			panic(err)
		}

		for i := lowerBound; i <= upperBound; i++ {
			if isPasswordValid(i) {
				numberOfValidPassword++
			}
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("Number of combinations: %v\n", numberOfValidPassword)
	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
