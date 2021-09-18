mod animal {
    fn nadar() {}
    pub mod mamifero {
        pub fn parir() {}
        pub fn amamantar() {}
        pub fn correr () {}
        pub fn comer() {
            correr();
            crate::animal::nadar();
            super::nadar();
        }
    }
    pub mod ave {
        pub fn desovar() {}
    }
    pub mod anfibio {
        pub fn desovar() {}
    }
}

use animal::mamifero;
use animal::anfibio;

fn osa_pare_amamanta() {
    animal::mamifero::parir();
    crate::animal::mamifero::amamantar();
}

fn tortuga_en_la_playa() {
    anfibio::desovar();
}

fn sembrar() {
    let mut arbol = plantas::Arbol::new();
    arbol.numero_hojas = 100000;
}

mod plantas {
    pub struct Arbol {
        pub numero_hojas: i64,
        es_frutal: bool,
    }
    impl Arbol {
        pub fn new() -> Self {
            Arbol {
                numero_hojas: 0,
                es_frutal: false,
            }
        }
    }
}

