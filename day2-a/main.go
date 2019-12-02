package main

import (
	"bufio"
	"fmt"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"strings"
)

func computeIntCodeSequence(intCodeSequence []int) {
	for i := 0; i < len(intCodeSequence); i++ {
		operation :=intCodeSequence[i]

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
			return
		default:
			log.Fatalf("Operation not valid: %v", i)
		}
	}
}

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var intCodeSequence []int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		// Get int code
		for _, strInt := range strings.Split(line, ",") {
			i, err := strconv.Atoi(strInt)
			if err != nil {
				panic(err)
			}
			intCodeSequence = append(intCodeSequence, i)
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	// Fix input data
	intCodeSequence[1] = 12
	intCodeSequence[2] = 2

	computeIntCodeSequence(intCodeSequence)

	fmt.Println("Result:", intCodeSequence[0])
}
