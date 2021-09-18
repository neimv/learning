use std::thread;
use std::time::Duration;

fn main() {
    let mut n = 10;
    let handle = thread::spawn(move || {
        let v = vec![1, 2, 3, 4, 5];
        for i in v {
            n += i;
            println!("hilo {} valor {}", i, n);
            thread::sleep(Duration::from_millis(100));
        }
    });
    println!("Paralelismo");

    for _ in 1..5 {
        println!("hilo principal");
        thread::sleep(Duration::from_millis(100));
    }

    handle.join().unwrap();
}
