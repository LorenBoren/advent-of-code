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

	input, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	// read in the entire input
	var strInputArr []string
	for scanner.Scan() {
		strInputArr = strings.Split(scanner.Text(), ",")
	}
	// convert the string slice into a slice of ints
	intInputArr := make([]int, len(strInputArr))
	for i, strInput := range strInputArr {
		intInputArr[i], err = strconv.Atoi(strInput)
		if err != nil {
			fmt.Printf("Error converting string slice to int slice")
		}
	}
	fmt.Printf("intArray: %v\n", intInputArr)
	for i := 0; i < len(intInputArr); i=i+4 {
		// the program is finished and should immediately halt
		if intInputArr[i] == 99 {
			fmt.Printf("Value at 0th element:  %v\n", intInputArr[0])
			break
		}
		// add
		if intInputArr[i] == 1 {
			fmt.Printf("adding: %v + %v to get %v\n", intInputArr[intInputArr[i+1]], intInputArr[intInputArr[i+2]], intInputArr[intInputArr[i+1]] + intInputArr[intInputArr[i+2]])
			intInputArr[intInputArr[i+3]] = intInputArr[intInputArr[i+1]] + intInputArr[intInputArr[i+2]]
			continue
		}
		// multiply
		if intInputArr[i] == 2 {
			fmt.Printf("multiplying: %v + %v to get %v\n", intInputArr[intInputArr[i+1]], intInputArr[intInputArr[i+2]], intInputArr[intInputArr[i+1]] * intInputArr[intInputArr[i+2]])
			intInputArr[intInputArr[i+3]] = intInputArr[intInputArr[i+1]] * intInputArr[intInputArr[i+2]]
		}
	}
}
