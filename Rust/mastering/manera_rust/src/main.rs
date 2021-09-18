fn main() {
    println!("La forma de hacer de rust");
    println!("Manejo de memoria segura");
    {
        let s1 = String::from("ownership");
        println!("{}", s1);
    } // se elimina s1, esto es el owner
    // println!("{}", s1); // error

    println!("se puede llamar Error de puntero, sin copy");
    let e1 = String::from("hola");
    let e2 = e1.clone();
    println!("{}", e1);

    let mut s1 = String::from("hola");
    let entero = 2;
    let s1 = regresa_ownership(s1);
    let z = entrega_ownership();
    crea_copia(entero);
    println!("la cadena desde main: {}", s1);
    println!("entrega ownership: {}", z);

    // ----------------------------------------------------------------------
    // Borrowin y referencias
    // ---------------------------------------------------------------------
    let len = length_cuadrado(&s1);
    println!("texto: {}, tamaño^2: {}", s1, len);

    let ref1 = &s1;
    println!("{}", ref1);

    let s = xy();

    println!("Lifetimes");
    let a;
    {
        let b = 4; // comienza el lifetime
        a = b;
    } // muere b
    println!("{}", a);

    let x = 3;
    let z = 7;
    let a = f(&x, &z); // usando &[numero]
    println!("{}", a);
}

fn regresa_ownership(cadena_texto: String) -> String {
    println!("{}", cadena_texto);
    cadena_texto
}

fn crea_copia(entero: i32) {
    println!("{}", entero);
}

fn entrega_ownership() -> String {
    let x = String::from("x");
    x
}

fn length_cuadrado(s: &String) -> usize {
    s.len() * s.len()
}

fn xy() -> String {
    let s = String::from("xx");
    s
}

#[allow(dead_code)]
#[allow(unused_variables)]
fn f<'a>(param1: &'a i32, param2: &'a i32) -> &'a i32 {
    if param1 > param2 {
        return param1;
    } else {
        return param2;
    }
}

