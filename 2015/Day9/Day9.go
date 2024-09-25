package main

import (
	"fmt"
	"log"
	"math"
	"os"
)

func permute(cities []string, start int, result *[][]string) {
	if start == len(cities)-1 {
		route := make([]string, len(cities))
		copy(route, cities)
		*result = append(*result, route)
		return
	}

	for i := start; i < len(cities); i++ {
		cities[start], cities[i] = cities[i], cities[start]
		permute(cities, start+1, result)
		cities[start], cities[i] = cities[i], cities[start]
	}
}

func calculateDistance(route []string, distances map[string]map[string]int) int {
	totalDistance := 0
	for i := 0; i < len(route)-1; i++ {
		totalDistance += distances[route[i]][route[i+1]]
	}
	return totalDistance
}

func findShortestRoute(cities []string, distances map[string]map[string]int) int {
	var permutations [][]string
	permute(cities, 0, &permutations)

	shortestDistance := math.MaxInt
	for _, route := range permutations {
		distance := calculateDistance(route, distances)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}
	return shortestDistance
}

func findLongestRoute(cities []string, distances map[string]map[string]int) int {
	var permutations [][]string
	permute(cities, 0, &permutations)

	longestDistance := 0
	for _, route := range permutations {
		distance := calculateDistance(route, distances)
		if distance > longestDistance {
			longestDistance = distance
		}
	}
	return longestDistance
}

func main() {
	_, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}	
	
	distances := map[string]map[string]int{
		"Faerun":        {"Tristram": 65, "Tambi": 129, "Norrath": 144, "Snowdin": 71, "Straylight": 137, "AlphaCentauri": 3, "Arbre": 149},
		"Tristram":      {"Faerun": 65, "Tambi": 63, "Norrath": 4, "Snowdin": 105, "Straylight": 125, "AlphaCentauri": 55, "Arbre": 14},
		"Tambi":         {"Faerun": 129, "Tristram": 63, "Norrath": 68, "Snowdin": 52, "Straylight": 65, "AlphaCentauri": 22, "Arbre": 143},
		"Norrath":       {"Faerun": 144, "Tristram": 4, "Tambi": 68, "Snowdin": 8, "Straylight": 23, "AlphaCentauri": 136, "Arbre": 115},
		"Snowdin":       {"Faerun": 71, "Tristram": 105, "Tambi": 52, "Norrath": 8, "Straylight": 101, "AlphaCentauri": 84, "Arbre": 96},
		"Straylight":    {"Faerun": 137, "Tristram": 125, "Tambi": 65, "Norrath": 23, "Snowdin": 101, "AlphaCentauri": 107, "Arbre": 14},
		"AlphaCentauri": {"Faerun": 3, "Tristram": 55, "Tambi": 22, "Norrath": 136, "Snowdin": 84, "Straylight": 107, "Arbre": 46},
		"Arbre":         {"Faerun": 149, "Tristram": 14, "Tambi": 143, "Norrath": 115, "Snowdin": 96, "Straylight": 14, "AlphaCentauri": 46},
	}

	cities := []string{"Faerun", "Tristram", "Tambi", "Norrath", "Snowdin", "Straylight", "AlphaCentauri", "Arbre"}

	shortest := findShortestRoute(cities, distances)
	longest := findLongestRoute(cities, distances)

	fmt.Println("Shortest Route Distance:", shortest)
	fmt.Println("Longest Route Distance:", longest)
}