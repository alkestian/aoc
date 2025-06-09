package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

/*
--- Day 12: JSAbacusFramework.io ---
Santa's Accounting-Elves need help balancing the books after a recent order. Unfortunately, their accounting software uses a peculiar storage format. That's where you come in.

They have a JSON document which contains a variety of things: arrays ([1,2,3]), objects ({"a":1, "b":2}), numbers, and strings. Your first job is to simply find all of the numbers throughout the document and add them together.

For example:

[1,2,3] and {"a":2,"b":4} both have a sum of 6.
[[[3]]] and {"a":{"b":4},"c":-1} both have a sum of 3.
{"a":[-1,1]} and [-1,{"a":1}] both have a sum of 0.
[] and {} both have a sum of 0.
You will not encounter any strings containing numbers.

What is the sum of all numbers in the document?
*/

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error opening input file")
	}
	defer file.Close()
	
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatal("Error reading file")
	}
	
	var origMap interface{}
	err = json.Unmarshal(content, &origMap)
	if err != nil {
		log.Fatal("Error unmarshalling JSON")
	}
	sum := part1(origMap)
	fmt.Println("Sum of all numbers:", sum)
	sum2 := part2(origMap)
	fmt.Println("Sum of all non-red branches:", sum2)
}

func part1(data interface{}) int {
	var sum int

	switch v := data.(type) {
	case float64:
		sum += int(v)
	case []interface{}:
		for _, item := range v {
			sum += part1(item)
		}
	case map[string]interface{}:
		for _, item := range v {
			sum += part1(item)
		}
	}

	return sum
}

/*--- Part Two ---
Uh oh - the Accounting-Elves have realized that they double-counted everything red.

Ignore any object (and all of its children) which has any property with the value "red". Do this only for objects ({...}), not arrays ([...]).

[1,2,3] still has a sum of 6.
[1,{"c":"red","b":2},3] now has a sum of 4, because the middle object is ignored.
{"d":"red","e":[1,2,3,4],"f":5} now has a sum of 0, because the entire structure is ignored.
[1,"red",5] has a sum of 6, because "red" in an array has no effect. */

func part2(data interface{}) int {
	switch v := data.(type) {
	case float64:
		return int(v)
	case []interface{}:
		sum := 0
		for _, item := range v {
			sum += part2(item)
		}
		return sum
	case map[string]interface{}:
		for _, val := range v {
			if str, ok := val.(string); ok && str == "red" {
				return 0
			}
		}

		sum := 0
		for _, item := range v {
			sum += part2(item)
		}
		return sum
	}
	return 0
}



