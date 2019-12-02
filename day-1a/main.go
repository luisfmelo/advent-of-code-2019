package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func calcFuel(mass int) int{
	return int(math.Floor(float64(mass)/3)) - 2
}

func main() {
	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	counter := 0
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		counter += calcFuel(mass)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Result:", counter)
}
