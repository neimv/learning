const { count } = require("console");

var tabla = [
    {
        a: 1,
        b: 'Z'
    }, {
        a: 2,
        b: 'Otra'
    }, {
        a: 3,
        c: 'kiubo'
    }
]

console.log('Hola');
console.info('info');
console.error('Se rompio');
console.warn('warning');
console.table(tabla)

console.group('Conversacion');
console.log('Hola');
console.log('blablabla');
console.log('Adios');
console.groupEnd('Conversacion');

console.count('veces');
console.count('veces');
console.count('veces');
console.count('veces');
console.countReset('veces');
console.count('veces');