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

type Point struct {
	x int
	y int
}

func (p *Point) toString() string {
	return fmt.Sprintf("%v,%v", p.x, p.y)
}

func (p *Point) goLeft() {
	p.x--
}

func (p *Point) goRight() {
	p.x++
}

func (p *Point) goUp() {
	p.y++
}

func (p *Point) goDown() {
	p.y--
}

func getPointsOfWire(ch chan map[string]int, wirePath string) {
	wirePoints := map[string]int{}
	currPoint := Point{
		x: 0,
		y: 0,
	}

	steps := 0
	for _, cmd := range strings.Split(wirePath, ",") {
		direction := string(cmd[0])
		dist, err := strconv.Atoi(cmd[1:])
		if err != nil {
			panic(err)
		}

		switch direction {
		case "R":
			for i := 0; i < dist; i++ {
				currPoint.goRight()
				steps++
				if value, ok := wirePoints[currPoint.toString()]; !ok || steps < value {
					wirePoints[currPoint.toString()] = steps
				}
			}
		case "L":
			for i := 0; i < dist; i++ {
				currPoint.goLeft()
				steps++
				if value, ok := wirePoints[currPoint.toString()]; !ok || steps < value {
					wirePoints[currPoint.toString()] = steps
				}
			}
		case "U":
			for i := 0; i < dist; i++ {
				currPoint.goUp()
				steps++
				if value, ok := wirePoints[currPoint.toString()]; !ok || steps < value {
					wirePoints[currPoint.toString()] = steps
				}
			}
		case "D":
			for i := 0; i < dist; i++ {
				currPoint.goDown()
				steps++
				if value, ok := wirePoints[currPoint.toString()]; !ok || steps < value {
					wirePoints[currPoint.toString()] = steps
				}
			}
		}
	}
	ch <- wirePoints
}

func calcMinimumWiresCrossSteps(wire1Path string, wire2Path string) int {
	ch := make(chan map[string]int, 2)
	go getPointsOfWire(ch, wire1Path)
	go getPointsOfWire(ch, wire2Path)

	result1 := <-ch
	result2 := <-ch

	minSteps := MaxInt
	for p1str, steps1 := range result1 {
		if steps2, ok := result2[p1str]; ok {
			if totalSteps := steps1 + steps2; totalSteps < minSteps {
				minSteps = totalSteps
			}
		}
	}

	return minSteps
}

func main() {
	start := time.Now()

	f, err := os.OpenFile("./input.txt", os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var wires []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		wires = append(wires, line)
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	if len(wires) == 2 {
		result := calcMinimumWiresCrossSteps(wires[0], wires[1])
		fmt.Println("Minimum steps:", result)
	} else {
		fmt.Println("Input invalid")
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
