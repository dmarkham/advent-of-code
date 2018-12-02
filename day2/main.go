package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/agnivade/levenshtein"
)

var sum = int64(0) //running total
var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

func main() {
	flag.Parse()
	lines := readFileToLines("data")
	if part2 {
		for _, l := range lines {
			for _, l2 := range lines {
				distance := levenshtein.ComputeDistance(l, l2)
				if distance == 1 {
					fmt.Println(l)
					fmt.Println(l2)
					os.Exit(0)
				}
			}
		}
		os.Exit(1)
	}

	pair := 0
	three := 0
	for _, l := range lines {
		hits := make(map[rune]int)

		for _, r := range l {
			hits[r]++
		}
		hadTwo := false
		hadThree := false
		for _, v := range hits {
			if v == 2 {
				hadTwo = true
			}
			if v == 3 {
				hadThree = true
			}

		}
		if hadTwo {
			pair++
		}
		if hadThree {
			three++
		}
	}
	fmt.Println(pair, three, pair*three)

}

// Pull all lines into a string slice
func readFileToLines(file string) []string {
	// open data
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	r := bufio.NewReader(fh)
	scanner := bufio.NewScanner(r)
	lines := make([]string, 0)
	// read it all in
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
