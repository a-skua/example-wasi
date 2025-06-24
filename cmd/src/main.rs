wit_bindgen::generate!({
    world: "root",
    path: "./wit",
    generate_all,
});

use example::calculator::calc;
use std::env;

fn main() -> Result<(), ()> {
    let args: Vec<String> = env::args().collect();

    if args.len() != 3 {
        eprintln!("Usage: {} <num1> <num2>", args[0]);
        return Err(());
    }

    let num1: i32 = match args[1].parse() {
        Ok(n) => n,
        Err(_) => {
            eprintln!("Error: First argument '{}' is not a valid number", args[1]);
            return Err(());
        }
    };

    let num2: i32 = match args[2].parse() {
        Ok(n) => n,
        Err(_) => {
            eprintln!("Error: Second argument '{}' is not a valid number", args[2]);
            return Err(());
        }
    };

    let result = calc::add(num1, num2);
    println!("{} + {} = {}", num1, num2, result);

    Ok(())
}
