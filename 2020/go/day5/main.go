package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

type Range struct {
	Min float64
	Max float64
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	part1 := 0.0
	seats := make([]int, 0)
	for _, l := range lines {

		//fmt.Println("Line:", l)
		val := Range{1, 128}
		final := 0.0
		r := 0.0
		for i, d := range l {
			//fmt.Println("Char:", string(d))
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
			//fmt.Println("Val:", val, "Final: ", final)
		}
		//fmt.Println("Line:", l)
		//parts := strings.Fields(l)

		//r := strings.Split(parts[0], "-")

		//fmt.Println("range:", parts[0])
		//fmt.Println("char:", string(parts[1][0]))
		//fmt.Println("Pass:", parts[2])
		//rule := &PassRule{ }
		//fmt.Println(rule, rule.Valid())
	}

	fmt.Println("Part1:", part1)
	sort.Sort(sort.IntSlice(seats))
	last := 0
	for _, s := range seats {
		//fmt.Println(s, last)
		if last != s {
			fmt.Println("YOUR SEAT:", s)
			last = s
		} else {
			fmt.Println(s, last)

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

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert string %s to integer: %v", s, err)
	}
	return i
}
func mustParseFloat(s string) float64 {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Fatalf("cannot convert string %s to float64: %v", s, err)
	}
	return i
}

func Paragraph(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of the separator substring
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + len("\n\n"), bytes.ReplaceAll(data[0:i], []byte{'\n'}, []byte{' '}), nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), bytes.ReplaceAll(data, []byte{'\n'}, []byte{' '}), nil
	}

	return
}
