package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

// https://adventofcode.com/2018/day/4
var part2 bool

type Shift struct {
	Guard    int
	Work     bool
	Duration time.Duration
	Start    time.Time
	End      time.Time
}

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

func main() {
	flag.Parse()

	shifts := readFileToShifts("data")
	mostAsleep := make(map[int]float64)
	for _, s := range shifts {
		if !s.Work {
			mostAsleep[s.Guard] += s.Duration.Minutes()
		}
	}

	gAsleepMost := 0
	gmin := float64(0)

	for k, v := range mostAsleep {
		if v >= gmin {
			gAsleepMost = k
			gmin = v
		}
	}
	if !part2 {
		fmt.Printf("Guard: %v, Sleep: %v\n", gAsleepMost, gmin)
	}
	hm := make(map[string]int)
	for _, s := range shifts {
		if !s.Work && (part2 || (gAsleepMost == s.Guard)) {
			start := s.Start // walk the mins
			for i := 0; i < int(s.Duration.Minutes()); i++ {
				hm[fmt.Sprintf("%v-%v", s.Guard, start.Format("15:04"))]++
				start = start.Add(1 + time.Minute)
			}
		}
	}
	var ss []kv
	for k, v := range hm {
		ss = append(ss, kv{k, v})
	}

	// Then sorting the slice by value, higher first.
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})
	// Print the x top values
	for _, kv := range ss[:1] {
		fmt.Printf("Gaurd: %v    MinCount: %v\n", kv.Key, kv.Value)
	}
}

type kv struct {
	Key   string
	Value int
}

// Pull all lines into a int slice
func readFileToShifts(file string) []*Shift {
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	r := bufio.NewReader(fh)
	scanner := bufio.NewScanner(r)
	shifts := make([]*Shift, 0)
	shift := &Shift{}
	for scanner.Scan() {

		ff := strings.Fields(scanner.Text())
		time, err := time.Parse("[2006-01-02 15:04]", ff[0]+" "+ff[1])
		if err != nil {
			panic(err)
		}
		switch ff[2] {
		case "Guard":

			shift = &Shift{}
			shift.Start = time
			ff[3] = strings.Replace(ff[3], "#", "", 1)
			gID, err := strconv.ParseInt(ff[3], 10, 64)
			if err != nil {
				panic(err)
			}
			shift.Guard = int(gID)
			shift.Work = true
		case "wakes":
			if shift.Guard != 0 {
				shift.End = time
				shift.Duration = shift.End.Sub(shift.Start)
				shifts = append(shifts, shift)
			}
			gID := shift.Guard
			shift = &Shift{}
			shift.Start = time
			shift.Guard = int(gID)
			shift.Work = true

		case "falls":
			if shift.Guard != 0 {
				shift.End = time
				shift.Duration = shift.End.Sub(shift.Start)
				shifts = append(shifts, shift)
			}
			gID := shift.Guard
			shift = &Shift{}
			shift.Start = time
			shift.Guard = int(gID)
			shift.Work = false

		default:
			panic("Messed up" + ff[2])
		}

	}
	return shifts
}
