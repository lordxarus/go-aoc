package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	id     int
	rounds []Round
}

type Round struct {
	red   int
	green int
	blue  int
}

func main() {
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	in, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	// round 1
	// test := Round{
	// 	red:   12,
	// 	blue:  13,
	// 	green: 14,
	// }
	// fmt.Println(test)
	// games := getGames(string(in))
	// idSum := 0
	// for _, game := range games {
	// 	possible := isPossible(game, test.red, test.blue, test.green)
	// 	fmt.Printf("%t %v\n", possible, game)
	// 	if possible {
	// 		idSum += game.id
	// 	}
	// }
	// fmt.Println(idSum)

	// round 2
	games := getGames(string(in))
	power := 0
	for _, game := range games {
		min := minimum(game)
		power += min.red * min.green * min.blue
	}
	fmt.Println(power)
}

func getGames(in string) []Game {
	games := make([]Game, 100)
	split := strings.Split(string(in), "\n")
	for i, line := range split {
		idStr := strings.Replace(strings.Split(line, ":")[0], "Game", "", 1)
		gameId, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err != nil {
			log.Fatal("couldn't parse int: ", err)
		}
		game := Game{
			id: gameId,
		}
		roundStr := strings.Replace(strings.Split(line, ":")[1], " ", "", 1)
		rounds := strings.Split(roundStr, ";")
		for _, roundStr := range rounds {
			round := Round{}
			each := strings.Split(roundStr, ",")
			for _, s := range each {
				if strings.Index(s, " ") == 0 {
					s = strings.Replace(s, " ", "", 1)
				}
				numStr := strings.Split(s, " ")[0]
				num, err := strconv.Atoi(numStr)
				if err != nil {
					log.Fatal("error parsing int: ", err)
				}
				color := strings.Split(s, " ")[1]
				switch color {
				case "red":
					round.red = num
				case "blue":
					round.blue = num
				case "green":
					round.green = num
				default:
					log.Fatalf("%s is not a color", color)
				}
			}
			game.rounds = append(game.rounds, round)
		}
		games[i] = game
	}
	return games
}

// give me a game and I will tell you if it is possible given we were shown red, green, blue number of respective cubes
func isPossible(game Game, red, green, blue int) bool {
	for _, round := range game.rounds {
		if red < round.red {
			return false
		}
		if green < round.green {
			return false
		}
		if blue < round.blue {
			return false
		}
	}
	return true
}

func minimum(game Game) Round {
	ret := Round{}
	for _, round := range game.rounds {
		if round.red > ret.red {
			ret.red = round.red
		}
		if round.green > ret.green {
			ret.green = round.green
		}
		if round.blue > ret.blue {
			ret.blue = round.blue
		}
	}
	return ret
}
