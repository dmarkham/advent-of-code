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
	generationsToDo := 50000000000
	lastGeneration := ""
	//patterns = make(map[string]string)
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
			//parts := strings.Split(line, " => ")
			//patterns[parts[0]] = parts[1]
			continue

		}
	}

	gen := 0
	//t := time.Now()
	for {
		lastGeneration = tick(lastGeneration)
		lastGeneration, lineStart = padGeneration(lastGeneration, lineStart)
		//generations = append(generations, newGen)
		//lastGeneration = newGen
		if gen == generationsToDo-1 {
			break
		}
		if gen == 2010 {
			diff := generationsToDo - 2 - gen
			gen = gen + diff
			lineStart = lineStart + diff
		}
		if gen == 19 {

			fmt.Print("Gen:", gen, " ")
			printTotal(lastGeneration, lineStart)
		}
		gen++

	}
	fmt.Print("Gen:", gen, " ")
	printTotal(lastGeneration, lineStart)

}

func printTotal(s string, lineStart int) {

	totalPots := 0
	aORb := regexp.MustCompile("#")
	matches := aORb.FindAllStringIndex(s, -1)
	for _, match := range matches {
		//fmt.Println(s, match[1], (lineStart + match[1] - 1))
		totalPots = totalPots + lineStart + match[1] - 1
	}
	fmt.Print(lineStart, s)
	fmt.Println("totalPots:", totalPots)
	return
}

func tick(s string) string {
	n := len(s)
	var b byte
	buf := make([]byte, n)
	copy(buf, s)

	for i := 0; i < n-5; i++ {

		switch s[i : i+5] {
		case ".#...":
			b = '#'
		case "...##", "#.#.#", ".###.", ".#..#", "#..#.", "###..", ".####", "#..##", "##.##", ".#.##", "##...", "##..#", "#.##.":
			b = '#'
		default:
			b = '.'

		}
		buf[i+2] = b
	}
	return string(buf)
}

func padGeneration(s string, lineStart int) (string, int) {

	index := strings.Index(s, "#")

	if index > 3 {
		s = s[index-4:]
		lineStart = lineStart + index - 4
	}
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
