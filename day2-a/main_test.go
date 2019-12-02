package main

import (
	"log"
	"reflect"
	"testing"
)

func TestCalcFuel(t *testing.T) {

	// Test intCode =1, 0, 0, 0, 99					Expected: 2, 0, 0, 0, 99
	intCode := []int{1, 0, 0, 0, 99}
	expected := []int{2, 0, 0, 0, 99}
	computeIntCodeSequence(intCode)
	if !reflect.DeepEqual(expected, intCode) {
		log.Fatalf("Expected %v, Got: %v", expected, intCode)
	}

	// Test intCode =2, 3, 0, 3, 99					  Expected: 2, 3, 0, 6, 99
	intCode = []int{2, 3, 0, 3, 99}
	expected = []int{2, 3, 0, 6, 99}
	computeIntCodeSequence(intCode)
	if !reflect.DeepEqual(expected, intCode) {
		log.Fatalf("Expected %v, Got: %v", expected, intCode)
	}

	// Test intCode =2, 4, 4, 5, 99, 0				   Expected: 2, 4, 4, 5, 99, 9801
	intCode = []int{2, 4, 4, 5, 99, 0}
	expected = []int{2, 4, 4, 5, 99, 9801}
	computeIntCodeSequence(intCode)
	if !reflect.DeepEqual(expected, intCode) {
		log.Fatalf("Expected %v, Got: %v", expected, intCode)
	}

	// Test intCode =1, 1, 1, 4, 99, 5, 6, 0, 99		Expected: 30, 1, 1, 4, 2, 5, 6, 0, 99
	intCode = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	expected = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	computeIntCodeSequence(intCode)
	if !reflect.DeepEqual(expected, intCode) {
		log.Fatalf("Expected %v, Got: %v", expected, intCode)
	}
}
