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

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

type Jolts struct {
	Min int
	Max int
	Set []int
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	jolts := make([]int, 0)
	jolts2 := make([]int, 0)
	for _, l := range lines {
		n := mustParseInt(l)
		jolts = append(jolts, n)
		jolts2 = append(jolts2, n)
		//fmt.Println("Line:", l)
		//parts := strings.Fields(l)

		//r := strings.Split(parts[0], "-")

		//fmt.Println("range:", parts[0])
		//fmt.Println("char:", string(parts[1][0]))
		//fmt.Println("Pass:", parts[2])
		//rule := &PassRule{ }
		//fmt.Println(rule, rule.Valid())
	}

	jolts = append(jolts, 0)
	sort.Sort(sort.IntSlice(jolts))
	jolts = append(jolts, jolts[len(jolts)-1]+3)
	max := jolts[len(jolts)-1]
	fmt.Println(jolts)
	threes := 0
	ones := 0

	for i := 1; i < len(jolts); i++ {
		diff := jolts[i] - jolts[i-1]
		//fmt.Println("DIFF:", diff)
		if diff == 1 {
			ones++
		} else if diff == 3 {
			threes++
		}
	}

	fmt.Println(ones, threes, ones*threes)

	sort.Sort(sort.IntSlice(jolts2))
	fmt.Println(jolts2)
	jo := &Jolts{
		Set: jolts2,
		Min: 0,
		Max: max,
	}

	count2 := subset(jo)

	fmt.Println(count2)
}
func remove(s []int, i int) []int {
	return append(s[:i], s[i+1:]...)
}

func subset(jo *Jolts) int {
	fmt.Println("SubSet:", jo.Min, jo.Set)
	if len(jo.Set) == 0 {
		return 0
	}
	count := 0
	start := 0
	end := -1
	for {
		jolts3 := make([]int, len(jo.Set))
		copy(jolts3, jo.Set)
		fmt.Println("Removing:", start, end, jolts3[start:end+1])
		if end >= len(jolts3)-1 {
			break
		}
		if start >= len(jolts3) {
			break
		}
		testSet := append(jolts3[0:start], jolts3[end+1:]...)
		testSet = append([]int{jo.Min}, testSet...)
		testSet = append(testSet, jo.Max)
		if diffGood(testSet) {
			if len(jolts3[start:end+1]) > 0 {
				count += subset(&Jolts{
					Max: jo.Max,
					Min: jo.Set[end+1],
					Set: jo.Set[end+2:],
				})
			} else {
				count++

			}

			end++
		} else {
			start++
			end = start
		}
	}
	fmt.Println("Returned Count: ", count)
	return count
}
func diffGood(jolts []int) bool {

	for i := 1; i < len(jolts); i++ {
		diff := jolts[i] - jolts[i-1]
		//fmt.Println("DIFF:", diff)
		if diff < 4 {
		} else {
			fmt.Println("MISS:", jolts, diff, jolts[i], jolts[i-1])
			return false
		}
	}
	fmt.Println("HIT:", jolts)
	return true

}

func adapt(jolts []int, min, max int) bool {
	jolts = append(jolts, min)
	sort.Sort(sort.IntSlice(jolts))
	jolts = append(jolts, max)
	for i := 1; i < len(jolts); i++ {
		diff := jolts[i] - jolts[i-1]
		//fmt.Println("DIFF:", diff)
		if diff < 4 {
		} else {
			fmt.Println("MISS:", jolts, diff, jolts[i], jolts[i-1])
			return false
		}
	}
	fmt.Println("HIT:", jolts)
	return true
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

func subsetSum(numbers []int, target int, partial []int) []int {
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
func sum(list []int) int {
	sum := 0
	for _, part := range list {
		sum = sum + part
	}
	return sum
}
