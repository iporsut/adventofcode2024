fn main() {
    let input = include_str!("day1-input.txt");
    let mut sum = 0;
    let mut left: Vec<i32> = Vec::new();
    let mut right: Vec<i32> = Vec::new();
    let s: Vec<Vec<&str>> = input.lines().map(|l| l.split("   ").collect::<Vec<&str>>()).collect();
    println!("{:?}", s);
    // for line in input.lines() {
    //     let words: Vec<&str> = line.split("   ").collect();
    //     for w in words {
    //         let x = w.parse::<i32>().unwrap();
    //         left.push(x);
    //         right.push(x);
    //     }
    // }
    // println!("left {:?}", left);
    // println!("right {:?}", right);
}
