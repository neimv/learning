
function hola(nombre, miCallback) {
    setTimeout(() => {
        console.log("Hola, " + nombre);
        miCallback();
    }, 1000);
}

function adios(nombre, otroCallback) {
    setTimeout(() => {
        console.log("Adios " + nombre)
        otroCallback();
    }, 1000)
}

console.log("iniciando proceso...");
hola('Neimv', () => {
    adios('Neimv', () => {
        console.log("terminando proceso...");
    });
});
