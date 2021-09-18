const { createBrotliCompress } = require("zlib");

function otraFuncion() {
    return seRompe();
}

function seRompe() {
    return 3 + z;
}

function seRompeAsincrona(cb) {
    setTimeout(() => {
        try {
            return 3 + z;
        } catch (err) {
            console.error("error en mi funcion asincrona")
            createBrotliCompress(err)
        }
    });
}

try {
    otraFuncion();
} catch(err) {
    console.error("Vaya, algo se ha roto");
    console.error(err.message);
}

try {
    seRompeAsincrona(() => {
        console.error("hay error");
    });
} catch(err) {
    console.error("Vaya, algo se ha roto");
    console.error(err.message);
}

console.log('Esto de aqui esta al final');