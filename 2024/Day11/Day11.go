package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "2 77706 5847 9258441 0 741 883933 12"
	var intSlice []int
	for _, val := range strings.Split(input, " ") {
		intConv, _ := strconv.Atoi(val)
		intSlice = append(intSlice, intConv)
	}
	blinks := 75
	for i := range blinks {
		applyRules(&intSlice)
		fmt.Printf("Iteration %d complete.\n", i + 1)
	}
	fmt.Printf("The length after %d blinks is %d\n", blinks, len(intSlice))
	
}

var memo = make(map[int][]int)
var frequency = make(map[int]int)

func applyRules(intSlice *[]int) {
	newSlice := make([]int, 0, len(*intSlice)*2)
	for _, entry := range *intSlice {
		if cached, found := memo[entry]; found {
			count := frequency[entry]
			for i := 0; i < count; i++ {
				newSlice = append(newSlice, cached...)
			}
		} else {
			var transformed []int
			if entry == 0 {
				transformed = append(transformed, 1)
			} else {
				numDigits := countDigits(entry)
				if numDigits%2 == 0 {
					half1 := entry / intPow(10, numDigits/2)
					half2 := entry % intPow(10, numDigits/2)
					transformed = append(transformed, half1, half2)
				} else {
					transformed = append(transformed, entry*2024)
				}
			}
			memo[entry] = transformed
			frequency[entry]++
			newSlice = append(newSlice, transformed...)
		}
	}
	*intSlice = newSlice
}

func countDigits(n int) int {
	count := 0
	for n > 0 {
		count++
		n /= 10
	}
	return count
}

func intPow(base, exp int) int {
	result := 1
	for i := 0; i < exp; i++ {
		result *= base
	}
	return result
}