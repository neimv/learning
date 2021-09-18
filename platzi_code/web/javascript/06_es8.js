/**
 * @class       : 06_es8
 * @author      : neimv (neimv@dark-world)
 * @created     : martes may 25, 2021 16:05:55 CDT
 * @description : 06_es8
 */

const data = {
    frontend: 'front',
    backend: 'back',
    desing: 'des',
}

const entries = Object.entries(data);
console.log(entries);
console.log(entries.length);

const values = Object.values(data);
console.log(values);
console.log(values.length);

const string = 'hello';
console.log(string.padStart(7, 'hi'));
console.log(string.padEnd(12, ' ----'));
console.log('food'.padEnd(12, ' ----'));

const helloWorld = () => {
    return new Promise((resolve, reject) => {
        (true)
        ? setTimeout(() => resolve('Hello world'), 3000)
        : reject(new Error('test Error'));
    });
}

const helloAsync = async () => {
    const hello = await helloWorld();

    console.log(hello);
}

helloAsync();

const anotherFunction = async () => {
    try {
        const hello = await helloWorld();
        console.log(hello);
    } catch (error) {
        console.log(error);
    }
}

anotherFunction();

