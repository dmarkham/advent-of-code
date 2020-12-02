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

type PassRule struct {
	Min  int
	Max  int
	Char string
	Pass string
}

func (p *PassRule) Valid() bool {

	c := strings.Count(p.Pass, p.Char)
	//fmt.Println("Count: ", c)
	return c >= p.Min && c <= p.Max
}
func (p *PassRule) Valid2() bool {

	if len(p.Pass) < p.Max {
		return false
	}

	if string(p.Pass[p.Min-1]) == p.Char && string(p.Pass[p.Max-1]) != p.Char {
		return true
	}
	if string(p.Pass[p.Min-1]) != p.Char && string(p.Pass[p.Max-1]) == p.Char {
		return true
	}
	return false
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")

	trues := 0
	trues2 := 0
	for _, l := range lines {

		parts := strings.Fields(l)

		r := strings.Split(parts[0], "-")
		min, _ := strconv.ParseInt(r[0], 10, 64)
		max, _ := strconv.ParseInt(r[1], 10, 64)

		//fmt.Println("range:", parts[0])
		//fmt.Println("char:", string(parts[1][0]))
		//fmt.Println("Pass:", parts[2])
		rule := &PassRule{
			Char: string(parts[1][0]),
			Pass: parts[2],
			Min:  int(min),
			Max:  int(max),
		}
		//fmt.Println(rule, rule.Valid())
		if rule.Valid() {
			trues++
		}
		if rule.Valid2() {
			trues2++
		}
	}
	fmt.Println(trues)
	fmt.Println(trues2)
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
