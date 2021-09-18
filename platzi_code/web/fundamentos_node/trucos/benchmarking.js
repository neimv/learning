let suma = 0

console.time('TodoListo')
console.time('bucle')
for (let i = 0; i < 1000000000; i++) {
    suma += 1;
}
console.timeEnd('bucle')

console.time('asincrono')
asincrona()
    .then(() => {
        console.timeEnd('asincrono')
    })

let suma2 = 0
console.time('bucle2')
for (let j = 0; j < 1000000000; j++) {
    suma2 += 1;
}
console.timeEnd('bucle2')
console.timeEnd('TodoListo')

function asincrona() {
    return new Promise((resolve) => {
        setTimeout(() => {
            console.log('Termina')
            resolve()
        });
    })
}