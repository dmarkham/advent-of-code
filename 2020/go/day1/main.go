package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

type Result struct {
	SomeValue float64
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	// in part 2 we run this until we hit a dup
	for true {

		sum := 0.0
		for _, l := range lines {
			r := processLine(l)
			sum = sum + r.SomeValue
		}
		fmt.Println(int(sum))
		os.Exit(0)
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
func processLine(line string) *Result {
	value, err := strconv.ParseFloat(line, 64)
	if err != nil {
		panic(err)
	}

	if !part2 {
		v := calcFuel(value)
		fmt.Println(line, "-> ", v)
		return &Result{SomeValue: v}
	}

	totalForModule := 0.0
	for {
		v := calcFuel(value)
		totalForModule = totalForModule + v
		if v > 0 {
			value = v
		} else {
			break
		}
	}
	return &Result{SomeValue: totalForModule}
}

func calcFuel(value float64) float64 {
	v := math.Floor(value/3) - 2
	if v < 0 {
		v = 0
	}
	fmt.Println(value, "-> ", v)
	return v
}
