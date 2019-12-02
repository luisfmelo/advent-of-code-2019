package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"strings"
)

func computeIntCodeSequence(intCodeSequence []int, noun, verb int) int {
	intCodeSequence[1] = noun
	intCodeSequence[2] = verb

	loop: for i := 0; i < len(intCodeSequence); i++ {
		operation := intCodeSequence[i]

		switch operation {
		case 1, 2:
			idxOp1 := intCodeSequence[i+1]
			idxOp2 := intCodeSequence[i+2]
			idxResult := intCodeSequence[i+3]
			switch operation {
			case 1:
				intCodeSequence[idxResult] = intCodeSequence[idxOp1] + intCodeSequence[idxOp2]
			case 2:
				intCodeSequence[idxResult] = intCodeSequence[idxOp1] * intCodeSequence[idxOp2]
			}
			i = i + 3
		case 99:
			break loop
		default:
			log.Fatalf("Operation not valid: %v", i)
		}
	}
	return intCodeSequence[0]
}

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var initialIntCodeSequence []int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		// Get int code
		for _, strInt := range strings.Split(line, ",") {
			i, err := strconv.Atoi(strInt)
			if err != nil {
				panic(err)
			}
			initialIntCodeSequence = append(initialIntCodeSequence, i)
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	// Brute force to try to achieve - 19690720
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			intCodeSequence := make([]int, len(initialIntCodeSequence))
			_ = copy(intCodeSequence, initialIntCodeSequence)
			res := computeIntCodeSequence(intCodeSequence, noun, verb)
			if res == 19690720 {
				fmt.Println("Result:", 100*noun+verb)
				return
			}
		}
	}
	fmt.Println("Not Found")
}
