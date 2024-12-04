use regex::Regex;
fn main() {
    part1();
    part2();
}

fn part1() {
    let input = include_str!("../../day3-input.txt");
    let mut sum = 0;
    let regex = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)").expect("invalid regex");
    for capture in regex.captures_iter(input) {
        let a = capture.get(1).expect("not match").as_str().parse::<i32>().expect("not a number");
        let b = capture.get(2).expect("not match").as_str().parse::<i32>().expect("not a number");
        sum += a * b;
    }
    println!("sum: {}", sum);
}

fn part2() {
    let input = include_str!("../../day3-input.txt");
    let mut sum = 0;
    let mut enabled = true;
    let regex = Regex::new(r"mul\((\d{1,3}),(\d{1,3})\)|do\(\)|don't\(\)").expect("invalid regex");
    for capture in regex.captures_iter(input) {
        let matched_str = capture.get(0).expect("not match").as_str();
        match matched_str {
            "do()" => enabled = true,
            "don't()" => enabled = false,
            _ => {
                if enabled {
                    let a = capture.get(1).expect("not match").as_str().parse::<i32>().expect("not a number");
                    let b = capture.get(2).expect("not match").as_str().parse::<i32>().expect("not a number");
                    sum += a * b;
                }
            },
        }
    }
    println!("sum: {}", sum);
}
