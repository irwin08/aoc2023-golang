package day3

import (
	"log"
	"os"
	"strconv"
	"strings"
	"slices"
)

func loadInput(fileName string) [][]string {
	body, err := os.ReadFile(fileName)
	if err != nil {
		path, _ := os.Getwd()
		log.Println(path);
		log.Fatalln("Unable to read input file.")
	}

	inputLines := strings.Split(string(body), "\r\n")
	
	grid := make([][]string, len(inputLines))
	for i := range grid {
		grid[i] = make([]string, len(inputLines[i]))
		for j := range grid[i] {
			grid[i][j] = string(inputLines[i][j])
		}
	}
	
	return grid
}

// returns number, end index
// assumes given a valid starting point
func extractNumberFromIndex(line []string, startIndex int) (int, int) {
	num := 0
	lastIndex := len(line)-1
	
	for i := startIndex; i < len(line); i++ {
		if _, err := strconv.Atoi(string(line[i])); err != nil {
			lastIndex = i-1
			break
		}
	}

	num, _ = strconv.Atoi(strings.Join(line[startIndex:lastIndex+1], ""))
	return num, lastIndex
}

func checkForAdjacentSymbols(grid [][]string, startLine int, startIndex int, endIndex int) bool {
	excludedSymbols := []string{".", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "\n"}
	// check above
	if startLine > 0 {
		startCursor := startIndex
		endCursor := endIndex
		if startIndex > 0 {
			startCursor = startIndex - 1
		}
		if endIndex < len(grid[startLine])-1 {
			endCursor = endIndex + 1
		}

		for i := startCursor; i < endCursor+1; i++ {
			if !slices.Contains(excludedSymbols, grid[startLine-1][i]) {
				myString := ""
				for a := startIndex; a <= endIndex; a++ {
					myString += grid[startLine][a]
				}
				
				return true
			}
		}
	}

	// check below
	if startLine < len(grid)-1 {
		startCursor := startIndex
		endCursor := endIndex
		if startIndex > 0 {
			startCursor = startIndex - 1
		}
		if endIndex < len(grid[startLine])-1 {
			endCursor = endIndex + 1
		}
		
		for i := startCursor; i < endCursor+1; i++ {
			if !slices.Contains(excludedSymbols, grid[startLine+1][i]) {
				myString := ""
				for a := startIndex; a <= endIndex; a++ {
					myString += grid[startLine][a]
				}
				return true
			}
		}
	}

	// check left
	if startIndex > 0 {
		if !slices.Contains(excludedSymbols, grid[startLine][startIndex-1]) {
			myString := ""
				for a := startIndex; a <= endIndex; a++ {
					myString += grid[startLine][a]
				}
			return true
		}
	}

	// check right
	if endIndex < len(grid[startLine])-1 {
		if !slices.Contains(excludedSymbols, grid[startLine][endIndex+1]) {
			myString := ""
				for a := startIndex; a <= endIndex; a++ {
					myString += grid[startLine][a]
				}
			return true
		}
	}

	return false
}


func RunDayP1() string {
	grid := loadInput("C:/Users/Jesse/Documents/Programming/aoc2023-golang/input/day3.txt")

	validSerials := []int{}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if _, err := strconv.Atoi(grid[i][j]); err == nil {
				num, lastIndex := extractNumberFromIndex(grid[i], j)

				if checkForAdjacentSymbols(grid, i, j, lastIndex) {
					validSerials = append(validSerials, num)
				}
				j = lastIndex
			}
		}
	}

	
	sum := 0
	for _, n := range validSerials {
		sum += n
	}
	log.Println(sum)
	return ""
}

func RunDayP2() string {
	return ""
}
