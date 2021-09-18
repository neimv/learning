use std::thread;
use std::time::Duration;
use std::sync::mpsc;

fn main() {
    println!("Channels");
    let mut n = 10;
    let (tx, rx) = mpsc::channel();

    thread::spawn(move || {
        let v = vec![1,2,3,4,5];

        for _i in v {
            n += 1;
            tx.send(n).unwrap();
            thread::sleep(Duration::from_millis(100));
        }
    });

    for r in rx {
        println!("mensaje desde tx: {}", r);
    }
}
