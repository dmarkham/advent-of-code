package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/floats"

	"gonum.org/v1/gonum/mat"
)

// https://adventofcode.com/2018/day/3

var size = 1000 // Just guessed an arbitrary size

func main() {

	m := readFileToMatrix("data", nil)
	overlap := floats.Count(func(f float64) bool { return f >= 2 }, m.RawMatrix().Data) // anything in the matrix over 2
	fmt.Println(overlap)                                                                // answer to part 1
	// Then for part 2 We just feed thr matrix in again looking for an entire
	// read in that was all 1's in the matrix
	// it will print out the matrix ID when it finds it
	readFileToMatrix("data", m)
}

// Pull all lines into a int slice
func readFileToMatrix(file string, second *mat.Dense) *mat.Dense {
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	r := bufio.NewReader(fh)
	scanner := bufio.NewScanner(r)
	c := mat.NewDense(size, size, nil) // Will hold everything at the end and get returned
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		// Parse the offest first
		offset := strings.Split(strings.Replace(ff[2], ":", "", 1), ",")
		oy, err := strconv.ParseInt(offset[0], 10, 64)
		if err != nil {
			panic(err)
		}
		ox, err := strconv.ParseInt(offset[1], 10, 64)
		if err != nil {
			panic(err)
		}
		// Now the deminsions
		dim := strings.Split(ff[3], "x")
		dy, err := strconv.ParseInt(dim[0], 10, 64)
		if err != nil {
			panic(err)
		}
		dx, err := strconv.ParseInt(dim[1], 10, 64)
		if err != nil {
			panic(err)
		}

		isIT := true // used for part 2
		for i := ox; i < ox+dx; i++ {
			for j := oy; j < oy+dy; j++ {
				if second != nil {
					v := second.At(int(i), int(j))
					if v != 1 { // it can not be the right id if it was conflicted
						isIT = false
					}
				}
				c.Set(int(i), int(j), 1+c.At(int(i), int(j))) // add them up
			}
		}
		if isIT && second != nil { // Found the part2 answer
			fmt.Println(ff)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return c
}
