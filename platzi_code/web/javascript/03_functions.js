/**
 * @class       : 03_functions
 * @author      : neimv (neimv@dark-world)
 * @created     : lunes may 24, 2021 20:22:09 CDT
 * @description : 03_functions
 */

let name = 'oscar';
let age = 32;

obj = {
    name,
    age
};
console.log(obj);

const names = [
    {'name': 'oscar', 'age': 32},
    {'name': 'yesica', 'age': 27}
]

let listOfNames = names.map(item => console.log(item.name));

const listOfNames2 = (name, age, country) => {
    //...
    console.log('hello');
}

const square = num => num * num;

const helloPromise = () => {
    return new Promise((resolve, reject) => {
        if (true) {
            resolve('Hey!');
        } else {
            reject('Ups');
        }
    });
}

helloPromise()
    .then(response => console.log(response))
    .catch(error => console.log(error))

