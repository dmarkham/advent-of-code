package main

import (
	"strings"
	"testing"
)

type test struct {
	start string
	end   string
	score int
}

func Test_Maps(t *testing.T) {

	tests := []test{
		test{
			start: `
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
`,
			end: `
#######
#G....#
#.G...#
#.#.#G#
#...#.#
#....G#
#######
`,
			score: 27730,
		},
		test{
			start: `
#######
#G..#E#
#E#E.E#
#G.##.#
#...#E#
#...E.#
#######`,
			end: `
#######
#...#E#
#E#...#
#.E##.#
#E..#E#
#.....#
#######`,
			score: 36334,
		},
		test{
			start: `
#######
#E..EG#
#.#G.E#
#E.##E#
#G..#.#
#..E#.#
#######
`,
			end: `
#######
#.E.E.#
#.#E..#
#E.##.#
#.E.#.#
#...#.#
#######
`,
			score: 39514,
		},
		test{
			start: `
#######
#E.G#.#
#.#G..#
#G.#.G#
#G..#.#
#...E.#
#######
`,
			end: `
#######
#G.G#.#
#.#G..#
#..#..#
#...#G#
#...G.#
#######
`,
			score: 27755,
		},
		test{
			start: `
#######
#.E...#
#.#..G#
#.###.#
#E#G#G#
#...#G#
#######
`,
			end: `
#######
#.....#
#.#G..#
#.###.#
#.#.#.#
#G.G#G#
#######
`,
			score: 28944,
		},
		test{
			start: `
#########
#G......#
#.E.#...#
#..##..G#
#...##..#
#...#...#
#.G...G.#
#.....G.#
#########
`,
			end: `
#########
#.G.....#
#G.G#...#
#.G##...#
#...##..#
#.G.#...#
#.......#
#.......#
#########
`,
			score: 18740,
		},
		test{
			start: `
#####
###G#
###.#
#.E.#
#G###
#####
`,
			end: `
#####
###.#
###.#
#.G.#
#G###
#####
`,
			score: 10030,
		},
		test{
			start: `
#########
#G..G..G#
#.......#
#.......#
#G..E..G#
#.......#
#.......#
#G..G..G#
#########
`,
			end: `
#########
#.......#
#..GGG..#
#..G.G..#
#G..G...#
#......G#
#.......#
#.......#
#########
`,
			score: 27828,
		},
		test{
			start: `
####
##E#
#GG#
####
`,
			end: `
####
##.#
#.G#
####
`,
			score: 13400,
		},
	}
	for _, args := range tests {
		world := ParseWorld(args.start)
		for world.Done == false {
			world = tick(world)
		}
		if world.RenderPath(nil) != strings.TrimSpace(args.end) {
			t.Errorf("Board: \n%v, want: \n%v\n", world.RenderPath(nil), args.end)
			continue
		}
		if world.Score != args.score {
			t.Errorf("Score = %v, want %v", world.Score, args.score)
			continue
		}
	}

}
