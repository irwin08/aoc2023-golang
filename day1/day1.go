package day1

import (
	"fmt"
	"os"
	"log"
	"strings"
	"unicode"
	"strconv"
)

func loadInput(fileName string) []string {
	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Unable to read input file.")
	}
	return strings.Split(string(body), "\n")
}

func RunDay1() int {
	inputLines := loadInput("input/day1.txt")
	
	runningCount := 0

	for _, line := range inputLines {
		set := false
		var first int
		var last int
		
		for _, c := range line {
			if unicode.IsNumber(c) {
				conv, _ := strconv.Atoi(string(c))
				first = conv
			}
		}

		for _, c := range line {
			if unicode.IsNumber(c) {
				conv, _ := strconv.Atoi(string(c))
				first = conv
				break
			}
		}

		for i := len(line)-1; i >= 0; i-- {
			if unicode.IsNumber(rune(line[i])) {
				conv, _ := strconv.Atoi(string(line[i]))
				last = conv
				set = true
				break
			}
		}
		
		if set {
			conv, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
			runningCount += conv
		}
	}

	return runningCount
}
