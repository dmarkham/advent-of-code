package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

type Intruction struct {
	Op  string
	Val int
	Ran int
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	intructions := make([]Intruction, 0)
	for _, l := range lines {

		parts := strings.Fields(l)
		in := Intruction{Op: parts[0], Val: mustParseInt(parts[1])}
		intructions = append(intructions, in)
	}

	// only needed for part 2
	try := 0
	cc := make([]Intruction, len(intructions))
	copy(cc, intructions)
	// only needed for part 2

	for {
		entry := 0
		count := 0

		if try > len(intructions) {
			panic("try too big")
		}
		if part2 {

			if intructions[try].Op == "nop" {
				intructions[try].Op = "jmp"
			} else if intructions[try].Op == "jmp" {
				intructions[try].Op = "nop"
			}
		}

		for {
			if entry >= len(intructions) {
				fmt.Println("Part2:", count, "Flipped line:", try) // part 2 we are looking for a clean exit
				os.Exit(0)
			}
			if intructions[entry].Ran > 0 {
				if !part2 {
					fmt.Println("Part1:", count) // part 1 we just detect any loop
					os.Exit(0)
				}
				break
			}

			intructions[entry].Ran++

			if intructions[entry].Op == "nop" {
				entry++
				continue
			}
			if intructions[entry].Op == "acc" {
				count = count + intructions[entry].Val
				entry++
				continue
			}
			if intructions[entry].Op == "jmp" {
				entry = entry + intructions[entry].Val
				continue
			}
		}
		// only needed for part 2
		copy(intructions, cc)
		try++
	}
	//fmt.Println("Done:", entry, count, intructions[entry])
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
