package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

const MaxInt = int(^uint(0) >> 1)

const (
	BLACK       = 0
	WHITE       = 1
	TRANSPARENT = 2
)

func processImageLayerPixelCount(layer []int) map[int]int {
	output := map[int]int{}
	for _, pixel := range layer {
		if _, exists := output[pixel]; exists {
			output[pixel]++
		} else {
			output[pixel] = 1
		}
	}
	return output
}

func processImage(input []int, pixelsWide int, pixelsTall int) int {
	imageArea := pixelsTall * pixelsWide

	minimumZeros := MaxInt
	result := 0
	for i := 0; i < len(input); i += imageArea {
		output := processImageLayerPixelCount(input[i : i+imageArea])
		if output[BLACK] < minimumZeros {
			result = output[WHITE] * output[TRANSPARENT]
			minimumZeros = output[BLACK]
		}
	}
	return result
}

func read(file string) []int {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var input []int
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		for i, _ := range line {
			integer, _ := strconv.Atoi(line[i : i+1])
			input = append(input, integer)
		}
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return input
}

func main() {
	start := time.Now()
	inputFilePath := os.Args[1]
	input := read(inputFilePath)

	result := processImage(input, 25, 6)
	log.Printf("Result: %d\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
