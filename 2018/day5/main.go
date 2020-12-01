package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

// https://adventofcode.com/2018/day/5
var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

func main() {
	flag.Parse()
	lines := readFileToLines("data")
	final := findPair([]rune(lines[0]), ' ')
	fmt.Println("Part1: ", len(final))

	if part2 {
		uniq := make(map[rune]bool)
		sizes := make(map[rune]int)
		for _, r := range strings.ToLower(lines[0]) {
			uniq[r] = true
		}

		for k := range uniq {
			d := findPair([]rune(lines[0]), k)
			sizes[k] = len(d)
		}
		var ss []kv
		for k, v := range sizes {
			ss = append(ss, kv{string(k), v})
		}

		sort.Slice(ss, func(i, j int) bool { // Then sorting the slice by value, higher first.
			return ss[i].Value < ss[j].Value
		})
		for _, kv := range ss[:1] { // Print the x top values
			fmt.Printf("Type:: %v    MinCount: %v\n", kv.Key, kv.Value)
		}

	}

}
func findPair(rr []rune, without rune) []rune {
	if without != ' ' {
		t := string(rr) // simple prepwork in the case of part2
		rU := strings.ToUpper(string(without))
		rL := strings.ToLower(string(without))
		t = strings.Replace(t, rU, "", -1)
		t = strings.Replace(t, rL, "", -1)
		rr = []rune(t)
	}

	for i := 0; i < len(rr)-1; i++ {
		if isRuneOpposite(rr[i], rr[i+1]) {
			rr = append(rr[:i], rr[i+2:]...) // remove the 2 that just matched
			if i > 0 {
				i = i - 2 // can only move back2 when your above 0
			} else {
				i = -1 //your at 0 the next time around
			}
		}
	}
	return rr
}

func isRuneOpposite(r1, r2 rune) bool { // are they the case opposite?

	switch strings.ToUpper(string(r1)) == string(r1) {
	case true:
		return strings.ToLower(string(r1)) == string(r2)
	case false:
		return strings.ToUpper(string(r1)) == string(r2)
	}
	return false
}

type kv struct {
	Key   string
	Value int
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
