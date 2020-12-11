package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

//go:generate go run github.com/dmarkham/enumer -type=Kind --linecomment
type Kind int

// Kind* constants refer to tile kinds for input and output.
const (
	KindFloor    Kind = iota // .
	KindOccupied             // #
	KindEmpty                // L
)

type Tile struct {
	X    int
	Y    int
	Kind Kind
	W    *World
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
	return fmt.Sprintf("Tile: X:%v Y:%v Kind:%v\n", t.X, t.Y, t.Kind.String())
}

func main() {
	flag.Parse()
	lines := readFileToLines("data.txt")
	w := World{Map: make(map[int]map[int]*Tile)}

	for x, l := range lines {
		//fmt.Println("Line:", l)
		for y, c := range l {

			//parts := strings.Fields(l)
			kind, err := KindString(string(c))
			if err != nil {
				kind = KindFloor
			}
			w.SetTile(&Tile{
				Kind: kind,
			}, x, y)

		}
	}
	for {
		//fmt.Println("Round:", w)
		w = tick(w)
		if w.Done {
			break
		}
	}
	//fmt.Println("Part1:", w)
	fmt.Println("Part1", len(w.FindAllKind(KindOccupied)))
	w = World{Map: make(map[int]map[int]*Tile)}

	for x, l := range lines {
		for y, c := range l {

			kind, err := KindString(string(c))
			if err != nil {
				kind = KindFloor
			}
			w.SetTile(&Tile{
				Kind: kind,
			}, x, y)

		}
	}

	for {
		w = tick2(w)
		if w.Done {
			break
		}
	}
	fmt.Println("Part2", len(w.FindAllKind(KindOccupied)))

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

func tick2(world World) World {

	flips := make([]*Tile, 0)
	for _, seat := range append(world.FindAllKind(KindEmpty), world.FindAllKind(KindOccupied)...) {
		//fmt.Println(seat)
		//fmt.Println(seat.LongRange())
		count := 0
		for _, t := range seat.LongRange() {
			if seat.Kind == KindEmpty && t.Kind == KindOccupied {
				count++
			} else if seat.Kind == KindOccupied && t.Kind == KindOccupied {
				count++
			}
		}
		//fmt.Println(seat, count)
		if seat.Kind == KindEmpty && count == 0 {
			flips = append(flips, seat)
		} else if seat.Kind == KindOccupied && count >= 5 {
			flips = append(flips, seat)
		}
	}

	for _, s := range flips {
		//fmt.Println("FLipping", s)
		if s.Kind == KindEmpty {
			s.Kind = KindOccupied
		} else if s.Kind == KindOccupied {
			s.Kind = KindEmpty

		}
		//fmt.Println("Flipp[ed", s)

	}

	world.Round++
	if len(flips) == 0 {
		world.Done = true
	}
	return world

}

func tick(world World) World {

	flips := make([]*Tile, 0)
	for _, seat := range append(world.FindAllKind(KindEmpty), world.FindAllKind(KindOccupied)...) {
		//fmt.Println(seat)
		//fmt.Println(seat.Range())
		count := 0
		for _, t := range seat.Range() {
			if seat.Kind == KindEmpty && t.Kind == KindOccupied {
				count++
			} else if seat.Kind == KindOccupied && t.Kind == KindOccupied {
				count++
			}
		}
		//fmt.Println(seat, count)
		if seat.Kind == KindEmpty && count == 0 {
			flips = append(flips, seat)
		} else if seat.Kind == KindOccupied && count >= 4 {
			flips = append(flips, seat)
		}
	}

	for _, s := range flips {
		//fmt.Println("FLipping", s)
		if s.Kind == KindEmpty {
			s.Kind = KindOccupied
		} else if s.Kind == KindOccupied {
			s.Kind = KindEmpty

		}
		//fmt.Println("Flipp[ed", s)

	}

	world.Round++
	if len(flips) == 0 {
		world.Done = true
	}
	return world

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
				if n.Kind == KindFloor {
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
