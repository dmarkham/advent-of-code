package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var seen map[int64]bool // keep track of values we have seen
var sum = int64(0)      //running total

func main() {
	var part2 = flag.Bool("part2", false, "run part2?")
	flag.Parse()
	seen = make(map[int64]bool)
	seen[0] = true
	lines := readFileToLines("data")
	// in part 2 we run this until we hit a dup
	for true {
		for _, l := range lines {
			foo(l)
		}
		if !*part2 {
			break
		}
	}
	fmt.Println("SUM:", sum)
}
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
func foo(line string) {
	switch line[0] {
	case '+':
		//fmt.Println("ADD", line[1:])
		v, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if err != nil {
			panic(err)
		}
		sum = sum + v
	case '-':
		//fmt.Println("Subtract: ", line[1:])
		v, err := strconv.ParseInt(string(line[1:]), 10, 64)
		if err != nil {
			panic(err)
		}
		sum = sum - v
	}
	if seen[sum] {
		fmt.Println("FIRST DUP:", sum)
		os.Exit(0)
	}
	seen[sum] = true
}
