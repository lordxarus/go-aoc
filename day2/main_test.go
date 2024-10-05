package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

func SumIdCheck(t *testing.T) {
	want := 2563
	// should really put this into a func but I'm being lazy and you should be happy I'm writing a test!
	file, err := os.Open("in.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	in, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	test := Round{
		red:   12,
		blue:  13,
		green: 14,
	}

	games := getGames(string(in))
	got := 0
	for _, game := range games {
		possible := isPossible(game, test.red, test.blue, test.green)
		fmt.Printf("%t %v\n", possible, game)
		if possible {
			got += game.id
		}
	}

	if got != want {
		t.Fatalf("failed sum id test, want: %d got: %d", want, got)
	}

}
