package main

import (
	"log"
	"testing"
)

func TestCalcFuel(t *testing.T) {

	// Test Mass = 12		Expected: 2
	expected := 2
	actual := calcFuel(12)
	if expected != actual {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}

	// Test Mass = 14		Expected: 2
	expected = 2
	actual = calcFuel(14)
	if expected != actual {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}

	// Test Mass = 1969	 	Expected: 966
	expected = 966
	actual = calcFuel(1969)
	if expected != actual {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}

	//// Test Mass = 100756	Expected: 50346
	//expected = 50346
	//actual = calcFuel(100756)
	//if expected != actual {
	//	log.Fatalf("Expected %v, Got: %v", expected, actual)
	//}
}
