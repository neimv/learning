/**
 * @class       : index
 * @author      : neimv (neimv@dark-world)
 * @created     : viernes may 21, 2021 18:32:35 CDT
 * @description : index
 */

// function newFunction(name, age, country) {
//     var name = name || 'ella';
//     var age = age || 32;
//     var country || 'MX';

//     console.log(name, age, country);
// }

// es6
function newFunction2(name='ella', age=32, country='MX') {
    console.log(name, age, country);
}

newFunction2();
newFunction2('otro', 35, 'CO');

let hello = "Hello";
let world = "World";
let epicPhrase = hello + " " + world;
console.log(epicPhrase);

// literal template
let epicPhrase2 = `${hello} ${world}`;
console.log(epicPhrase2);

