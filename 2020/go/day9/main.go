package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")

	preambleSize := 25
	i := 0
	for {

		//fmt.Println(len(lines[i:preambleSize+i]), lines[i:preambleSize+i])
		//fmt.Println("Looking for:", lines[preambleSize+i], i)
		pair := subsetSum(lines[i:preambleSize+i], mustParseFloat(lines[preambleSize+i]), nil)
		if len(pair) == 0 {
			fmt.Println("Part1:", lines[preambleSize+i])
			break
		}
		i++
	}
	pair := subsetSum2(lines, 36845998, nil)
	fmt.Println("Part2 Pairs:", pair)
	list := make([]int, 0)
	for _, l := range pair {
		list = append(list, mustParseInt(l))
	}
	sort.Sort(sort.IntSlice(list))
	fmt.Println("Part2: ", list[0]+list[len(list)-1])
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

func subsetSum2(numbers []string, target float64, partial []string) []string {

	for i := range numbers {
		for j := i; j < len(numbers); j++ {

			s := sum(numbers[i:j])

			if s == target {
				return numbers[i:j]
			}
			if s > target {
				break
			}

		}
	}
	return nil
}

func subsetSum(numbers []string, target float64, partial []string) []string {
	s := sum(partial)

	if s == target {
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
