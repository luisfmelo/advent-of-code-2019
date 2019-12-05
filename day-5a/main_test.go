package main

import (
	"log"
	"math"
	"reflect"
	"testing"
)

func TestComputeIntCodeSequence(t *testing.T) {

	output := make(chan int)
	input := make(chan int)
	intCode := []int{4, 2, 99}
	expected := 99
	go computeIntCodeSequence(intCode, input, output)

	actual := 0
	idx := 0
	for o := range output {
		actual += int(math.Pow10(idx)) * o
		idx++
	}
	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}

func TestGetValueByMode(t *testing.T) {
	mode := POSITION
	intCode := []int{4, 2, 99}
	idx := 1
	expected := 99
	actual := getValueByMode(intCode, idx, mode)

	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}

	mode = IMMEDIATE
	intCode = []int{4, 2, 99}
	idx = 1
	expected = 2
	actual = getValueByMode(intCode, idx, mode)

	if actual != expected {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}

func TestNewInstruction(t *testing.T) {
	inst := 1002
	expected := instruction{
		opCode:          2,
		firstParamMode:  0,
		secondParamMode: 1,
		thirdParamMode:  0,
	}
	actual := newInstruction(inst)

	if !reflect.DeepEqual(actual, expected) {
		log.Fatalf("Expected %v, Got: %v", expected, actual)
	}
}

