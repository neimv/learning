
trait Animal {
    fn habla(&self);
}

struct Gato {}
struct Perro {}

impl Animal for Gato {
    fn habla(&self) {
        println!("Miau");
    }
}

impl Animal for Perro {
    fn habla(&self) {
        println!("Guau");
    }
}

fn main() {
    println!("Hello, world!");
    let g = Gato{};
    let p = Perro{};

    emite_sonido(g);
    emite_sonido(p);
}

fn emite_sonido(a: impl Animal) {
    a.habla();
}
