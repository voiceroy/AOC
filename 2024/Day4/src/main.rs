use std::{
    fs,
    io::{self, BufRead},
    time::Instant,
};

fn timer(name: &str) -> impl FnOnce() {
    let start = Instant::now();
    let name = name.to_string();

    move || {
        let duration = Instant::now() - start;
        println!("{} took {:?}", name, duration);
    }
}

fn parse_input() -> Result<Vec<String>, io::Error> {
    let file = match fs::read("input.txt") {
        Ok(content) => content,
        Err(err) => {
            eprintln!("File open error: {err}");
            return Err(err);
        }
    };

    Ok(file
        .lines()
        .map(|line| line.unwrap())
        .collect::<Vec<String>>())
}

fn part_one(text: &[String]) -> u32 {
    text.iter() // Horizontal
        .map(|line| (line.matches("XMAS").count() + line.matches("SAMX").count()) as u32)
        .sum::<u32>()
        + (0..text.len() - 3) // Vertical
            .map(|i| {
                (0..text[0].len())
                    .map(|col| {
                        (text[i].as_bytes()[col] == b'X'
                            && text[i + 1].as_bytes()[col] == b'M'
                            && text[i + 2].as_bytes()[col] == b'A'
                            && text[i + 3].as_bytes()[col] == b'S') as u32
                            + (text[i].as_bytes()[col] == b'S'
                                && text[i + 1].as_bytes()[col] == b'A'
                                && text[i + 2].as_bytes()[col] == b'M'
                                && text[i + 3].as_bytes()[col] == b'X')
                                as u32
                    })
                    .sum::<u32>()
            })
            .sum::<u32>()
        + (0..(2 * text.len() - 1)) //
            .map(|k| {
                (0..text.len())
                    .map(|row| {
                        let col = k as isize - row as isize;

                        if col >= 0
                            && col < text.len() as isize
                            && row + 3 < text.len()
                            && col as usize + 3 < text.len()
                        {
                            (&[
                                text[row].as_bytes()[col as usize],
                                text[row + 1].as_bytes()[(col + 1) as usize],
                                text[row + 2].as_bytes()[(col + 2) as usize],
                                text[row + 3].as_bytes()[(col + 3) as usize],
                            ] == b"XMAS") as u32
                                + (&[
                                    text[row].as_bytes()[col as usize],
                                    text[row + 1].as_bytes()[(col + 1) as usize],
                                    text[row + 2].as_bytes()[(col + 2) as usize],
                                    text[row + 3].as_bytes()[(col + 3) as usize],
                                ] == b"SAMX") as u32
                        } else {
                            0
                        }
                    })
                    .sum::<u32>()
            })
            .sum::<u32>()
        + (0..(2 * text.len() - 1))
            .map(|k| {
                (0..text.len())
                    .map(|row| {
                        let col = k as isize - (text.len() - 1 - row) as isize;

                        if col >= 0 && col < text.len() as isize && row + 3 < text.len() && col >= 3
                        {
                            (&[
                                text[row].as_bytes()[col as usize],
                                text[row + 1].as_bytes()[(col - 1) as usize],
                                text[row + 2].as_bytes()[(col - 2) as usize],
                                text[row + 3].as_bytes()[(col - 3) as usize],
                            ] == b"XMAS") as u32
                                + (&[
                                    text[row].as_bytes()[col as usize],
                                    text[row + 1].as_bytes()[(col - 1) as usize],
                                    text[row + 2].as_bytes()[(col - 2) as usize],
                                    text[row + 3].as_bytes()[(col - 3) as usize],
                                ] == b"SAMX") as u32
                        } else {
                            0
                        }
                    })
                    .sum::<u32>()
            })
            .sum::<u32>()
}

fn part_two(text: &[String]) -> u32 {
    (1..text.len() - 1)
        .map(|i| {
            (1..text[0].len() - 1)
                .map(|j| match text[i].as_bytes()[j] == b'A' {
                    true => {
                        ((text[i - 1].as_bytes()[j - 1] == b'M'
                            && text[i - 1].as_bytes()[j + 1] == b'S'
                            && text[i + 1].as_bytes()[j - 1] == b'M'
                            && text[i + 1].as_bytes()[j + 1] == b'S')
                            || (text[i - 1].as_bytes()[j - 1] == b'M'
                                && text[i - 1].as_bytes()[j + 1] == b'M'
                                && text[i + 1].as_bytes()[j - 1] == b'S'
                                && text[i + 1].as_bytes()[j + 1] == b'S')
                            || (text[i - 1].as_bytes()[j - 1] == b'S'
                                && text[i - 1].as_bytes()[j + 1] == b'M'
                                && text[i + 1].as_bytes()[j - 1] == b'S'
                                && text[i + 1].as_bytes()[j + 1] == b'M')
                            || (text[i - 1].as_bytes()[j - 1] == b'S'
                                && text[i - 1].as_bytes()[j + 1] == b'S'
                                && text[i + 1].as_bytes()[j - 1] == b'M'
                                && text[i + 1].as_bytes()[j + 1] == b'M'))
                            as u32
                    }
                    false => 0,
                })
                .sum::<u32>()
        })
        .sum::<u32>()
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
