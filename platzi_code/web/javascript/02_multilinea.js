/**
 * @class       : 02_multilinea
 * @author      : neimv (neimv@dark-world)
 * @created     : lunes may 24, 2021 20:01:50 CDT
 * @description : 02_multilinea
 */


let lorem = `otra frase epica
ahora es otra frase epica
`

console.log(lorem);

let person = {
    'name': 'oscar',
    'age': 32,
    'country': 'MX'
}

let { name, age, country } = person;
console.log(name, age, country);

let team1 = ['Oscar', 'Julian', 'Ricardo'];
let team2 = ['Valeria', 'Yesica', 'Camila'];

let education = ['David', ...team1, ...team2];

console.log(education);


