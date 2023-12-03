package day2

import (
	"log"
	"os"
	"strconv"
	"strings"
)

type Subset struct {
	red   int
	blue  int
	green int
}

type Game struct {
	gameNumber int
	subsets []Subset
}


func loadInput(fileName string) []Game{
	body, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatalln("Unable to read input file.")
	}

	games := []Game{}

	for i, gameStr := range strings.Split(string(body), "\n") {
		game := Game{i+1, []Subset{}}
		subsets := strings.Split(gameStr, ": ")[1]

		for _, subsetStr := range strings.Split(subsets, "; ") {
			subset := Subset{}

			for _, colorPair := range strings.Split(subsetStr, ", ") {
				numStr := strings.Split(colorPair, " ")[0]
				color := strings.TrimSpace(strings.Split(colorPair, " ")[1])
				num, _ := strconv.Atoi(numStr)

				switch color {
				case "red":
					subset.red += num
				case "green":
					subset.green += num
				case "blue":
					subset.blue += num
				}
			}

			game.subsets = append(game.subsets, subset)
		}

		games = append(games, game)
	}

	return games
}


func RunDayP1() string {
	games := loadInput("input/day2.txt")

	// Given configuration
	totalRed := 12
	totalGreen := 13
	totalBlue := 14
	
	gameNumSum := 0
	
	for _, game := range games {
		possible := true
		for _, subset := range game.subsets {
			if subset.red > totalRed || subset.green > totalGreen || subset.blue > totalBlue {
				possible = false
				break
			}
		}
		if possible {
			gameNumSum += game.gameNumber
		}
	}

	return strconv.Itoa(gameNumSum)
}

func RunDayP2() string {
	games := loadInput("input/day2.txt")

	// Given configuration
	
	sumOfPower := 0
	
	for _, game := range games {
		minRed := 0
		minGreen := 0
		minBlue := 0

		for _, subset := range game.subsets {
			if subset.red > minRed {
				minRed = subset.red
			}
			if subset.green > minGreen {
				minGreen = subset.green
			}
			if subset.blue > minBlue {
				minBlue = subset.blue
			}
		}

		power := minRed * minGreen * minBlue
		sumOfPower += power
	}

	return strconv.Itoa(sumOfPower)
}


