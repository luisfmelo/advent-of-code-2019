package main

import (
	"log"
	"testing"
)

func TestIsPasswordValid(t *testing.T) {

	password := 111111
	actual := isPasswordValid(password)
	if actual != true {
		log.Fatalf("%v: Expected %v, Got: %v", password, true, actual)
	}

	password = 123789
	actual = isPasswordValid(password)
	if actual != false {
		log.Fatalf("%v: Expected %v, Got: %v", password, false, actual)
	}

	password = 223450
	actual = isPasswordValid(password)
	if actual != false {
		log.Fatalf("%v: Expected %v, Got: %v", password, false, actual)
	}
}
