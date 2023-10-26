
use regex::Regex;

fn main() {
    println!("Hello, world!");

    // Expresiones regulares
    let re_add = Regex::new(r"(\d+)\s?\+\s?(\d+)").unwrap();
    let re_mult = Regex::new(r"(\d+)\s?\*\s?(\d+)").unwrap();

    // Traer datos del usuario
    println!("Por favor introduce tu expresion: ");
    let mut expresion = String::new();
    std::io::stdin().read_line(&mut expresion).unwrap();

    // multiplicacion
    loop {
        // Validaciones de operaciones
        let caps = re_mult.captures(expresion.as_str());

        if caps.is_none() {
            break;
        }

        let caps = caps.unwrap();

        let cap_expression = caps.get(0).unwrap().as_str();
        let left_value: i32 = caps.get(1).unwrap().as_str().parse().unwrap();
        let right_value: i32 = caps.get(2).unwrap().as_str().parse().unwrap();
        println!("{:?}, izq: {}, der: {}", caps, left_value, right_value);

        let addition = left_value * right_value;

        expresion = expresion.replace(cap_expression, &addition.to_string());
    }

    // suma
    loop {
        // Validaciones de operaciones
        let caps = re_add.captures(expresion.as_str());

        if caps.is_none() {
            break;
        }

        let caps = caps.unwrap();

        let cap_expression = caps.get(0).unwrap().as_str();
        let left_value: i32 = caps.get(1).unwrap().as_str().parse().unwrap();
        let right_value: i32 = caps.get(2).unwrap().as_str().parse().unwrap();
        println!("{:?}, izq: {}, der: {}", caps, left_value, right_value);

        let addition = left_value + right_value;

        expresion = expresion.replace(cap_expression, &addition.to_string());
    }

    // Mostrar resultado
    println!("resultado: {}", expresion)
}
