use std::{fs, io, time::Instant};

fn timer(name: &str) -> impl FnOnce() {
    let start = Instant::now();
    let name = name.to_string();

    move || {
        let duration = Instant::now() - start;
        println!("{} took {:?}", name, duration);
    }
}

fn parse_input() -> Result<Vec<u8>, io::Error> {
    let mut file = match fs::read("input.txt") {
        Ok(content) => content,
        Err(err) => {
            eprintln!("File open error: {err}");
            return Err(err);
        }
    };

    file.truncate(
        file.iter()
            .rposition(|&c| !c.is_ascii_whitespace())
            .unwrap(),
    );

    Ok(file)
}

fn part_one(text: &[u8]) -> u32 {
    let mut total = 0;
    let mut i = 0;

    'outer: while i < text.len() {
        if let Some(next_mul) = (i..text.len() - 3).find(|&k| &text[k..k + 4] == b"mul(") {
            i = next_mul + 4;
        } else {
            break;
        }

        let mut num_one = 0;
        let mut num_two = 0;

        while text[i] != b',' {
            let digit = text[i] - b'0';
            if digit <= 9 {
                num_one = num_one * 10 + digit as u32;
            } else {
                i += 1;
                continue 'outer;
            }

            i += 1;
        }

        i += 1;

        while text[i] != b')' {
            let digit = text[i] - b'0';
            if digit <= 9 {
                num_two = num_two * 10 + digit as u32;
            } else {
                i += 1;
                continue 'outer;
            }

            i += 1;
        }

        total += num_one * num_two;
        i += 1;
    }

    total
}

fn part_two(text: &[u8]) -> u32 {
    let mut total = 0;
    let mut i = 0;

    let mut next_dont = text
        .windows(7)
        .position(|window| window == b"don't()")
        .unwrap_or(usize::MAX);

    'outer: while i < text.len() {
        if let Some(next_mul) = text[i..].windows(4).position(|window| window == b"mul(") {
            i += next_mul + 4;
        } else {
            break;
        }

        if i > next_dont {
            if let Some(next_do) = text[i..].windows(4).position(|window| window == b"do()") {
                i += next_do + 4;

                next_dont = text[i..]
                    .windows(7)
                    .position(|window| window == b"don't()")
                    .map_or(usize::MAX, |pos| pos + i);
                continue;
            } else {
                break;
            }
        }

        let mut num_one = 0;
        let mut num_two = 0;

        while text[i] != b',' {
            let digit = text[i] - b'0';
            if digit <= 9 {
                num_one = num_one * 10 + digit as u32;
            } else {
                i += 1;
                continue 'outer;
            }

            i += 1;
        }
        i += 1;

        while text[i] != b')' {
            let digit = text[i] - b'0';
            if digit <= 9 {
                num_two = num_two * 10 + digit as u32;
            } else {
                i += 1;
                continue 'outer;
            }

            i += 1;
        }
        i += 1;

        total += num_one * num_two;
    }

    total
}

fn main() {
    let timer_parse = timer("Parsing");
    let input = match parse_input() {
        Ok(parsed) => parsed,
        Err(_) => return,
    };
    timer_parse();

    let timer_one = timer("Part One");
    let one = part_one(&input);
    timer_one();
    println!("Part One: {one}");

    let timer_two = timer("Part Two");
    let two = part_two(&input);
    timer_two();
    println!("Part Two: {two}");
}
