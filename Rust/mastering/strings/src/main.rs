fn main() {
    println!("ascii y utf8");

    let s = "hola, ía";
    println!("{:?}", s.as_bytes());

    println!("String y str");
    let str_: &str = "hola";
    let elem = &str_[0..2];
    println!("{}", str_);
    println!("{}", elem);

    println!("Format text");
    let texto = format!("este es el numero tres: {}", 3);
    println!("{}", texto);

    let mama = "maria";
    let papa = "juan";
    let hija = "alicia";
    let edad_hija = 3;
    let texto2 = format!(
        "{madre} {padre} tiene una hija, su nombre es {alicia}, {alicia} tiene {edad}",
        madre=mama, padre=papa, alicia=hija, edad=edad_hija
    );

    println!("{}", texto2);
}
