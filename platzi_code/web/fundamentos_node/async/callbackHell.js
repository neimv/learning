
function hola(nombre, miCallback) {
    setTimeout(() => {
        console.log("Hola, " + nombre);
        miCallback(nombre);
    }, 1500);
}

function adios(nombre, otroCallback) {
    setTimeout(() => {
        console.log("Adios " + nombre)
        otroCallback();
    }, 1000);
}

function hablar(callbackHablar) {
    setTimeout(() => {
        console.log("bla bla bla");
        callbackHablar()
    }, 1000);
}

function conversacion(nombre, veces, callback) {
    if (veces > 0) {
        hablar(() => {
            conversacion(nombre, --veces, callback);
        });
    } else {
        adios(nombre, callback);
    }
}

console.log("iniciando proceso...");
hola('Neimv', (nombre) => {
    conversacion(nombre, 3, () => {
        console.log('Proceso terminado');
    })
});
// hola('Neimv', (nombre) => {
//     hablar(() => {
//         hablar(() => {
//             hablar(() => {
//                 adios(nombre, () => {
//                     console.log("terminando proceso...");
//                 });
//             });
//         });
//     });
// });
