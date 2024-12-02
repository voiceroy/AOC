use std::{fs, io, time::Instant};

fn timer(name: &str) -> impl FnOnce() {
    let start = Instant::now();
    let name = name.to_string();

    move || {
        let duration = Instant::now() - start;
        println!("{} took {:?}", name, duration);
    }
}

fn parse_input() -> Result<Vec<Vec<i32>>, io::Error> {
    let file = match fs::read_to_string("input.txt") {
        Ok(content) => content,
        Err(err) => {
            eprintln!("File open error: {err}");
            return Err(err);
        }
    };

    Ok(file
        .trim()
        .lines()
        .map(|x| x.split_whitespace())
        .map(|x| {
            x.map(|y| str::parse::<i32>(y).unwrap())
                .collect::<Vec<i32>>()
        })
        .collect::<Vec<Vec<i32>>>())
}

fn part_one<T>(lists: &[T]) -> i32
where
    T: AsRef<[i32]>,
{
    lists
        .iter()
        .map(|x| {
            let x = x.as_ref();
            let mut increasing = x[1] - x[0];
            if increasing.abs().clamp(1, 3) == increasing.abs() {
                increasing = increasing.signum();

                (x.len() - 2
                    == (1..x.len() - 1)
                        .map(|i| x[i + 1] - x[i])
                        .take_while(|&v| {
                            v != 0 && v.abs().clamp(1, 3) == v.abs() && v.signum() == increasing
                        })
                        .count()) as i32
            } else {
                0
            }
        })
        .sum()
}

fn part_two(lists: &[Vec<i32>]) -> i32 {
    lists
        .iter()
        .map(|x| {
            if part_one(&[&x]) == 1
                || (0..x.len())
                    .map(|i| {
                        let mut v = x.clone();
                        v.remove(i);
                        v
                    })
                    .any(|v| part_one(&[v]) == 1)
            {
                1
            } else {
                0
            }
        })
        .sum()
}

fn main() {
    let input = match parse_input() {
        Ok(parsed) => parsed,
        Err(_) => return,
    };

    let timer_one = timer("Part One");
    let one = part_one(&input);
    timer_one();
    println!("Part One: {one}");

    let timer_two = timer("Part Two");
    let two = part_two(&input);
    timer_two();
    println!("Part Two: {two}");
}
