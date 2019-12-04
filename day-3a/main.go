package main

import (
	"bufio"
	"fmt"
	"math"
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

func newPointFromString(str string) Point {
	splitted := strings.Split(str, ",")
	x, err := strconv.Atoi(splitted[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(splitted[1])
	if err != nil {
		panic(err)
	}
	return Point{x: x, y: y}
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

func (p *Point) calcManhattanDistanceToPoint(p2 Point) int {
	return int(math.Abs(float64(p.x-p2.x)) + math.Abs(float64(p.y-p2.y)))
}

func getPointsOfWire(result chan map[string]bool, wirePath string) {
	wirePointsStr := map[string]bool{}
	currPoint := Point{
		x: 0,
		y: 0,
	}

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
				wirePointsStr[currPoint.toString()] = true
			}
		case "L":
			for i := 0; i < dist; i++ {
				currPoint.goLeft()
				wirePointsStr[currPoint.toString()] = true
			}
		case "U":
			for i := 0; i < dist; i++ {
				currPoint.goUp()
				wirePointsStr[currPoint.toString()] = true
			}
		case "D":
			for i := 0; i < dist; i++ {
				currPoint.goDown()
				wirePointsStr[currPoint.toString()] = true
			}
		}
	}
	result <- wirePointsStr
}

func calcMinimumWiresCrossDistance(wire1Path string, wire2Path string) int {
	ch := make(chan map[string]bool, 2)

	go getPointsOfWire(ch, wire1Path)
	go getPointsOfWire(ch, wire2Path)

	result1 := <-ch
	result2 := <-ch
	minDist := MaxInt
	for p1str, _ := range result1 {
		if _, ok := result2[p1str]; ok {
			point := newPointFromString(p1str)
			dist := point.calcManhattanDistanceToPoint(Point{x: 0, y: 0})
			if dist < minDist {
				minDist = dist
			}

		}
	}

	return minDist
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
		result := calcMinimumWiresCrossDistance(wires[0], wires[1])
		fmt.Println("Minimum distance:", result)
	} else {
		fmt.Println("Input invalid")
	}

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
