
const anotherFunction = () => {
    for (let i = 0; i < 10; i++) {
        setTimeout(() => {
            console.log(i);
        }, 1000)
    }
}

anotherFunction();


// errores
// ¿Qué representa el siguiente código?
// JavaScript solo utiliza el hoisting en declaraciones, mas no en inicializaciones