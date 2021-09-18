
#[derive(Debug)]
struct Persona {
    nombre: String,
    edad: u8,
    altura: u8,
    email: Option<String>,
}

#[derive(Debug)]
struct Punto(i32, i32);

#[derive(Debug)]
struct Unidad;

enum Direccion {
    Arriba,
    Abajo,
    Izquierda,
    Derecha,
    Noreste,
    Sureste,
}

#[derive(Debug)]
enum Logo {
    Descripcion(String),
    Color(u8, u8, u8),
    Archivo(File),
}

#[derive(Debug)]
struct File {}

fn main() {
    println!("struct");
    let mut juan = Persona {
        nombre: String::from("juan"),
        edad: 25,
        altura: 180,
        email: Some("juan@email.com".to_string()),
    };
    juan.nombre = "Juan".to_string();
    let ramon = Persona {
        nombre: "Ramon".to_string(),
        email: Some("ramon@email.com".to_string()),
        ..juan
    };

    println!("Struct juan: {:?}", juan);
    println!("Struct juan: {:?}", ramon);

    let punto = Punto(34, -62);
    println!("struct punto: {:?}", punto);

    let u = Unidad;
    println!("struct unidad: {:?}", u);

    // Enums y option
    println!("enum");
    let arriba = Direccion::Noreste;
    match arriba {
        Direccion::Arriba => println!("Arriba"),
        Direccion::Abajo => println!("Abajo"),
        Direccion::Izquierda => println!("Izquierda"),
        Direccion::Derecha => println!("Derecha"),
        _ => println!("No se a donde ir")
    }

    let color = Logo::Color(23, 55, 100);
    let descripcion = Logo::Descripcion("mi logo".to_string());
    let v = vec![color, descripcion];
    println!("{:?}", v);

    let mut pedro = Persona {
        nombre: "Pedro".to_string(),
        edad: 23,
        altura: 170,
        email: None,
    };
    pedro.email = Some("pedro@email.com".to_string());

    println!("El email es: {}", pedro.email.unwrap_or("Pedro no tiene email".to_string()));

    let v = vec![0, 3, 34, 53, 53];
    let valor_alto = v.get(30);
    match valor_alto {
        Some(cualquier_cosa) => println!("EL numero es {}", cualquier_cosa),
        None => println!("Fuera de rango"),
    }

    // Arrays
    println!("Arrays");
    let a: [u8; 50] = [0; 50];
    println!("primer elemento = {}, ultimo elemento = {} y de tamaño = {}",a[0], a[a.len() - 1], a.len());
    println!("tamaño en bytes: {}", std::mem::size_of_val(&a));

    // Recorriendo el array
    for i in 0..a.len() {
        println!("indice: {}, valor: {}", i, a[i]);
    }

    // slices, sin tamaño en tiempo de compilacion
    let mut a = [34, 56, 12, 99, 82, 300];
    let slice = &a[..4];
    let get = get_slice(&a);
    println!("slice 1: {:?}", slice);
    println!("slice 1: {:?}", get);
    modifica(&mut a);
    println!("slice 1: {:?}", a);

    // tuplas, guardan diferentes tipos
    let t1 = (4, -5);
    let t2 = (3, 23, false, t1);
    println!("{:?}", t2);

    let t = retorna_tupla(2, 4);
    println!("{:?}", t);
}

fn get_slice(s: &[i32]) -> &[i32] {
    &s[2..5]
}

fn modifica(s: &mut[i32]) {
    s[0] = 0;
}

fn retorna_tupla(x: u8, y: u8) -> (u8, u8) {
    (x, y)
}

