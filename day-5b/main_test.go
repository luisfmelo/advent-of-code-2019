package main

import (
	"log"
	"reflect"
	"testing"
)

func fillInput(arr []int, inputCh chan<- int) {
	for _, in := range arr {
		inputCh <- in
	}
}

func TestComputeIntCodeSequence(t *testing.T) {

	// Using position mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
	{
		inputCh := make(chan int)
		intCode := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
		expected := 1
		go fillInput([]int{8}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
		expected := 0
		go fillInput([]int{4}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}

	// Using position mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
	{
		inputCh := make(chan int)
		intCode := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
		expected := 1
		go fillInput([]int{4}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
		expected := 0
		go fillInput([]int{10}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}

	// Using immediate mode, consider whether the input is equal to 8; output 1 (if it is) or 0 (if it is not).
	{
		inputCh := make(chan int)
		intCode := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
		expected := 1
		go fillInput([]int{8}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 3, 1108, -1, 8, 3, 4, 3, 99}
		expected := 0
		go fillInput([]int{4}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}

	// Using immediate mode, consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not).
	{
		inputCh := make(chan int)
		intCode := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
		expected := 1
		go fillInput([]int{6}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
		expected := 0
		go fillInput([]int{10}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}

	// Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero: POSITION MODE
	{
		inputCh := make(chan int)
		intCode := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
		expected := 0
		go fillInput([]int{0}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
		expected := 1
		go fillInput([]int{10}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}

	// Here are some jump tests that take an input, then output 0 if the input was zero or 1 if the input was non-zero: IMMEDIATE MODE
	{
		inputCh := make(chan int)
		intCode := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
		expected := 0
		go fillInput([]int{0}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
		expected := 1
		go fillInput([]int{10}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}

	// Here's a larger example:
	// The example below uses an input instruction to ask for a single number.
	// The program will then output 999 if the input value is below 8,
	// output 1000 if the input value is equal to 8, or output 1001 if the input value is greater than 8.
	{
		inputCh := make(chan int)
		intCode := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0,
			1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
		}
		expected := 999
		go fillInput([]int{7}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0,
			1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
		}
		expected := 1000
		go fillInput([]int{8}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
	}
	{
		inputCh := make(chan int)
		intCode := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0,
			1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99,
		}
		expected := 1001
		go fillInput([]int{9}, inputCh)
		outputCh := make(chan int)
		go computeIntCodeSequence(intCode, inputCh, outputCh)

		actual := getOutput(outputCh)
		if actual != expected {
			log.Fatalf("Expected %v, Got: %v", expected, actual)
		}
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
