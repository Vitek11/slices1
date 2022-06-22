// На вхід подано стрінг з цілими числами, котри розділені пробілами.
//Треба повернути найбільше та найменше число.

package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1 9 3 4 -5 16 -1 -1"
	var result string
	var s = 0

	inputSplit := strings.Split(input, " ")
	minInput, err := strconv.Atoi(inputSplit[0])
	maxInput, _ := strconv.Atoi(inputSplit[0])

	if err != nil {
		fmt.Println("Invalid syntax")
		return
	}

	for i := 1; i < len(inputSplit); i++ {
		s, err = strconv.Atoi(inputSplit[i])

		if err != nil {
			fmt.Println("Invalid syntax")
			return
		}

		if minInput > s {
			minInput = s
		}
		if maxInput < s {
			maxInput = s
		}
	}
	result = strconv.Itoa(maxInput) + " " + strconv.Itoa(minInput)
	fmt.Println(input)
	fmt.Println(result)
}
