package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type operation int

const (
	ADD      operation = 1
	MULTIPLY operation = 2
	INPUT    operation = 3
	OUTPUT   operation = 4
	JIT      operation = 5 // Jump if true
	JIF      operation = 6 // Jump if false
	LESS     operation = 7
	EQUALS   operation = 8
	HALT     operation = 99
)

type mode int

const (
	POSITION mode = iota
	IMMEDIATE
)

type instruction struct {
	opCode          operation
	firstParamMode  mode
	secondParamMode mode
	thirdParamMode  mode
}

func newInstruction(i int) instruction {
	return instruction{
		opCode:          operation(i % 100),
		firstParamMode:  mode(i / 100 % 10),
		secondParamMode: mode(i / 1000 % 10),
		thirdParamMode:  mode(i / 10000 % 10),
	}
}

func stringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
	}
	return i
}
func intToString(i int) string {
	return strconv.Itoa(i)
}

func getValueByMode(intCodeSequence []int, idx int, m mode) int {
	value := intCodeSequence[idx]
	switch m {
	case POSITION: // Position mode
		return intCodeSequence[value]
	case IMMEDIATE: // immediate mode
		return value
	}
	log.Fatalf("Mode %v not valid", m)
	return -1
}

func readInput() int {
	var inputValue int
	fmt.Print("Insert an input: ")
	_, err := fmt.Scanf("%d", &inputValue)
	if err != nil {
		log.Fatalf("Error reading input. %s", err.Error())
	}
	return inputValue
}

func computeIntCodeSequence(intCodeSequence []int, inputCh <-chan int, outputCh chan<- int) {
	for ptr := 0; ptr < len(intCodeSequence); {
		inst := newInstruction(intCodeSequence[ptr])

		switch inst.opCode {
		case ADD:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			param2 := getValueByMode(intCodeSequence, ptr+2, inst.secondParamMode)
			param3 := intCodeSequence[ptr+3]
			intCodeSequence[param3] = param1 + param2
			ptr += 4

		case MULTIPLY:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			param2 := getValueByMode(intCodeSequence, ptr+2, inst.secondParamMode)
			param3 := intCodeSequence[ptr+3]
			intCodeSequence[param3] = param1 * param2
			ptr += 4

		case INPUT:
			intCodeSequence[intCodeSequence[ptr+1]] = <-inputCh
			ptr += 2

		case OUTPUT:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			outputCh <- param1
			ptr += 2

		case JIT:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			param2 := getValueByMode(intCodeSequence, ptr+2, inst.secondParamMode)
			if param1 != 0 {
				ptr = param2
			} else {
				ptr += 3
			}

		case JIF:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			param2 := getValueByMode(intCodeSequence, ptr+2, inst.secondParamMode)
			if param1 == 0 {
				ptr = param2
			} else {
				ptr += 3
			}

		case LESS:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			param2 := getValueByMode(intCodeSequence, ptr+2, inst.secondParamMode)
			param3 := intCodeSequence[ptr+3]
			if param1 < param2 {
				intCodeSequence[param3] = 1
			} else {
				intCodeSequence[param3] = 0
			}
			ptr += 4

		case EQUALS:
			param1 := getValueByMode(intCodeSequence, ptr+1, inst.firstParamMode)
			param2 := getValueByMode(intCodeSequence, ptr+2, inst.secondParamMode)
			param3 := intCodeSequence[ptr+3]
			if param1 == param2 {
				intCodeSequence[param3] = 1
			} else {
				intCodeSequence[param3] = 0
			}
			ptr += 4

		case HALT:
			close(outputCh)
			return

		default:
			panic("Op code not valid")
		}
	}

	log.Fatalf("The code did not halt")
	return
}

func read(file string) []int {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
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
			i := stringToInt(strInt)
			initialIntCodeSequence = append(initialIntCodeSequence, i)
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return initialIntCodeSequence
}

func getOutput(outputCh chan int) int {
	output := ""
	for o := range outputCh {
		output += intToString(o)
	}
	return stringToInt(output)
}

func main() {
	start := time.Now()
	inputFilePath := os.Args[1]
	initialIntCodeSequence := read(inputFilePath)

	// Process inputs
	inputCh := make(chan int)
	go func() {
		for _, in := range os.Args[2:] {
			var i = stringToInt(in)
			inputCh <- i
		}
	}()
	outputCh := make(chan int)
	go computeIntCodeSequence(initialIntCodeSequence, inputCh, outputCh)
	log.Printf("Output: %d\n", getOutput(outputCh))

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
