package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func transpose(lines []string) []string {
	rowLen := len(lines[0])
	colLen := len(lines)
	result := make([]string, rowLen)

	for i := 0; i < rowLen; i++ {
		var column []byte
		for j := 0; j < colLen; j++ {
			column = append(column, lines[j][i])
		}
		result[i] = string(column)
	}

	return result
}

func countOccurrences(lines []string, target string) int {
	count := 0
	reversed := reverseString(target)
	for _, line := range lines {
		count += strings.Count(line, target)  
		count += strings.Count(line, reversed) 
	}
	return count
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func extractDiagonals(lines []string) ([]string, []string) {
	n, m := len(lines), len(lines[0])
	var diag1, diag2 []string

	for d := 0; d < n+m-1; d++ {
		var sb1, sb2 strings.Builder
		for i := 0; i < n; i++ {
			j1 := d - i
			if j1 >= 0 && j1 < m {
				sb1.WriteByte(lines[i][j1])
			}
			j2 := i - d + m - 1
			if j2 >= 0 && j2 < m {
				sb2.WriteByte(lines[i][j2])
			}
		}
		if sb1.Len() > 0 {
			diag1 = append(diag1, sb1.String())
		}
		if sb2.Len() > 0 {
			diag2 = append(diag2, sb2.String())
		}
	}

	return diag1, diag2
}

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal("Error parsing input")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Printf("There were %d occurences in part 1.\n", part1(lines))
	fmt.Printf("There were %d occurences in part 2.\n", part2(lines))
}

func part1(lines []string) int {
	horizontalMatches := countOccurrences(lines, "XMAS")

	transposed := transpose(lines)
	verticalMatches := countOccurrences(transposed, "XMAS")

	diag1, diag2 := extractDiagonals(lines)
	diagonalMatches := countOccurrences(diag1, "XMAS") + countOccurrences(diag2, "XMAS")

	totalMatches := horizontalMatches + verticalMatches + diagonalMatches
	return totalMatches
}

func part2(lines []string) int {
	xmasMatches := 0
	rowLen := len(lines[0])
	colLen := len(lines)

	for x := 1; x < colLen - 1; x++ {
		for y := 1; y < rowLen - 1; y++ {
			var d1, d2 strings.Builder
			if lines[x][y] == 'A' {
				d1.WriteByte(lines[x-1][y-1])
				d1.WriteByte(lines[x][y])
				d1.WriteByte(lines[x+1][y+1])
				d2.WriteByte(lines[x-1][y+1])
				d2.WriteByte(lines[x][y])
				d2.WriteByte(lines[x+1][y-1])			
			}
			if (d1.String() == "SAM" || d1.String() ==  "MAS") && (d2.String() == "SAM" || d2.String() == "MAS") {
				xmasMatches++
			}
		}
	}

	return xmasMatches
}