package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	totalFuel := 0
	for scanner.Scan() {
		mass, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		totalFuel += calculateFuel(mass)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	fmt.Printf("Total Fuel Needed: %v\n", totalFuel)
}

func calculateFuel(mass int) int {
	if mass <= 6 {
		return 0
	}
	fuel := (mass/3) - 2

	return fuel + calculateFuel(fuel)
}