package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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
	spl := strings.Split(string(in), "\n")
	cal := make([]int, len(spl))
	for i, line := range spl {
		var first rune = 0
		var last rune = 0
		for _, c := range line {
			if c >= '0' && c <= '9' {
				if first == 0 {
					first = c
				} else {
					last = c
				}
			}
		}
		if last == 0 {
			last = first
		}
		// fmt.Println(string(first), string(last))
		out := strings.Builder{}
		if first != 0 {
			out.WriteRune(first)
			if last != 0 {
				out.WriteRune(last)
			}
			int, err := strconv.Atoi(out.String())
			if err != nil {
				log.Fatal(err)
			}
			cal[i] = int
		}

	}
	sum := 0
	for _, i := range cal {
		sum += i
	}
	fmt.Println(sum)
}
