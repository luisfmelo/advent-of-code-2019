package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type Object struct {
	name                  string
	objectsInOrbit        []*Object
	indirectOrbitsMemoize *int
}

func (o *Object) addObjectToOrbit(object *Object) {
	o.objectsInOrbit = append(o.objectsInOrbit, object)
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

func computeNumberOfOrbits(input []string) int {
	objects := map[string]*Object{}

	for _, elem := range input {
		e := strings.Split(elem, ")")
		center := getOrCreateObject(objects, e[0])
		objectInOrbit := getOrCreateObject(objects, e[1])

		center.addObjectToOrbit(objectInOrbit)
	}

	nOrbits := 0
	for _, object := range objects {
		nOrbits += object.getTotalOfOrbits()
	}

	return nOrbits
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
	input := read(inputFilePath)

	nOrbits := computeNumberOfOrbits(input)
	log.Printf("Number of direct/indirect orbits: %d\n", nOrbits)

	elapsed := time.Since(start)
	fmt.Printf("Execution took %s", elapsed)
}
