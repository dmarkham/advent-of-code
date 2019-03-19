package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime/pprof"
	"strings"
	"time"
)

var patterns map[string]string
var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to `file`")

func main() {

	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	lines := readFileToLines("data")
	generationsToDo := 20
	lastGeneration := ""
	patterns = make(map[string]string)
	lineStart := 0
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if strings.HasPrefix(line, "initial state: ") {
			line = strings.TrimPrefix(line, "initial state: ")
			line, lineStart = padGeneration(line, lineStart)

			lastGeneration = line
			continue
		}
		if strings.Contains(line, "=>") {
			parts := strings.Split(line, " => ")
			patterns[parts[0]] = parts[1]
			continue

		}
	}

	gen := 0
	t := time.Now()
	for {
		newGen := tick(lastGeneration)
		newGen, lineStart = padGeneration(newGen, lineStart)
		//generations = append(generations, newGen)
		lastGeneration = newGen
		if gen == generationsToDo-1 {
			break
		}
		if gen%100 == 0 {
			taken := time.Since(t).Truncate(time.Second)
			donePercent := float64(gen) / float64(generationsToDo)
			fmt.Println("Generation:", gen, taken, donePercent)
		}
		gen++

	}

	totalPots := 0
	aORb := regexp.MustCompile("#")
	matches := aORb.FindAllStringIndex(lastGeneration, -1)
	for _, match := range matches {
		fmt.Println(lineStart, match[1], (lineStart + match[1] - 1))
		totalPots = totalPots + lineStart + match[1] - 1
	}
	fmt.Println("totalPots:", totalPots)
}

func tick(s string) string {
	newGen := s
	for i := 0; i < len(s)-5; i++ {

		//foo := patterns[s[i:i+5]]
		//if foo == "" {
		//	foo = "."
		//}
		//fmt.Println(s)
		//fmt.Println(i, foo)
		//fmt.Println(newGen)
		//fmt.Println(newGen[0:i+2], foo, newGen[i+3:])
		newGen = newGen[0:i+2] + patterns[s[i:i+5]] + newGen[i+3:]
		//fmt.Println(newGen)
		//fmt.Println("")
	}
	//fmt.Println("GEN DONE")
	return newGen
}

func padGeneration(s string, lineStart int) (string, int) {

	if s[0] == '#' {
		s = "...." + s
		lineStart = lineStart - 4
	} else if s[1] == '#' {
		s = "..." + s
		lineStart = lineStart - 3
	} else if s[2] == '#' {
		s = ".." + s
		lineStart = lineStart - 2
	} else if s[3] == '#' {
		s = "." + s
		lineStart = lineStart - 1
	}

	if s[len(s)-1] == '#' {
		s = s + "...."
	} else if s[len(s)-2] == '#' {
		s = s + "..."
	} else if s[len(s)-3] == '#' {
		s = s + ".."
	} else if s[len(s)-4] == '#' {
		s = s + "."
	}
	return s, lineStart
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
