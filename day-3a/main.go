package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MaxInt = int(^uint(0) >> 1)

type Point struct {
	x int
	y int
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

func getPointsOfWire(wirePath string) []Point {
	var wirePoints []Point
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
				wirePoints = append(wirePoints, currPoint)
			}
		case "L":
			for i := 0; i < dist; i++ {
				currPoint.goLeft()
				wirePoints = append(wirePoints, currPoint)
			}
		case "U":
			for i := 0; i < dist; i++ {
				currPoint.goUp()
				wirePoints = append(wirePoints, currPoint)
			}
		case "D":
			for i := 0; i < dist; i++ {
				currPoint.goDown()
				wirePoints = append(wirePoints, currPoint)
			}
		}
	}
	return wirePoints
}

func getIntersections(wire1Path []Point, wire2Path []Point) []Point {
	var intersections []Point
	for _, p1 := range wire1Path {
		for _, p2 := range wire2Path {
			if p1.x == p2.x && p1.y == p2.y {
				intersections = append(intersections, p1)
			}
		}
	}
	return intersections
}

func calcMinimumWiresCrossDistance(wire1Path string, wire2Path string) int {
	wire1Points := getPointsOfWire(wire1Path)
	wire2Points := getPointsOfWire(wire2Path)

	intersections := getIntersections(wire1Points, wire2Points)
	if len(intersections) == 0 {
		return -1
	}

	minDist := MaxInt
	for _, intersectionPoint := range intersections {
		dist := intersectionPoint.calcManhattanDistanceToPoint(Point{x: 0, y: 0})
		if dist < minDist {
			minDist = dist
		}
	}

	return minDist
}

func main() {
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
		return
	}

	fmt.Println("Input invalid")
}
