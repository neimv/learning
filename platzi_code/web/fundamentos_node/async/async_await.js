
async function hola(nombre) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log("Hola, " + nombre);
            resolve(nombre);
        }, 1500);
    });
}

async function adios(nombre) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log("Adios " + nombre);
            resolve();
        }, 1000);
    });
}

async function hablar(nombre) {
    return new Promise((resolve, reject) => {
        setTimeout(() => {
            console.log("bla bla bla");
            resolve(nombre)
        }, 1000);
    });
}

async function main(){
    let nombre = await hola("Neimv");
    await hablar();
    await hablar();
    await hablar();
    await adios(nombre);
}

main();