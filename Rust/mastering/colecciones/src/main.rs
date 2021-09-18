
use std::collections::HashMap;
use std::collections::HashSet;

fn main() {
    println!("Vectores");

    let mut v = Vec::new();
    v.push(2);
    v.push(12);
    v.push(21);
    v.push(121);
    v.push(1212);
    println!("Vector: {:?} {} {}", v, v.len(), v.capacity());

    // iterar el vector
    for i in v {
        println!("item {}", i);
    }

    let mut v2 = Vec::new();
    v2.push(2);
    v2.push(12);
    v2.push(21);
    v2.push(121);
    v2.push(1212);
    let ultimo = v2.pop();
    let segundo = v2.get(10);
    println!("ultimo valor {:?}", ultimo);
    println!("ultimo valor {:?}", segundo);

    let mut v3 = vec![9, 8, 6];
    println!("{:?}", v3);

    let mut v4: Vec<i32> = Vec::new();
    v4.push(2);
    v4.push(12);
    v4.push(21);
    v4.push(121);
    v4.push(1212);
    println!("{:?}", v4);

    let mut v5: Vec<Tipo> = Vec::new();
    #[derive(Debug)]
    enum Tipo {
        Entero(i32),
        Decimal(f64),
        Texto(String)
    }
    v5.push(Tipo::Texto(String::from("hola")));
    v5.push(Tipo::Decimal(2.3));
    v5.push(Tipo::Entero(333));
    v5.push(Tipo::Entero(1));
    println!("{:?}", v5);

    println!("Hashmap");
    let mut num = HashMap::new();
    num.insert(10, String::from("diez"));
    num.insert(20, String::from("veinte"));
    num.insert(30, String::from("treinta"));
    println!("{:?}", num);


    let mut num2 = HashMap::new();
    let diez = String::from("diez");
    let veinte = String::from("veinte");
    let treinta = String::from("treinta");
    let x = 10;
    let y = 20;
    let z = 30;
    num2.insert(x, &diez);
    num2.insert(y, &veinte);
    let q = num2.entry(z).or_insert(&treinta);
    println!("entry {:?}", q);
    println!("{:?} {} {}", num, diez, x);

    for (k, v) in &num2 {
        println!("{} => {}", k, v);
    }

    let elem = num2.get(&x);
    match elem {
        Some(i) => println!("Valor: {}", i),
        _ => println!("No encontre el valor")
    }

    println!("Hash set");
    let mut c1 = HashSet::new();
    c1.insert('a');
    c1.insert('b');
    c1.insert('c');
    c1.insert('d');
    let letra = 'z';
    println!("{:?} contiene {}? {}", c1, letra, c1.contains(&letra));
}

