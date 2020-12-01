package main

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"io/ioutil"
	"strings"
)

//go:generate go run github.com/dmarkham/enumer -type=Kind --linecomment
type Kind int

// Kind* constants refer to tile kinds for input and output.
const (
	// KindPlain (.) is a plain tile with a movement cost of 1.
	KindPlain Kind = iota // .
	// KindRiver (~) is a river tile with a movement cost of 2.
	KindRiver // ~
	// KindMountain (M) is a mountain tile with a movement cost of 3.
	KindMountain // M
	// KindBlocker (#) is a tile which blocks movement.
	KindBlocker // #
	// KindFrom (F) is a tile which marks where the path should be calculated // from.
	KindFrom // F
	// KindTo (T) is a tile which marks the goal of the path.
	KindTo // T
	// KindPath (+) is a tile to represent where the path is in the output.
	KindPath  // +
	KindElf   // E
	KindGnome // G
)

type Tile struct {
	X           int
	Y           int
	Kind        Kind
	AttackPower int
	HP          int
	W           *World
}

func (t *Tile) String() string {
	return fmt.Sprintf("Tile: X:%v Y:%v Kind:%v HP:%v\n", t.X, t.Y, t.Kind.String(), t.HP)
}

// World is a two dimensional map of Tiles.
type World struct {
	Map   map[int]map[int]*Tile
	Done  bool
	Round int
	HP    int
	Score int
}

func tick(world World) World {
	players := world.Fighters()

	for _, fighter := range players {
		if fighter.IsDead() {
			fighter.Kind = KindPlain
		}
	}

	for i, fighter := range players {

		var vs []*Tile
		switch fighter.Kind {
		case KindElf:
			vs = world.FindAllKind(KindGnome)
		case KindGnome:
			vs = world.FindAllKind(KindElf)
		default:
			continue
		}
		if len(vs) == 0 {
			if i != len(players)-1 {
				fmt.Println("I", i, "lenplayer:", len(players))
				//world.Round--
			}
			world.Done = true
			players := world.Fighters()
			hpLeft := 0
			for _, p := range players {
				hpLeft += p.HP
			}
			world.HP = hpLeft
			world.Score = hpLeft * world.Round

			return world
		}

		if !fighter.CanAttack() {
			fighter.W = &world
			p := bestPath(fighter, vs)
			fighter.Move(&world, p)
		}
		if fighter.CanAttack() {
			fighter.W = &world
			//fmt.Println("Can Attack", fighter)
			fighter.Attack()
			for _, p := range players {
				if p.IsDead() {
					p.Kind = KindPlain
				}
			}
		}
	}
	world.Round++
	return world

}

func main() {
	data, err := ioutil.ReadFile("data")
	if err != nil {
		panic(err)
	}

	world := ParseWorld(string(data))
	//fmt.Println(world)
	round := 0
	for world.Done == false {
		//time.Sleep(time.Millisecond * 300)
		var input string
		fmt.Scanln(&input)
		fmt.Printf("After Round: %v\n%s\n", round, world.RenderPath(nil))
		/*
			players := world.Fighters()
			for _, p := range players {
				fmt.Println(p)
			}
		*/
		world = tick(world)
		round++
	}

	fmt.Printf("Done Score:%v on Round: %v \n%s\n", world.Score, world.Round, world.RenderPath(nil))
	players := world.Fighters()
	for _, p := range players {
		fmt.Println(p)
	}
}

func bestPath(from *Tile, to []*Tile) []astar.Pather {
	var best []astar.Pather
	bestVal := float64(1000000.0)
	for _, start := range from.Range() {
		start.W = from.W
		for _, p := range to {
			p.W = from.W
			for _, inRange := range p.Range() {

				p, dist, found := astar.Path(start, inRange)
				if !found {
					//fmt.Println("Could not find a path\n", to, from)
					//fmt.Printf("Resulting path\n%s\n", from.W.RenderPath(nil))
				} else {
					if dist < bestVal {
						best = p
						bestVal = dist
					}
					//fmt.Printf("Resulting path\n%s\n", from.W.RenderPath(p))
				}
			}
		}
	}
	//fmt.Printf("Resulting path\n%s\n", from.W.RenderPath(best))
	return best
}

// ParseWorld parses a textual representation of a world into a world map.
func ParseWorld(input string) World {
	w := World{Map: make(map[int]map[int]*Tile)}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			kind, err := KindString(string(raw))
			if err != nil {
				kind = KindBlocker
			}
			w.SetTile(&Tile{
				Kind:        kind,
				HP:          200,
				AttackPower: 3,
			}, x, y)
		}
	}
	return w
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

// RenderPath renders a path on top of a world.
func (w World) RenderPath(path []astar.Pather) string {
	width := len(w.Map)
	if width == 0 {
		return ""
	}
	height := len(w.Map[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*Tile)
		pathLocs[fmt.Sprintf("%d,%d", pT.X, pT.Y)] = true
	}
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.Tile(x, y)
			r := ' '
			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
				r = rune(KindPath.String()[0])
			} else if t != nil {
				r = rune(t.Kind.String()[0])
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

func (t *Tile) IsDead() bool {
	return t.HP < 1
}
func (t *Tile) PathNeighbors() []astar.Pather {
	neighbors := []astar.Pather{}
	if d := t.Up(); d != nil &&
		d.Kind != KindBlocker &&
		d.Kind != KindGnome &&
		d.Kind != KindElf {
		neighbors = append(neighbors, d)
	}
	if d := t.Left(); d != nil &&
		d.Kind != KindBlocker &&
		d.Kind != KindGnome &&
		d.Kind != KindElf {
		neighbors = append(neighbors, d)
	}
	if d := t.Right(); d != nil &&
		d.Kind != KindBlocker &&
		d.Kind != KindGnome &&
		d.Kind != KindElf {
		neighbors = append(neighbors, d)
	}
	if d := t.Down(); d != nil &&
		d.Kind != KindBlocker &&
		d.Kind != KindGnome &&
		d.Kind != KindElf {
		neighbors = append(neighbors, d)
	}

	return neighbors
}

func (t *Tile) Move(w *World, p []astar.Pather) {

	if len(p) < 1 {
		return
	}
	//fmt.Println(p)
	t2, ok := p[len(p)-1].(*Tile)
	if !ok {
		panic("Error")
	}
	w.SetTile(&Tile{
		Kind: KindPlain}, t.X, t.Y)

	w.SetTile(t, t2.X, t2.Y)
}

func (t *Tile) CanAttack() bool {
	if t.HP < 1 {
		return false
	}

	if d := t.Up(); d.HP > 0 {
		switch d.Kind {
		case KindElf, KindGnome:
			if d.Kind != t.Kind {
				return true
			}
		}
	}
	if d := t.Left(); d.HP > 0 {
		switch d.Kind {
		case KindElf, KindGnome:
			if d.Kind != t.Kind {
				return true
			}
		}
	}
	if d := t.Right(); d.HP > 0 {
		switch d.Kind {
		case KindElf, KindGnome:
			if d.Kind != t.Kind {
				return true
			}
		}
	}
	if d := t.Down(); d.HP > 0 {
		switch d.Kind {
		case KindElf, KindGnome:
			if d.Kind != t.Kind {
				return true
			}
		}
	}
	return false
}

func (t *Tile) Attack() {
	if t.HP < 1 {
		return
	}

	var lowest *Tile

	for _, offset := range [][]int{
		{0, -1},
		{-1, 0},
		{1, 0},
		{0, 1},
	} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil {
			switch n.Kind {
			case KindElf, KindGnome:
				if n.Kind != t.Kind && n.HP > 0 {
					if lowest == nil || n.HP < lowest.HP {
						lowest = n
					}
				}
			}
		}
	}
	if lowest != nil {
		lowest.HP -= t.AttackPower
		//fmt.Println("Hit:", lowest)
		//if lowest.IsDead() {
		//fmt.Println("Dead:", lowest)
		//	lowest.W.SetTile(&Tile{
		//		Kind: KindPlain}, lowest.X, lowest.Y)
		//}

	}
}

func (t *Tile) Range() []*Tile {
	neighbors := make([]*Tile, 0)
	for _, offset := range [][]int{
		{0, -1},
		{-1, 0},
		{1, 0},
		{0, 1},
	} {
		if n := t.W.Tile(t.X+offset[0], t.Y+offset[1]); n != nil &&
			n.Kind == KindPlain {
			neighbors = append(neighbors, n)
		}
	}
	return neighbors
}

// PathNeighborCost returns the movement cost of the directly neighboring tile.
func (t *Tile) PathNeighborCost(to astar.Pather) float64 {
	return 1
	toT := to.(*Tile)
	drag := float64(0)

	if toT.Y < t.Y {
		drag = 0
	} else if toT.Y > t.Y {
		drag = .01
	} else if toT.X < t.X {
		drag = .00001
	} else if toT.X > t.X {
		drag = .00004
	}

	return 1 + drag
}

func (t *Tile) Up() *Tile {
	if n := t.W.Tile(t.X, t.Y+1); n != nil {
		return n
	}
	return nil
}
func (t *Tile) Down() *Tile {
	if n := t.W.Tile(t.X, t.Y-1); n != nil {
		return n
	}
	return nil
}
func (t *Tile) Left() *Tile {
	if n := t.W.Tile(t.X-1, t.Y); n != nil {
		return n
	}
	return nil
}
func (t *Tile) Right() *Tile {
	if n := t.W.Tile(t.X+1, t.Y); n != nil {
		return n
	}
	return nil
}

// PathEstimatedCost uses Manhattan distance to estimate orthogonal distance
// between non-adjacent nodes.
func (t *Tile) PathEstimatedCost(to astar.Pather) float64 {
	toT := to.(*Tile)
	absX := toT.X - t.X
	if absX < 0 {
		absX = -absX
	}
	absY := toT.Y - t.Y
	if absY < 0 {
		absY = -absY
	}
	return float64(absX + absY)
}

func (w *World) Fighters() []*Tile {
	f := make([]*Tile, 0)
	xlen := len(w.Map)
	ylen := len(w.Map[0])

	for y := 0; y < ylen; y++ {
		for x := 0; x < xlen; x++ {
			if w.Map[x][y].Kind == KindElf || w.Map[x][y].Kind == KindGnome {
				f = append(f, w.Map[x][y])
			}
		}
	}
	return f
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
