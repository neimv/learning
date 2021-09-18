process.on('beforeExit', () => {
    console.log("El proceso va a terminar");
})

process.on('exit', () => {
    console.log("El proceso acabó");
})

process.on('uncaughtException', (err, origen) => {
    console.error("Vaya se nos ha olvidado capturar un error");
    console.error(err);
})
// process.on('unhandledRejection')

functionQueNoExiste();

console.log('Esto si el error no se recoje, no sale');