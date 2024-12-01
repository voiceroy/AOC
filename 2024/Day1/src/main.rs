use std::{
    collections::{BinaryHeap, HashMap},
    fs, io,
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

fn parse_input() -> Result<Vec<Vec<i32>>, io::Error> {
    let file = match fs::read_to_string("input.txt") {
        Ok(content) => content,
        Err(err) => {
            eprintln!("File open error: {err}");
            return Err(err);
        }
    };

    let mut parsed = vec![Vec::with_capacity(1000), Vec::with_capacity(1000)];
    file.trim()
        .lines()
        .map(|x| x.split_whitespace())
        .map(|mut x| {
            (
                str::parse::<i32>(x.next().unwrap()).unwrap(),
                str::parse::<i32>(x.next().unwrap()).unwrap(),
            )
        })
        .for_each(|x| {
            parsed[0].push(x.0);
            parsed[1].push(x.1);
        });

    Ok(parsed)
}

fn part_one(lists: &[Vec<i32>]) -> i32 {
    let mut heap_one = BinaryHeap::from(lists[0].iter().collect::<Vec<&i32>>());
    let mut heap_two = BinaryHeap::from(lists[1].iter().collect::<Vec<&i32>>());

    (0..lists[0].len())
        .map(|_| (heap_one.pop().unwrap() - heap_two.pop().unwrap()).abs())
        .sum()
}

fn part_two(lists: &[Vec<i32>]) -> i32 {
    let mut count: HashMap<i32, i32> = HashMap::with_capacity(lists[1].len());

    lists[1]
        .iter()
        .for_each(|x| *count.entry(*x).or_insert(0) += 1);

    lists[0]
        .iter()
        .map(|x| x * count.get(x).unwrap_or(&0))
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
