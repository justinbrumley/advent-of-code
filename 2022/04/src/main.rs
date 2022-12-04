use std::fs;

fn main() {
    let input = fs::read_to_string("./input").expect("Should have been able to read the file");

    let split = input.split("\n");
    let lines: Vec<&str> = split.collect();

    // Track total number of _complete_ overlaps
    let mut overlaps = 0;
    let mut i = 0;

    loop {
        if i >= lines.len() {
            break;
        }

        if lines[i] == "" {
            i += 1;
            continue;
        }

        // Leaving a comment to say I don't actually know rust,
        // and all this splitting and parsing code seems crazy to me
        // (which means there's probably an easier way)

        let pair: Vec<&str> = lines[i].split(",").collect();

        let v1 = pair[0];
        let mut sections: Vec<&str> = v1.split("-").collect();
        let min1 = sections[0].parse::<i32>().unwrap();
        let max1 = sections[1].parse::<i32>().unwrap();

        let v2 = pair[1];
        sections = v2.split("-").collect();
        let min2 = sections[0].parse::<i32>().unwrap();
        let max2 = sections[1].parse::<i32>().unwrap();

        // Check if two sections DON'T overlap, then invert that
        if !(max1 < min2 || min1 > max2) {
            overlaps += 1;
        }

        i += 1;
    }

    println!("Found {overlaps} overlaps")
}
