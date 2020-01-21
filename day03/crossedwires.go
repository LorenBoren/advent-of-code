package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// read it in
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	var line1 []string
	var line2 []string
	scanner.Scan()
	line1 = strings.Split(scanner.Text(), ",")
	scanner.Scan()
	line2 = strings.Split(scanner.Text(), ",")
	fmt.Printf("Line 1: %v\n", line1)
	fmt.Printf("Line 2: %v\n", line2)
	// range through each line and add each coordinate the line
	// passes through to the list of coordinates for that line

	var line1CoordList []Point
	line1CoordList = append(line1CoordList, Point{0,0})
	for _, opCode := range line1 {
		line1CoordList = addCoordinatesToListFromOpCode(line1CoordList, opCode)
	}
	var line2CoordList []Point
	line2CoordList = append(line2CoordList, Point{0,0})
	for _, opCode := range line2 {
		line2CoordList = addCoordinatesToListFromOpCode(line2CoordList, opCode)
	}

	// find common coordinate and compare distance
	var shortestDist int
	var closestPoint Point
	var stepsToIntersection int
	for i, point1 := range line1CoordList {
		for j, point2 := range line2CoordList {
			if point1.x == point2.x && point1.y == point2.y {
				if point1.x == 0 && point2.x == 0 && point1.y == 0 && point2.y == 0 {
					continue
				}
				fmt.Printf("Found intersection at %v \n", point1)
				distance := Abs(point1.x) + Abs(point1.y)
				fmt.Printf("Found distance: %v \n", distance)
				if stepsToIntersection == 0 {
					stepsToIntersection = i+j
					fmt.Printf("Total steps to first intersection: %v\n", stepsToIntersection)
				}
				if shortestDist == 0 {
					shortestDist = distance
				}
				if distance < shortestDist {
					shortestDist = distance
					closestPoint = point1
				}

			}
		}
	}

	fmt.Printf("Shortest Manhattan Distance: %v, closest point: %v\n", shortestDist, closestPoint)
}


func addCoordinatesToListFromOpCode(lineCoordList []Point, opCode string) []Point {
	//   x,y
	// R/L, U/D
	// [Point{0,0}, Point{0,1}]
	var numOfPositionsToMove int

	if strings.HasPrefix(opCode, "R") {
		numOfPositionsToMove, _ = strconv.Atoi(strings.TrimPrefix(opCode, "R"))
		for i := 0; i < numOfPositionsToMove; i++ {
			coordsToAdd := Point{lineCoordList[len(lineCoordList)-1].x + 1, lineCoordList[len(lineCoordList)-1].y}
			lineCoordList = append(lineCoordList, coordsToAdd)
		}
	}
	if strings.HasPrefix(opCode, "L") {
		numOfPositionsToMove, _ = strconv.Atoi(strings.TrimPrefix(opCode, "L"))
		for i := 0; i < numOfPositionsToMove; i++ {
			coordsToAdd := Point{lineCoordList[len(lineCoordList)-1].x - 1, lineCoordList[len(lineCoordList)-1].y}
			lineCoordList = append(lineCoordList, coordsToAdd)

		}
	}
	if strings.HasPrefix(opCode, "U") {
		numOfPositionsToMove, _ = strconv.Atoi(strings.TrimPrefix(opCode, "U"))
		for i := 0; i < numOfPositionsToMove; i++ {
			coordsToAdd := Point{lineCoordList[len(lineCoordList)-1].x, lineCoordList[len(lineCoordList)-1].y + 1}
			lineCoordList = append(lineCoordList, coordsToAdd)

		}
	}
	if strings.HasPrefix(opCode, "D") {
		numOfPositionsToMove, _ = strconv.Atoi(strings.TrimPrefix(opCode, "D"))
		for i := 0; i < numOfPositionsToMove; i++ {
			coordsToAdd := Point{lineCoordList[len(lineCoordList)-1].x, lineCoordList[len(lineCoordList)-1].y - 1}
			lineCoordList = append(lineCoordList, coordsToAdd)

		}
	}
	return lineCoordList
}

type Point struct {
	x int
	y int
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
