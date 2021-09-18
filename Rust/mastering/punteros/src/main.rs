use std::rc::Rc;

enum Lista {
    Nodo(i32, Box<Lista>),
    None,
}

enum Lista2 {
    Nodo(i32, Rc<Lista2>),
    None,
}

use Lista::*;

fn main() {
    println!("Box");
    let a = 10;
    let b = Box::new(10);

    if a == *b{
        println!("A y B son iguales");
    }

    let lista = Nodo(2, Box::new(Nodo(4, Box::new(Nodo(7, Box::new(None),)))));

    println!("Reference count");

    let listaA = Rc::new(Lista2::Nodo(
        2, Rc::new(
            Lista2::Nodo(
                4, Rc::new(Lista2::Nodo(7, Rc::new(Lista2::None),))
            )
        )
    ));
    println!("contador de referencias listaA: {}", Rc::strong_count(&listaA));
    let listaB = Lista2::Nodo(45, Rc::clone(&listaA));
    let listaC = Lista2::Nodo(45, Rc::clone(&listaA));
    println!("contador de referencias listaA: {}", Rc::strong_count(&listaA));
}
