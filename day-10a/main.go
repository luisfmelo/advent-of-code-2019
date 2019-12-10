package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"time"
)

const (
	EMPTY    = '.'
	ASTEROID = '#'
)

type coordinate struct {
	x int
	y int
}

func (c *coordinate) toString() string {
	return fmt.Sprintf("(%v,%v)", c.x, c.y)
}

func (c *coordinate) calcAngle(c2 coordinate) float64 {
	angle := math.Atan2(float64(c2.x-c.x), float64(c.y-c2.y)) * 180 / math.Pi
	if angle < 0 {
		angle = angle + 360
	}
	return angle
}

func getAsteroidBelt(inputMap []string) []coordinate {
	var belt []coordinate
	for y, row := range inputMap {
		for x, object := range row {
			if object == ASTEROID {
				belt = append(belt, coordinate{x: x, y: y})
			}

		}
	}
	return belt
}

func getBestMonitoringStationLocation(inputMap []string) (coordinate, int) {
	belt := getAsteroidBelt(inputMap)
	bestBet := coordinate{}
	maxAsteroidsInSight := 0

	for _, asteroid := range belt {
		angles := map[float64]bool{}

		for _, target := range belt {
			if asteroid.x == target.x && asteroid.y == target.y {
				continue
			}

			// Calc Angle
			angle := asteroid.calcAngle(target)

			if _, exists := angles[angle]; !exists {
				angles[angle] = true
			}
		}
		if len(angles) > maxAsteroidsInSight {
			maxAsteroidsInSight = len(angles)
			bestBet = asteroid
		}

	}

	return bestBet, maxAsteroidsInSight
}

func read(file string) []string {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var input []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		input = append(input, line)
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

	bestBet, maxAsteroidsInSight := getBestMonitoringStationLocation(input)
	fmt.Printf("Best location: %s. Can detect %v asteroids.\n\n", bestBet.toString(), maxAsteroidsInSight)

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
