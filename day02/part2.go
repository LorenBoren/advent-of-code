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
	compMemory := make([] int, len(intInputArr))
	copy(compMemory, intInputArr)
	// Find the input noun and verb that cause the program to produce the output 19690720
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			intInputArr[1] = noun
			intInputArr[2] = verb
			//fmt.Printf("Tying: %v and %v\n", noun, verb)
			for i := 0; i < len(intInputArr); i=i+4 {
				// the program is finished and should immediately halt
				if intInputArr[i] == 99 {
					if intInputArr[0] == 19690720 {
						fmt.Printf("FOUND IT:  noun: %v verb: %v\n", noun, verb)
					}
					break
				}
				// add
				if intInputArr[i] == 1 {
					sum := intInputArr[intInputArr[i+1]] + intInputArr[intInputArr[i+2]]
					intInputArr[intInputArr[i+3]] = sum
					continue
				}
				// multiply
				if intInputArr[i] == 2 {
					product := intInputArr[intInputArr[i+1]] * intInputArr[intInputArr[i+2]]
					intInputArr[intInputArr[i+3]] = product
					continue
				} else {
					fmt.Printf("SOMETHING WENT WRONG!!! INVADID OPCODE:  %v\n", intInputArr[i])
				}
			}
			// reset the input with the computer's memory
			copy(intInputArr, compMemory)
		}
	}
}
