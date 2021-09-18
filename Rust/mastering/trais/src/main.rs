
use std::fmt::Debug;

trait Mostrar {
    fn mostrar(&self) -> String;
}

impl Mostrar for i32 {
    fn mostrar(&self) -> String {
        format!("entero cuatro bytes: {}", self)
    }
}

impl Mostrar for f64 {
    fn mostrar(&self) -> String {
        format!("decima ocho bytes: {}", self)
    }
}

trait Sonido {
    fn hace(&self) -> String;
    fn get(&self) -> &Self {
        &self
    }
}

#[derive(Debug)]
struct Gato {}
#[derive(Debug)]
struct Perro {}
#[derive(Debug)]
struct Perro2 { nombre: String }
#[derive(Debug)]
struct Tractor {}

impl Sonido for Gato {
    fn hace(&self) -> String {
        "Miau".to_string()
    }
}

impl Sonido for Perro {
    fn hace(&self) -> String {
        "Guau".to_string()
    }
}

impl Sonido for Perro2 {
    fn hace(&self) -> String {
        "Guau".to_string()
    }
}

impl Sonido for Tractor {
    fn hace(&self) -> String {
        "Brrrrrrr".to_string()
    }
}

impl Sonido for i32 {
    fn hace(&self) -> String {
        match &self {
            1 => "U-ENE-O".to_string(),
            2 => "DE-O-ESE".to_string(),
            3 => "TE-ERRE-E-ESE".to_string(),
            _ => "ni idea".to_string()
        }
    }
}

// POO
#[derive(Debug)]
struct Gato2 {
    nombre: String,
}

impl Sonido for Gato2 {
    fn hace(&self) -> String {
        "Miau".to_string()
    }
}

impl Gato2 {
    fn new() -> Self {
        Self {
            nombre: "Minino".to_string(),
        }
    }

    fn hace(&self) -> String {
        format!("{} hace miau", self.nombre)
    }
}

#[allow(dead_code)]
enum SimpleEnum<T> {
    OpcionA(T),
    OpcionB(T),
    OpcionC(T),
}

struct Punto<T, U> {
    x: T,
    y: U,
}

impl<T, U> Punto<T, U> {
    fn get_x(&self) -> &T {
        &self.x
    }
    fn get_y(&self) -> &U {
        &self.y
    }
}

#[allow(unused_variables)]
fn main() {
    println!("Traits");
    let uno = 1;
    let dos = 2;
    let x = 44;
    let y = 2.7;
    let gato = Gato {};
    let perro = Perro {};
    let tractor = Tractor {};
    let v: Vec<&Mostrar> = vec![&x, &y];

    for d in v.iter() {
        println!("muestra {}", d.mostrar());
    }

    println!("El perro hace {}", perro.hace());
    println!("El gato hace {}", gato.hace());
    println!("El tractor hace {}", tractor.hace());
    println!("El uno hace {}", uno.hace());
    println!("El dos hace {}", dos.hace());
    println!("La x hace {}", x.hace());

    println!("POO");
    let gato2 = Gato2::new();
    println!("{}", gato2.hace());

    println!("Enums");
    let x = SimpleEnum::OpcionA(60);
    let y = SimpleEnum::OpcionB("Soy una cadena de texto".to_string());

    println!("Generics");
    let punto = Punto{x: 1.10, y: -40};
    println!("x: {} y: {}", punto.x, punto.y);

    println!("funciones genericas");
    let gato3 = Gato2 {
        nombre: "min".to_string(),
    };
    let perro = Perro2 {
        nombre: "solovino".to_string(),
    };
    emite_sonido(gato3, perro);

    let a = Punto {x: 1, y: 3};
    println!("x: {}, y: {}", a.get_x(), a.get_y());
}

fn emite_sonido<T, U>(emisor: T, emisor2: U) 
where T: Sonido + Debug,
      U: Sonido + Debug {
    println!("{:?}", emisor.get());
    println!("{:?}", emisor2.get());
}
