use std::collections::HashSet;
use std::fs::File;
use std::io;
use std::io::*;
use std::vec::Vec;

fn main() -> io::Result<()> {
    let f = &File::open("data")?;
    let buf = BufReader::new(f);
    let nums: Vec<i64> = buf
        .lines()
        .map(|line| {
            line.unwrap()
                .trim()
                .parse::<i64>()
                .expect("Wanted a number")
        })
        .collect();
    let total: i64 = nums.iter().sum();
    println!("{:?}", total);

    // Part 2
    let f = nums.iter().cycle().scan(0, |state, &x| {
        *state += x;
        Some(*state)
    });

    let mut seen = HashSet::new();
    for element in f {
        //println!("the value is: {}", element);
        if seen.contains(&element) {
            println!("First Dup:{}", element);
            break;
        } else {
            seen.insert(element);
        }
    }
    Ok(())
}
