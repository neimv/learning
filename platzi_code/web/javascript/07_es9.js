/**
 * @class       : 07_es9
 * @author      : neimv (neimv@dark-world)
 * @created     : martes may 25, 2021 17:01:55 CDT
 * @description : 07_es9
 */

const obj = {
    name: 'zatara',
    age: 1000,
    country: 'MX',
}

let {name, ...all} = obj;
console.log(name, all);

const obj2 = {
    ...obj,
    last: 'zatara',
}

console.log(obj2);

const helloWorld = () => {
    return new Promise((resolve, reject) => {
        (true)
            ? resolve('Hello world')
            : reject(new Error('Test Error'))
    });
};

helloWorld()
    .then(response => console.log(response))
    .catch(error => console.log(error))
    .finally(() => console.log('Finalizo'));

