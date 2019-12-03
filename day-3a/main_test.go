package main

import (
	"log"
	"testing"
)

func TestCalcFuel(t *testing.T) {

	w1Path := "R75,D30,R83,U83,L12,D49,R71,U7,L72"
	w2Path := "U62,R66,U55,R34,D71,R55,D58,R83"
	expected := 159
	actual := calcMinimumWiresCrossDistance(w1Path, w2Path)
	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}

	w1Path = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51"
	w2Path = "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"
	expected = 135
	actual = calcMinimumWiresCrossDistance(w1Path, w2Path)
	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}
