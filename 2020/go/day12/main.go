package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

//go:generate go run github.com/dmarkham/enumer -type=Kind --linecomment
type Kind int

// Kind* constants refer to tile kinds for input and output.
const (
	KindF    Kind = iota // F
	KindN                // N
	KindR                // R
	KindL                // L
	KindS                // S
	KindE                // E
	KindW                // W
	KindShip             // Ship
	KindWP               // WP
)

type Tile struct {
	X         int
	Y         int
	Kind      Kind
	W         *World
	Direction int
}

// World is a two dimensional map of Tiles.
type World struct {
	Map   map[int]map[int]*Tile
	Done  bool
	Round int
	Score int
}

// Tile gets the tile at the given coordinates in the world.
func (w *World) Tile(x, y int) *Tile {
	if w.Map[x] == nil {
		return nil
	}
	return w.Map[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w *World) SetTile(t *Tile, x, y int) {

	if w.Map[x] == nil {
		w.Map[x] = map[int]*Tile{}
	}
	t.X = x
	t.Y = y
	t.W = w
	w.Map[x][y] = t
}

func (t *Tile) String() string {
	return fmt.Sprintf("Tile: X:%v Y:%v Kind:%v, Direction %v \n", t.X, t.Y, t.Kind.String(), t.Direction)
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	w := World{Map: make(map[int]map[int]*Tile)}
	kind, _ := KindString("Ship")
	ship := &Tile{
		Kind:      kind,
		Direction: 90,
		X:         0,
		Y:         0,
	}
	w.SetTile(ship, 0, 0)
	// Read in the directions
	for _, l := range lines {
		//fmt.Println("Line:", l)
		k := string(l[0])
		amount := mustParseInt(l[1:])
		kind, err := KindString(k)
		if err != nil {
			panic(k)
		}

		//fmt.Println(amount, kind)
		switch kind {
		case KindN:
			for i := 0; i < amount; i++ {
				ship.Up()
			}
		case KindS:
			for i := 0; i < amount; i++ {
				ship.Down()
			}
		case KindE:
			for i := 0; i < amount; i++ {
				ship.Right()
			}
		case KindW:
			for i := 0; i < amount; i++ {
				ship.Left()
			}
		case KindF:
			switch math.Abs(float64(ship.Direction)) {
			case 0:
				for i := 0; i < amount; i++ {
					ship.Up()
				}
			case 90:
				for i := 0; i < amount; i++ {
					ship.Right()
				}
			case 180:
				for i := 0; i < amount; i++ {
					ship.Down()
				}
			case 270:
				for i := 0; i < amount; i++ {
					ship.Left()
				}
			}

		case KindR:
			ship.Direction = (ship.Direction + amount) % 360
		case KindL:
			ship.Direction = (ship.Direction - amount) % 360
			if ship.Direction < 0 {
				ship.Direction = ship.Direction + 360
			}
		}

		//fmt.Println(ship)

	}
	//fmt.Println(ship)
	part1 := math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y))
	fmt.Println("PArt1", part1)

	// part2
	w = World{Map: make(map[int]map[int]*Tile)}

	ship = &Tile{
		Kind:      kind,
		Direction: 90,
		X:         0,
		Y:         0,
	}
	w.SetTile(ship, 0, 0)
	kind, _ = KindString("WP")
	wp := &Tile{
		Kind:      kind,
		Direction: 0,
		X:         10,
		Y:         1,
	}
	w.SetTile(wp, 0, 0)
	// Read in the directions
	for _, l := range lines {
		//fmt.Println("Line:", l)
		k := string(l[0])
		amount := mustParseInt(l[1:])
		kind, err := KindString(k)
		if err != nil {
			panic(k)
		}

		//fmt.Println(amount, kind)
		switch kind {
		case KindN:
			for i := 0; i < amount; i++ {
				wp.Up()
			}
		case KindS:
			for i := 0; i < amount; i++ {
				wp.Down()
			}
		case KindE:
			for i := 0; i < amount; i++ {
				wp.Right()
			}
		case KindW:
			for i := 0; i < amount; i++ {
				wp.Left()
			}
		case KindF:
			switch math.Abs(float64(ship.Direction)) {
			case 0:
				for i := 0; i < amount; i++ {
					ship.Up()
				}
			case 90:
				for i := 0; i < amount; i++ {
					ship.Right()
				}
			case 180:
				for i := 0; i < amount; i++ {
					ship.Down()
				}
			case 270:
				for i := 0; i < amount; i++ {
					ship.Left()
				}
			}

		case KindR:

			Xdiff := ship.X - wp.X
			Ydiff := ship.Y - wp.Y
			if Xdiff > 0 && Ydiff > 0 {
				wp.X += Xdiff
				wp.Y += Ydiff

			}

			wp.X += Xdiff
			wp.Y += Ydiff

			//ship.Direction = (ship.Direction + amount) % 360
		case KindL:
			//ship.Direction = (ship.Direction - amount) % 360
			//if ship.Direction < 0 {
			//	ship.Direction = ship.Direction + 360
			//	}
		}

		//fmt.Println(ship)

	}
	//fmt.Println(ship)
	part1 := math.Abs(float64(ship.X)) + math.Abs(float64(ship.Y))
	fmt.Println("PArt1", part1)

}

func mustParseInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("cannot convert string %s to integer: %v", s, err)
	}
	return i
}
func (w World) FindAllKind(k Kind) []*Tile {
	f := make([]*Tile, 0)
	xlen := len(w.Map)
	ylen := len(w.Map[0])
	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if w.Map[x][y].Kind == k {
				f = append(f, w.Map[x][y])
			}
		}
	}
	return f
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

func (t *Tile) Range() []*Tile {
	neighbors := make([]*Tile, 0)
	for _, offset := range [][]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

func (t *Tile) LongRange() []*Tile {
	neighbors := make([]*Tile, 0)
	for _, offset := range [][]int{
		{0, -1},
		{0, 1},
		{1, 0},
		{-1, 0},
		{-1, -1},
		{-1, 1},
		{1, 1},
		{1, -1},
	} {
		x := offset[0]
		y := offset[1]
		for {
			if n := t.W.Tile(t.X+x, t.Y+y); n != nil {
				//fmt.Println(n)
				if n.Kind == KindN {
					x += offset[0]
					y += offset[1]
					continue
				}
				neighbors = append(neighbors, n)
				break
			} else if n == nil {
				break
			}
		}
	}
	return neighbors
}

func (t *Tile) Up() {
	t.Y += 1

}
func (t *Tile) Down() {
	t.Y -= 1
}
func (t *Tile) Left() {
	t.X -= 1
}
func (t *Tile) Right() {

	t.X += 1
}
