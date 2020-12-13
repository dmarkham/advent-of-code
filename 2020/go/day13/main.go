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

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")

	time := mustParseInt(lines[0])
	busss := strings.Split(lines[1], ",")
	ids := make([]uint64, 0)

	for _, b := range busss {
		if b == "x" {
			ids = append(ids, 0)
		}
		if b != "x" {
			ids = append(ids, mustParseInt(b))
		}
	}

	//fmt.Println(time, ids)
	part1 := uint64(0)
	tick := uint64(time)
DONE:
	for {
		for _, b := range ids {
			if b == 0 {
				continue
			}
			if tick%b == 0 {
				part1 = (tick - time) * b
				break DONE
			}
		}
		tick++
	}
	fmt.Println("Part1:", part1)
	inc := uint64(1)

	// this is the first rev of part2
	// it does work it is just slower than is should be,
	// TODO: Chinese Remainder Theorem

	nextMul := map[uint64]uint64{
		37:  1,
		41:  37,
		433: 41,
		23:  443,
		17:  23,
		19:  17,
		29:  19,
		593: 29,
		13:  593,
	}

	tick = uint64(0)
LOOP:
	for {
		for i, b := range ids {
			if b == 0 {
				continue
			}
			if (tick+uint64(i))%b != 0 {
				tick += inc
				continue LOOP
			} else {
				if b*nextMul[b] > inc {

					inc = b * nextMul[b]
					fmt.Printf("TickCheck: %v+%v=%v mod %v  =%v   INC:%v \n", tick, uint64(i), tick+uint64(i), b, (tick+uint64(i))%b, inc)
				}

			}
		}
		fmt.Println("Part2", tick)
		break
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

func mustParseInt(s string) uint64 {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert string %s to integer: %v", s, err)
	}
	return uint64(i)
}
