package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"
)

var size = 1000

func main() {
	m := readFileToMatrix("data", nil)

	r, c := m.Dims()
	good := 0
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			v := m.At(i, j)
			if v >= 2 {
				//fmt.Println(i, j, v)
				good++
			}
		}
	}

	//fc := mat.Formatted(m, mat.Prefix("    "), mat.Squeeze())
	//fmt.Printf("m = %v", fc)
	fmt.Println(good)
	_ = readFileToMatrix("data", m)
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
	c := mat.NewDense(size, size, nil)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())
		offset := strings.Split(strings.Replace(ff[2], ":", "", 1), ",")
		oy, err := strconv.ParseInt(offset[0], 10, 64)
		if err != nil {
			panic(err)
		}
		ox, err := strconv.ParseInt(offset[1], 10, 64)
		if err != nil {
			panic(err)
		}

		dim := strings.Split(ff[3], "x")
		dy, err := strconv.ParseInt(dim[0], 10, 64)
		if err != nil {
			panic(err)
		}
		dx, err := strconv.ParseInt(dim[1], 10, 64)
		if err != nil {
			panic(err)
		}
		temp := mat.NewDense(size, size, nil)
		isIT := true
		for i := ox; i < ox+dx; i++ {
			for j := oy; j < oy+dy; j++ {

				if second != nil {
					v := second.At(int(i), int(j))
					if v != 1 {
						isIT = false
					}
				}
				temp.Set(int(i), int(j), 1)
			}
		}
		if isIT && second != nil {
			fmt.Println(ff)
		}
		var foo mat.Dense
		foo.Add(c, temp)

		c = &foo
		//fmt.Println(offset, dim)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return c
}
