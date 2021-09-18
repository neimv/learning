var hello = "Hello";
var hello = "hello +" // no se debe de hacer
let world = "Hello World";
const helloWorld = "Hello World!";

const anotherFunction = () => {
    console.log(hello);
    console.log(world);
    console.log(helloWorld);
}

anotherFunction();

const helloWorldFunc = () => {
    globalVar = "I'm global";
}

helloWorldFunc();
console.log(globalVar);

const anotherFunction2 = () => {
    var localVar = globalVar = "I'm global";
}

anotherFunction2();
console.log(globalVar)
