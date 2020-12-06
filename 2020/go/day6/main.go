package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	lines := readFileToLines("data.txt")
	total1 := 0
	total2 := 0
	for _, l := range lines {
		seen := make(map[string]int)

		parts := strings.Fields(l)

		for _, p := range parts {
			for _, c := range p {
				seen[string(c)]++
			}
		}

		total1 = total1 + len(seen)
		groupSize := len(parts)

		for _, v := range seen {
			if v == groupSize {
				total2++
			}
		}
	}
	fmt.Println("Part1 ", total1)
	fmt.Println("Part2:", total2)
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
	scanner.Split(Paragraph)
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

func Paragraph(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of the separator substring
	if i := strings.Index(string(data), "\n\n"); i >= 0 {
		return i + len("\n\n"), bytes.ReplaceAll(data[0:i], []byte{'\n'}, []byte{' '}), nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), bytes.ReplaceAll(data, []byte{'\n'}, []byte{' '}), nil
	}

	return
}
