package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

const (
	BLACK       = 0
	WHITE       = 1
	TRANSPARENT = 2
)

func getVisiblePixel(wg *sync.WaitGroup, input []int, pixelIdx int, imageArea int, output []int) {
	layers := len(input) / imageArea
	for layerIdx := 0; layerIdx < layers; layerIdx++ {
		pixel := input[imageArea*layerIdx+pixelIdx]
		if pixel != TRANSPARENT {
			output[pixelIdx] = pixel
			wg.Done()
			return
		}
	}
}

func processImage(input []int, pixelsWide int, pixelsTall int) []int {
	var wg sync.WaitGroup
	imageArea := pixelsTall * pixelsWide

	finalImage := make([]int, imageArea, imageArea)
	for pixelIdx := 0; pixelIdx < imageArea; pixelIdx++ {
		wg.Add(1)
		go getVisiblePixel(&wg, input, pixelIdx, imageArea, finalImage)
	}
	wg.Wait()

	return finalImage
}

func drawImage(input []int, pixelsWide int) {

	for idx, pixel := range input {
		if idx%pixelsWide == 0 {
			fmt.Printf("\n")
		}
		if pixel == BLACK {
			fmt.Printf("■")
		} else if pixel == WHITE {
			fmt.Printf(" ")
		} else {
			fmt.Printf("?")
		}
	}

	fmt.Printf("\n\n")
}

func renderImage(input []int, pixelsWide int, pixelsTall int) []int {
	imageArea := pixelsTall * pixelsWide

	finalImage := make([]int, imageArea, imageArea)
	pixelIdx := imageArea - 1
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] != TRANSPARENT {
			finalImage[pixelIdx] = input[i]
		}
		if pixelIdx == 0 {
			pixelIdx = imageArea
		}
		pixelIdx--
	}

	return finalImage
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

	result := renderImage(input, 25, 6)
	drawImage(result, 25)

	result = processImage(input, 25, 6)
	drawImage(result, 25)

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
