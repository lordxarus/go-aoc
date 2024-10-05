package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

var nums = []string{
	"zero",
	"one",
	"two",
	"three",
	"four",
	"five",
	"six",
	"seven",
	"eight",
	"nine",
	"ten",
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
	spl := strings.Split(string(in), "\n")
	parsed := parse(spl)
	sum := 0
	for _, i := range parsed {
		if i != -1 {
			sum += i
		}
	}
	fmt.Println(sum)
}

// this returns a number comprised of the first and last numerals found in the input string
// after being run through toNumeral
// ex: 1nine9eight would return 18
func firstAndLast(in string) int {
	line := toNumeral(in)
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
		return int
	}
	return -1
}

func parse(in []string) []int {
	ret := make([]int, len(in))
	for i, l := range in {
		parsed := firstAndLast(l)
		if parsed == -1 {
			fmt.Printf("parsed returned -1 for line %s\n", l)
		}
		ret[i] = parsed
	}
	return ret
}

func toNumeral(in string) string {
	idx := make([]string, len(in))
	// if we have any numerals already, add them into the arrays
	for i, c := range in {
		_, err := strconv.Atoi(string(c))
		if err == nil {
			idx[i] = string(c)
		}
	}
	for _, strNum := range nums {

		firstIdx := strings.Index(in, strNum)

		if firstIdx == -1 {
			continue
		}
		idx[firstIdx] = strNum

		lastIdx := strings.LastIndex(in, strNum)
		if firstIdx != lastIdx {
			idx[lastIdx] = strNum
		}

		idx[firstIdx] = strNum
	}

	ret := ""
	for _, s := range idx {
		if s != "" {
			ret += s
		}
	}
	for i, strNum := range nums {
		ret = strings.ReplaceAll(ret, strNum, strconv.Itoa(i))
	}

	return ret
}
