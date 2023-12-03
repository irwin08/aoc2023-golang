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

func checkForNumber(s string, n int) (int, bool) {
	wordDigits := map[string]int { "zero": 0, "one" : 1, "two": 2, "three": 3, "four": 4, "five": 5, "six": 6, "seven": 7, "eight": 8, "nine": 9 }

	if unicode.IsNumber(rune(s[n])) {
		conv, _ := strconv.Atoi(string(rune(s[n])))
		return conv, true
	} else {
		for w, d := range wordDigits {
			if len(s)-n < len(w) {
				continue
			}

			skip := false

			for i := 0; i < len(w); i++ {
				if s[n+i] != w[i]{
					skip = true
					break
				}
			}
			if skip {
				continue
			}

			return d, true
		}
		return 0, false
	}
}

func RunDayP1() string {
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
	
	return strconv.Itoa(runningCount)
}

func RunDayP2() string {
	inputLines := loadInput("input/day1.txt")
	
	runningCount := 0

	for _, line := range inputLines {
		var first int
		var last int

		for i := range line {
			num, exists := checkForNumber(line, i)
				
			if exists {
				first = num
				break
			}
		}

		for i := len(line)-1; i >= 0; i-- {
			num, exists := checkForNumber(line, i)

			if exists {
				last = num
				break
			}
		}
		
		conv, _ := strconv.Atoi(fmt.Sprintf("%d%d", first, last))
		runningCount += conv
	}

	return strconv.Itoa(runningCount)
}
