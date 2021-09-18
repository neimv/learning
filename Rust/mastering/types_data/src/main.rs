
use std::mem;

fn main() {
    // Tipos de datos
    let x: u8 = 255; // i = integer with sign, 8 = 8bits
    // u8 integer without sign
    // i8, i16, i32, i64
    // u8, u16, u32, u64
    // isize, usize
    
    let y: usize = 10000000000;
    let size = mem::size_of_val(&y);
    println!("x = {}, y = {} OS {}bits", x, y, size*8);

    // punto flotante
    // f32 f64
    let x: f64 = 34.83;
    let hexadecimal = 0x2B;
    let octal = 0o10;
    let binario = 0b101010101;
    println!("x = {}", x);
    println!("Hexadecimal = {}, octal = {}, binario = {}", hexadecimal, octal, binario);
    println!("Tamaño de binario = {}", mem::size_of_val(&binario));

    let c: char = 'x';
    println!("x = {}, tamaño = {}", c, mem::size_of_val(&c));

    let b: bool = true;
    println!("b = {}, tamaño = {}", b, mem::size_of_val(&b));

    // variables, constantes, shadowing
    // Todas las variables son inmutables
    let a: i32 = 3456;      // declaracion de variables con tipo
    let b = "hola";            // declaracion de variables sin tipo
    // shadowing
    let b = b.len();
    const SIZE: i8 = 5;
    println!("{} {} {}", a, b, SIZE);

    linea();
    linea();
    linea();

    suma(1, 2);

    let x = suma_ret(1, 2);
    println!("{}", x);
}

// declaracion de una funcion
fn linea() {
    println!("hello");
}

fn suma(a: i8, b: i8) {
    println!("{}", a + b);
}

fn suma_ret(a: i8, b: i8) -> i8 {
    a + b
}

