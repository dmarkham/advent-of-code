package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
)

type Range struct {
	Min float64
	Max float64
}

func main() {
	lines := readFileToLines("data.txt")
	part1 := 0.0
	seats := make([]int, 0)
	for _, l := range lines {

		val := Range{1, 128}
		final := 0.0
		r := 0.0
		for i, d := range l {
			if d == 'F' || d == 'L' {
				val.Max -= math.Round((val.Max - val.Min) / 2)
			} else {
				val.Min += math.Round((val.Max - val.Min) / 2)
			}
			if i == 6 {
				if d == 'F' {
					final = val.Min - 1
				} else {
					final = val.Max - 1
				}
				val = Range{1, 8}
			}
			if i == 9 {
				if d == 'L' {
					r = val.Min - 1
				} else {
					r = val.Max - 1
				}
				final = final*8.0 + r
				seats = append(seats, int(final))
				if final > part1 {
					part1 = final
				}
			}
		}
	}

	fmt.Println("Part1:", part1)
	sort.Sort(sort.IntSlice(seats))
	last := 0
	for i, s := range seats {
		if last != s {
			if i != 0 {
				fmt.Println("Part2 YOUR SEAT:", s-1)
			}
			last = s
		}
		last++
	}
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

func mustParseFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("cannot convert string %s to float64: %v", s, err)
	}
	return i
}
