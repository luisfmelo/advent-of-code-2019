package main

import (
	"log"
	"testing"
)

func TestCountDirectAndIndirectOrbits(t *testing.T) {

	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L"}
	actual := computeNumberOfOrbits(input)
	expected := 42
	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}
