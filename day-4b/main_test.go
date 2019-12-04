package main

import (
	"log"
	"testing"
)

func TestIsPasswordValid(t *testing.T) {

	password := 112233
	actual := isPasswordValid(password)
	if actual != true {
		log.Fatalf("%v: Expected %v, Got: %v", password, true, actual)
	}

	password = 123444
	actual = isPasswordValid(password)
	if actual != false {
		log.Fatalf("%v: Expected %v, Got: %v", password, false, actual)
	}

	password = 111122
	actual = isPasswordValid(password)
	if actual != true {
		log.Fatalf("%v: Expected %v, Got: %v", password, true, actual)
	}
}
