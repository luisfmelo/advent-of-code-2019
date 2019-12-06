package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const MaxInt = int(^uint(0) >> 1)

type Object struct {
	name                  string
	objectsInOrbit        []*Object
	orbit                 *Object
	indirectOrbitsMemoize *int
}

func (o *Object) addObjectToOrbit(object *Object) {
	o.objectsInOrbit = append(o.objectsInOrbit, object)
}

func (o *Object) addOrbitObject(object *Object) {
	o.orbit = object
}

func (o *Object) getNumberOfDirectOrbits() int {
	return len(o.objectsInOrbit)
}

func (o *Object) getNumberOfIndirectOrbits() int {
	if o.indirectOrbitsMemoize != nil {
		return *o.indirectOrbitsMemoize
	}

	nIndirectOrbits := 0
	for _, objectInOrbit := range o.objectsInOrbit {
		nIndirectOrbits += objectInOrbit.getTotalOfOrbits()
	}
	o.indirectOrbitsMemoize = &nIndirectOrbits

	return nIndirectOrbits
}

func (o *Object) getTotalOfOrbits() int {
	return o.getNumberOfDirectOrbits() + o.getNumberOfIndirectOrbits()
}

func getOrCreateObject(objects map[string]*Object, name string) *Object {
	if obj, ok := objects[name]; ok {
		return obj
	}
	object := Object{
		name:           name,
		objectsInOrbit: []*Object{},
	}

	objects[name] = &object
	return &object
}

func getObjectMap(input []string) map[string]*Object {
	objects := map[string]*Object{}

	for _, elem := range input {
		e := strings.Split(elem, ")")
		center := getOrCreateObject(objects, e[0])
		objectInOrbit := getOrCreateObject(objects, e[1])

		center.addObjectToOrbit(objectInOrbit)
		objectInOrbit.addOrbitObject(center)
	}
	return objects
}

func getObjectsToCOM(objects map[string]*Object, start string) map[string]int {
	objectsToCOM := map[string]int{}

	dist := 0
	for ptr := objects[start].orbit.name; objects[ptr].orbit.name != "COM"; ptr = objects[ptr].orbit.name {
		objectsToCOM[ptr] = dist
		dist ++
	}

	return objectsToCOM
}

func computeNumberOfTransfers(input []string, start string, end string) int {
	objects := getObjectMap(input)

	startObjectToCOM := getObjectsToCOM(objects, start)
	endObjectToCOM := getObjectsToCOM(objects, end)

	minDistance := MaxInt
	for objectName, distanceToStart := range startObjectToCOM {
		if distanceToEnd, ok := endObjectToCOM[objectName]; ok {
			if distanceToEnd + distanceToStart < minDistance {
				minDistance = distanceToEnd + distanceToStart
			}
		}
	}
	return minDistance
}

func read(file string) []string {
	f, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer func() { _ = f.Close() }()

	var input []string
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		input = append(input, sc.Text())
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}
	return input
}

func main() {
	start := time.Now()
	inputFilePath := os.Args[1]
	fromObject := os.Args[2]
	toObject := os.Args[3]
	input := read(inputFilePath)

	nOrbits := computeNumberOfTransfers(input, fromObject, toObject)
	log.Printf("Number of direct/indirect orbit: %d\n", nOrbits)

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
