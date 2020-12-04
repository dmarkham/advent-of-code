package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var part2 bool

func init() {
	flag.BoolVar(&part2, "part2", false, "Run Part2?")
}

var heightsRegex = regexp.MustCompile(`^(\d+)(cm|in)$`)
var hcolorRegex = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var ecolorRegex = regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
var pidRegex = regexp.MustCompile(`^[0-9]{9}$`)

type Passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func (p *Passport) Valid() bool {
	if p.byr == "" {
		return false
	}
	if mustParseInt(p.byr) > 2002 || mustParseInt(p.byr) < 1920 {
		return false
	}

	if p.iyr == "" {
		return false
	}
	if mustParseInt(p.iyr) > 2020 || mustParseInt(p.iyr) < 2010 {
		return false
	}

	if p.eyr == "" {
		return false
	}
	if mustParseInt(p.eyr) > 2030 || mustParseInt(p.eyr) < 2020 {
		return false
	}

	if p.hgt == "" {
		return false
	}
	if !heightsRegex.MatchString(p.hgt) {
		return false
	}
	t := heightsRegex.FindStringSubmatch(p.hgt)
	h := mustParseInt(t[1])
	if t[2] == "cm" {
		if h < 150 || h > 193 {
			return false
		}
	}
	if t[2] == "in" {
		if h < 59 || h > 76 {
			return false
		}
	}
	if p.hcl == "" {
		return false
	}

	if !hcolorRegex.MatchString(p.hcl) {
		return false
	}

	if p.ecl == "" {
		return false
	}

	if !ecolorRegex.MatchString(p.ecl) {
		return false
	}

	if p.pid == "" {
		return false
	}

	if !pidRegex.MatchString(p.pid) {
		return false
	}

	return true
}

func main() {
	flag.Parse()
	passports := readFileToLines("data.txt")
	fmt.Println(len(passports))
	valid := 0
	for _, p := range passports {
		//fmt.Println(p.hcl)
		if p.Valid() {
			valid++
		}
	}
	fmt.Println(valid)
}

// Pull all lines into a string slice
func readFileToLines(file string) []*Passport {
	// open data
	fh, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer fh.Close()
	r := bufio.NewReader(fh)
	scanner := bufio.NewScanner(r)
	// read it all in

	passports := make([]*Passport, 0)
	passport := &Passport{}
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		//fmt.Println(line)
		parts := strings.Fields(line)
		if line == "" {
			passports = append(passports, passport)
			passport = &Passport{}
		}
		for _, p := range parts {

			kv := strings.Split(p, ":")
			switch kv[0] {
			case "byr":
				passport.byr = kv[1]
			case "iyr":
				passport.iyr = kv[1]
			case "eyr":
				passport.eyr = kv[1]
			case "hgt":
				passport.hgt = kv[1]
			case "hcl":
				passport.hcl = kv[1]
			case "ecl":
				passport.ecl = kv[1]
			case "pid":
				passport.pid = kv[1]
			case "cid":
				passport.cid = kv[1]
			default:
				panic(p)
			}

		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	passports = append(passports, passport)
	return passports
}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert string %s to integer: %v", s, err)
	}
	return i
}
