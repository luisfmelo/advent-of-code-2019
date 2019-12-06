package main

import (
	"log"
	"testing"
)

func TestNumberOfTransfers(t *testing.T) {

	input := []string{"COM)B", "B)C", "C)D", "D)E", "E)F", "B)G", "G)H", "D)I", "E)J", "J)K", "K)L", "K)YOU", "I)SAN",}
	actual := computeNumberOfTransfers(input, "YOU", "SAN")
	expected := 4
	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}
