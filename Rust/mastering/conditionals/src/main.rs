fn main() {
    let n = -7;
    let a = if n > 0 { true } else { false };

    println!(
        "{}",
        if n > 0 {
            "El número es positivo"
        } else if n == 0 {
            "El número es cero"
        } else {
            "El número no es positivo"
        }
    );
    println!("a = {}", a);

    // ciclos
    let mut i = 1;
    // 1 .. 50
    // no imprimir multiplos de 3 (i % 3 != 0)
    // i*i < 400
    while i <= 50 {
        i += 1;

        if i % 3 == 0 { continue; }
        if i * i >= 400 { break; }
        println!("{}", i * i);
    }

    i = 0;
    loop { // = while true
        let x = i * i;
        if x >= 200 { break; }
        println!("{}", x);
        i += 1;
    }

    // Ciclo for
    for i in 1..10 {
        if i == 5 { continue; }
        if i > 8 { break; }
        println!("{}^2 = {}", i, i * i);
    }

    // match parecido a switch
    let numero = 4;

    match numero {
        1 => println!("uno"),
        2 => println!("dos"),
        3 => println!("tres"),
        4 | 5 => println!("cuatro o cinco"),
        6..=10 => println!("en rango"),
        _ => println!("desconocido")
    }

    // patrones
    let x = (10, 20);
    match x {
        (0, 0) => println!("es el origen"),
        (0, y) => println!("x es cero y el valor Y = {}", y),
        (x, 0) => println!("y es cero y X = {}", x),
        (x, y) => println!("X = {}, Y = {}", x, y),
        _ => println!("otro caso")
    }

    // maneras especiales de Rust
    let mut x: Option<u8> = Some(20);
    x = None;
    
    /*match x {
        Some(i) => println!("{}", i),
        _ => println!("No hay valor")
    }*/
    if let Some(i) = x{
        println!("{}", i);
    } else {
        println!("No hay valor (if let)");
    }

    // while let
    x = Some(20);
    while let Some(i) = x {
        if i == 0 {
            x = None;
        } else {
            println!("{}", i);
            x = Some(i-1);
        }
    }
}
