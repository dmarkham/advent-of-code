package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")

	list := subsetSum(lines, 2020, nil)
	fmt.Println(list)
	fmt.Println(int(product(list)))
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

func subsetSum(numbers []string, target float64, partial []string) []string {
	s := sum(partial)
	match := 2
	if part2 {
		match = 3
	}

	if s == target && len(partial) == match {
		return partial
	}
	if s > target {
		return nil
	}

	for i := range numbers {
		n := numbers[i]
		r := subsetSum(numbers[i+1:], target, append(partial, n))
		if r != nil {
			return r
		}
	}
	return nil
}
func sum(list []string) float64 {
	sum := 0.0
	for _, part := range list {
		cm, err := strconv.ParseFloat(part, 64)
		if err != nil {
			panic(err)
		}
		sum = sum + cm
	}
	return sum
}

func product(list []string) float64 {
	p := 0.0
	for _, part := range list {
		cm, err := strconv.ParseFloat(part, 64)
		if err != nil {
			panic(err)
		}
		if p == 0 {
			p = cm
		} else {
			p = p * cm
		}
		fmt.Println(part, int(p))
	}
	return p
}
