fn main() {
    println!("Closures");
    let i = 5;
    let suma_uno = |n| n + 1;
    println!("{}", suma_uno(9));
    let hola = dice_hola;
    hola();

    let suma_dos = |n: u8| -> u8 {n + 2};
    println!("{}", suma_dos(90));

    let suma_i = |n| {n + i};
    println!("{}", suma_i(9));

    // FnOnce
    let mut dict = std::collections::HashMap::new();
    dict.insert(1, "uno");
    dict.insert(2, "dos");
    dict.insert(3, "tres");

    let cl = || {
        for (k, v) in &dict {
            println!("K: {}, v: {}", k, v);
        }
    };
    cl();
    cl();

    // fnMut
    let mut i = 0;
    let mut inc = || {
        i += 1;
        i
    };

    println!("{}", inc());

    let mut x = vec![1, 2, 3];
    let mut y = || x.pop();
    println!("{:?}", y());
    println!("{:?}", x);

    println!("iteradores");
    let n = 10;
    
    for i in 0..n { // Range<i32>
        println!("{}", i);
    }

    let v = vec![1, 2, 3];
    let mut iter = v.iter();
    assert_eq!(iter.next(), Some(&1));
    println!("{:?}", iter.next());

    println!("Adapters");
    let v = vec![1, 2, 3];
    let vec_suma: Vec<i32> = v.iter().map(|n| n + 10).collect();
    println!("{:?}", vec_suma);

    let animales = "camellos vacas gatos perros elefantes";
    let v_filter: Vec<&str> = animales
        .split_whitespace()
        .filter(|animal| animal != &"gatos")
        .collect();
    println!("{:?}", v_filter);

    let v_consumer = vec![45, 23, 12];
    // let total: u8 = v_consumer.iter().sum();
    // println!("{}", total);

    // let min = v_consumer.iter().min();
    // println!("{:?}", min);

    let max = v_consumer.iter().max();
    println!("{:?}", max);
}

fn dice_hola() {
    println!("hola");
}

