function hola(nombre) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log("Hola, " + nombre);
            resolve(nombre);
        }, 1500);
    });
}

function adios(nombre) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log("Adios " + nombre)
            resolve();
        }, 1000);
    });
}

function hablar(nombre) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log("bla bla bla");
            resolve(nombre)
        }, 1000);
    });
}

console.log("iniciando el proceso...");
hola('Neimv')
    .then(hablar)
    .then(hablar)
    .then(hablar)
    .then(adios)
    .then((nombre) => {
        console.log("terminado el proceso");
    });