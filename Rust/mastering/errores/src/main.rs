use std::fs::File;
use std::io::ErrorKind;
use std::io;
use std::io::Read;

fn main() {
    println!("Errores recuperables");
    let f = File::open("f.txt");

    let f = match f {
        Ok(file) => file,
        Err(err) => match err.kind(){
            ErrorKind::NotFound => match File::create("f.txt") {
                Ok(file_created) => file_created,
                Err(err2) => panic!("{:?}", err2),
            }
            _ => panic!("Hubo un problema"),
        },
    };

    println!("{:?}", f);

    let f2 = File::open("f2.txt").unwrap();
    println!("{:?}", f2);

    let f3 = File::open("f3.txt").expect("Ups!");
    println!("{:?}", f3);

    let f4 = read_file();
    println!("{:?}", f4);
}

fn read_file() -> Result<String, io::Error> {
    // let mut f = match File::open("f4.txt") {
    //     Ok(f) => f,
    //     Err(e) => return Err(e),
    // };
    let mut f = File::open("f4.txt")?;
    let mut s = String::new();

    // match f.read_to_string(&mut s) {
    //     Ok(_) => Ok(s),
    //     Err(e) => return Err(e),
    // }
    f.read_to_string(&mut s)?;
    Ok(s)
}

