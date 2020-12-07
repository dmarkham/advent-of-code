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

type Bag struct {
	Name   string
	Amount int
	Holds  []*Bag
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	bags := make([]*Bag, 0)
	for _, l := range lines {
		l = strings.Replace(l, "contain", ",", -1)
		l = strings.Replace(l, ", no other bags.", "", -1)
		l = strings.Replace(l, "bags", "", -1)
		l = strings.Replace(l, "bag", "", -1)
		l = strings.Replace(l, ".", "", -1)

		r := strings.Split(l, ",")
		b := &Bag{Name: strings.TrimSpace(r[0]), Holds: toBags(r[1:])}
		bags = append(bags, b)
	}
	count1 := 0
	bgs := holds("shiny gold", bags)
	seen := make(map[string]bool)
	for _, b := range bgs {
		if !seen[b.Name] {
			count1++
		}
		seen[b.Name] = true
	}

	fmt.Println("Part1:", count1)
	fmt.Println("Part2:", contains2("shiny gold", bags))
}

func contains2(name string, bags []*Bag) int {
	h := 0
	for _, b := range bags {
		if b.Name == name {

			for _, has := range b.Holds {
				temp := contains2(has.Name, bags)
				if temp == 0 {
					h += has.Amount
				} else {
					h = h + has.Amount + (has.Amount * temp)
				}
			}
		}
	}
	return h
}

func holds(name string, bags []*Bag) []*Bag {
	h := make([]*Bag, 0)
	for _, b := range bags {
		matched := false
		for _, c := range b.Holds {
			if c.Name == name {
				matched = true
			}
		}
		if matched {
			h = append(h, b)
			h = append(h, holds(b.Name, bags)...)

		}
	}
	return h
}

func toBags(s []string) []*Bag {
	bags := make([]*Bag, 0)
	for _, s := range s {
		parts := strings.Fields(s)

		b := &Bag{Name: strings.Join(parts[1:], " "), Amount: mustParseInt(parts[0])}
		//fmt.Println("BAG: ", b)
		bags = append(bags, b)
	}
	return bags
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
