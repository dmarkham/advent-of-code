use std::collections::HashSet;
use std::vec::Vec;

fn main() {
    let puzzle = include_str!("data");
    let nums = puzzle
        .split_whitespace()
        .map(|c| c.parse::<i64>().unwrap())
        .collect::<Vec<_>>();

    let total: i64 = nums.iter().sum();
    println!("{:?}", total);

    // Part 2
    let f = nums.iter().cycle().scan(0, |state, &x| {
        *state += x;
        Some(*state)
    });

    let mut seen = HashSet::new();
    seen.insert(0);
    for element in f {
        //println!("the value is: {}", element);
        if seen.contains(&element) {
            println!("First Dup:{}", element);
            break;
        } else {
            seen.insert(element);
        }
    }
}
