/**
 * @class       : 04_clases
 * @author      : neimv (neimv@dark-universe)
 * @created     : lunes may 24, 2021 21:52:14 CDT
 * @description : 04_clases
 */

class Calculator {
    constructor() {
        this.valueA = 0;
        this.valueB = 0;
    }

    sum(valueA, valueB) {
        this.valueA = valueA;
        this.valueB = valueB;

        return this.valueA + this.valueB;
    }
}

const calc = new Calculator();
console.log(calc.sum(2, 2));

// generators
function *helloWorld() {
    if(true) {
        yield 'Hello, ';
    }

    if(true) {
        yield 'world';
    }
};

const generatorHello = helloWorld();
console.log(generatorHello.next().value);
console.log(generatorHello.next().value);
console.log(generatorHello.next().value);

