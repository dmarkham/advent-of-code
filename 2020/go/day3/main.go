package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

type point struct {
	X int
	Y int
}
type myMap [][]string

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	m := make(myMap, 0)
	for i, line := range lines {
		m = append(m, make([]string, 0))
		for _, c := range line {
			m[i] = append(m[i], string(c))
		}
	}
	p := &point{}
	p1Count := 0
	for {
		//fmt.Println(m[p.X][p.Y])
		p = Move(1, 3, p, m)
		if m[p.X][p.Y] == "#" {
			p1Count++
		}
		//fmt.Println(m[p.X][p.Y])
		if p.X == len(lines)-1 {
			break
		}
	}
	fmt.Println(p1Count)

	slopes := []struct{ X, Y int }{
		{1, 1},
		{1, 3},
		{1, 5},
		{1, 7},
		{2, 1},
	}
	total := 1
	for _, t := range slopes {
		//fmt.Println(t)
		p := &point{}
		temp := 0
		for {
			//fmt.Println(m[p.X][p.Y])
			p = Move(t.X, t.Y, p, m)
			if m[p.X][p.Y] == "#" {
				temp++
			}
			//fmt.Println(m[p.X][p.Y])
			if p.X == len(lines)-1 {
				break
			}
		}
		total = total * temp
	}
	fmt.Println(total)
}

func Move(x, y int, p *point, m myMap) *point {
	p.X = p.X + x
	p.Y = (p.Y + y) % len(m[0])

	return p
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
