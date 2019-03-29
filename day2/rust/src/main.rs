use edit_distance::edit_distance;
use itertools::Itertools;
use std::iter::Iterator;

extern crate edit_distance;

fn main() {
    let puzzle = include_str!("data");
    let lines = puzzle.lines();
    let mut three = 0;
    let mut two = 0;

    for l in lines {
        let mut chars = l.chars().collect::<Vec<char>>();
        chars.sort_by(|a, b| b.cmp(a));
        let maps = chars
            .iter()
            .map(|c| (c, 1))
            .coalesce(|(c, n), (d, m)| {
                if c == d {
                    Ok((c, n + m))
                } else {
                    Err(((c, n), (d, m)))
                }
            })
            .collect::<Vec<_>>();

        let mut th = false;
        let mut tw = false;
        for (_, v) in maps {
            if v == 3 {
                th = true
            }
            if v == 2 {
                tw = true
            }
        }
        if th {
            three += 1;
        }

        if tw {
            two += 1;
        }
    }
    println!("Part #1: {}", two * three);

    'outer: for l1 in puzzle.lines() {
        for l2 in puzzle.lines() {
            if edit_distance(l1, l2) == 1 {
                print!("{}\n{}\n", l1, l2);
                break 'outer;
            }
        }
    }
}
