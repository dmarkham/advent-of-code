package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"

	"gonum.org/v1/gonum/mat"

	"gonum.org/v1/gonum/floats"
)

// https://adventofcode.com/2018/day/6
var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

type Cord struct {
	RowNum int
	X      int
	Y      int
	Score  float64
}

func main() {
	flag.Parse()
	lines := readFileToLines("data")
	rows := make([]float64, 0)
	cols := make([]float64, 0)
	cords := make([]*Cord, 0)
	lineNumber := 0
	for _, l := range lines {
		ff := strings.Split(l, ", ")
		col, err := strconv.ParseInt(ff[0], 10, 64)
		if err != nil {
			panic(err)
		}
		row, err := strconv.ParseInt(ff[1], 10, 64)
		if err != nil {
			panic(err)
		}
		rows = append(rows, float64(row))
		cols = append(cols, float64(col))
		c := &Cord{
			X:      int(row),
			Y:      int(col),
			RowNum: lineNumber,
		}
		cords = append(cords, c)
		lineNumber++
	}
	maxRow := floats.Max(rows) + 1
	maxCol := floats.Max(cols) + 2
	matrix := mat.NewDense(int(maxRow), int(maxCol), nil)
	dimRows, dimCol := matrix.Dims()
	part2Count := 0
	for i := 0; i < dimRows; i++ {
		for j := 0; j < dimCol; j++ {
			minC := intergerDistance(i, j, cords)
			if minC == nil {
				matrix.Set(i, j, float64(99))
			} else if minC.Score == 0 {
				matrix.Set(i, j, float64(-1*minC.RowNum))
			} else {
				matrix.Set(i, j, float64(minC.RowNum))
			}
			if part2 {
				tScore := intergerTotalDistance(i, j, cords)
				if tScore < 10000 {
					part2Count++
				}
			}
		}
	}
	if part2 {
		fmt.Println("Part2 Count:", part2Count)
		os.Exit(0)
	}
	seen := make(map[int]bool)
	for i := 0; i < dimRows; i++ {
		for j := 0; j < dimCol; j++ {
			value := matrix.At(i, j)
			seen[int(value)] = true
			if i != 0 && i != dimRows-1 && j != dimCol-1 {
				j = dimCol - 1
			}
		}
	}
	counts := make(map[int]int)
	for _, v := range cords {
		if seen[v.RowNum] {
			continue
		}
		count := floats.Count(func(f float64) bool { return f == float64(v.RowNum) }, matrix.RawMatrix().Data)
		counts[v.RowNum] = count + 1
	}
	var ss []kv
	for k, v := range counts {
		ss = append(ss, kv{fmt.Sprintf("%v", k), v})
	}

	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	for _, kv := range ss[:1] { // Print the x top values
		fmt.Printf("Item %v    MinCount: %v\n", kv.Key, kv.Value)
	}
}

func intergerTotalDistance(x, y int, c []*Cord) float64 {
	totalScore := float64(0)
	for i := 0; i < len(c); i++ {
		s := math.Abs(float64(x - c[i].X))
		s += math.Abs(float64(y - c[i].Y))
		totalScore += s
	}
	return totalScore
}
func intergerDistance(x, y int, c []*Cord) *Cord {
	min := math.MaxFloat64
	var minCord *Cord
	seen := make(map[float64]int)
	totalScore := float64(0)
	for i := 0; i < len(c); i++ {
		s := math.Abs(float64(x - c[i].X))
		s += math.Abs(float64(y - c[i].Y))
		totalScore += min
		if s <= min {
			//fmt.Printf("I:%v, X:%v, Y:%v, X1:%v Y1:%v  Score:%v\n", i, x, y, c[i].X, c[i].Y, s)
			minCord = c[i]
			min = s
			seen[min]++
		}
	}
	if seen[min] > 1 {
		return nil
	}
	minCord.Score = min
	return minCord
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

type kv struct {
	Key   string
	Value int
}
