package main

import (
	"log"
	"reflect"
	"testing"
)

func TestProcessImage(t *testing.T) {
	input := []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0}
	expected := []int{0, 1, 1, 0}
	actual := processImage(input, 2, 2)
	if reflect.DeepEqual(actual, expected) {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}
