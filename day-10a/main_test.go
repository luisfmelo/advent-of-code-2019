package main

import (
	"log"
	"testing"
)

func TestGetBestMonitoringStationLocation(t *testing.T) {
	input := []string{
		"......#.#.",
		"#..#.#....",
		"..#######.",
		".#.#.###..",
		".#..#.....",
		"..#....#.#",
		"#..#....#.",
		".##.#..###",
		"##...#..#.",
		".#....####",
	}
	expectedCoordinate := coordinate{x: 5, y: 8}
	expectedAsteroidsDetected := 33
	actualCoordinate, actualAsteroidsInSight := getBestMonitoringStationLocation(input)
	if expectedCoordinate.toString() != actualCoordinate.toString() {
		log.Fatalf("Expected %v, Got: %v", expectedCoordinate.toString(), actualCoordinate.toString())
	}
	if expectedAsteroidsDetected != actualAsteroidsInSight {
		log.Fatalf("Expected %v, Got: %v", expectedAsteroidsDetected, actualAsteroidsInSight)
	}

	input = []string{
		"#.#...#.#.",
		".###....#.",
		".#....#...",
		"##.#.#.#.#",
		"....#.#.#.",
		".##..###.#",
		"..#...##..",
		"..##....##",
		"......#...",
		".####.###.",
	}
	expectedCoordinate = coordinate{x: 1, y: 2}
	expectedAsteroidsDetected = 35
	actualCoordinate, actualAsteroidsInSight = getBestMonitoringStationLocation(input)
	if expectedCoordinate.toString() != actualCoordinate.toString() {
		log.Fatalf("Expected %v, Got: %v", expectedCoordinate.toString(), actualCoordinate.toString())
	}
	if expectedAsteroidsDetected != actualAsteroidsInSight {
		log.Fatalf("Expected %v, Got: %v", expectedAsteroidsDetected, actualAsteroidsInSight)
	}

	input = []string{
		".#..#..###",
		"####.###.#",
		"....###.#.",
		"..###.##.#",
		"##.##.#.#.",
		"....###..#",
		"..#.#..#.#",
		"#..#.#.###",
		".##...##.#",
		".....#.#..",
	}
	expectedCoordinate = coordinate{x: 6, y: 3}
	expectedAsteroidsDetected = 41
	actualCoordinate, actualAsteroidsInSight = getBestMonitoringStationLocation(input)
	if expectedCoordinate.toString() != actualCoordinate.toString() {
		log.Fatalf("Expected %v, Got: %v", expectedCoordinate.toString(), actualCoordinate.toString())
	}
	if expectedAsteroidsDetected != actualAsteroidsInSight {
		log.Fatalf("Expected %v, Got: %v", expectedAsteroidsDetected, actualAsteroidsInSight)
	}

	input = []string{
		".#..##.###...#######",
		"##.############..##.",
		".#.######.########.#",
		".###.#######.####.#.",
		"#####.##.#.##.###.##",
		"..#####..#.#########",
		"####################",
		"#.####....###.#.#.##",
		"##.#################",
		"#####.##.###..####..",
		"..######..##.#######",
		"####.##.####...##..#",
		".#####..#.######.###",
		"##...#.##########...",
		"#.##########.#######",
		".####.#.###.###.#.##",
		"....##.##.###..#####",
		".#.#.###########.###",
		"#.#.#.#####.####.###",
		"###.##.####.##.#..##",
	}
	expectedCoordinate = coordinate{x: 11, y: 13}
	expectedAsteroidsDetected = 210
	actualCoordinate, actualAsteroidsInSight = getBestMonitoringStationLocation(input)
	if expectedCoordinate.toString() != actualCoordinate.toString() {
		log.Fatalf("Expected %v, Got: %v", expectedCoordinate.toString(), actualCoordinate.toString())
	}
	if expectedAsteroidsDetected != actualAsteroidsInSight {
		log.Fatalf("Expected %v, Got: %v", expectedAsteroidsDetected, actualAsteroidsInSight)
	}
}
