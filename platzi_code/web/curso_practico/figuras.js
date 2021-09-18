/**
 * @class       : figuras
 * @author      : neimv (neimv@dark-universe)
 * @created     : miércoles jul 21, 2021 10:18:59 CDT
 * @description : figuras
 */

// Codigo del cuadrado
console.group("cuadrado");
// const ladoCuadrado = 5;
// console.log("Los lados del cuadrado miden: " + ladoCuadrado + "cm");

function perimetroCuadrado(lado) {
    return lado * 4;
}
// console.log("El perimetro del cuadrado es: " + perimetroCuadrado(ladoCuadrado) + "cm");

function areaCuadrado(lado) {
    return lado * lado;
}
// console.log("El area del cuadrado es: " + areaCuadrado + "cm^2");
console.groupEnd();

// Codigo del triangulo
console.group("Triangulo");
// const ladoTriangulo1 = 6;
// const ladoTriangulo2 = 6;
// const baseTriangulo = 4;
// const alturaTriangulo = 5.5;

// console.log(
//     "los lados del triangulo miden: " 
//     + ladoTriangulo1 
//     + "cm, " 
//     + ladoTriangulo2 
//     + "cm, " 
//     + baseTriangulo + "cm"
// );
// console.log("La altura del triangulo es: " + alturaTriangulo + "cm");

function perimetroTriangulo(lado1, lado2, base) {
    lado1 + lado2 + base;
}
// console.log("El perimetro del triangulo es: " + perimetroTriangulo + "cm");

function areaTriangulo(base, altura) {
    (base * altura) / 2
}
// console.log("El area del triangulo es: " + areaTriangulo + "cm^2")

console.groupEnd();

// Codigo del circulo
console.group("Circulo");
// const radioCirculo = 4;
function diametroCirculo(radio) {
    return radio * 2;
}
// const PI = Math.PI;
// console.log("El radio del circulo es: " + radioCirculo + "cm, el diametro del circulo es :" + diametroCirculo + "cm");

function perimetroCirculo(radio) {
    return diametroCirculo(radio) * Math.PI;
}
// console.log("El perimetro del circulo es: " + perimetroCirculo + "cm");

function areaCirculo(radio) {
    return (radio * radio) * PI;
}
// console.log("El area del circulo es: " + areaCirculo + "cm^2");

console.groupEnd();

// Triangulo isosceles
console.group("triangulo isoseles");
function alturaTrianguloIsosceles(lado1, lado2, lado3) {
    const miList = [lado1, lado2, lado3]
    const miSet = new Set(miList);

    if (miSet.size === 1) {
        return "todos los lados del triangulo son iguales, no es un isosceles";
    } else if (miSet.size === 3) {
        return "todos los lados del triangulo son diferentes, no es un isosceles";
    } else {
        const counts = {};
        var base;
        var lado;
        const es_valido = [
            (lado1 + lado2) > lado3, (lado1 + lado3) > lado2, (lado2 + lado3) > lado1
        ].every(e => e == true);

        if (!es_valido) {
            return "Los lados no forman un triangulo";
        }

        for (const num of miList) {
            counts[num] = counts[num] ? counts[num] + 1 : 1;
        }

        for (const num of [...miSet]) {
            if (counts[num] == 2) {
                lado = num;
            } else {
                base = num;
            }
        }

        const altura = Math.sqrt((lado ** 2) - ((base / 2) ** 2));

        return altura
    }
}

resp = areaTrianguloIsosceles(1, 2, 2);
console.log(resp)

console.groupEnd();

// aqui interactuamos con el HTML
function calcularPerimetroCuadrado() {
    const input = document.getElementById("InputCuadrado");
    const value = input.value;

    const perimetro = perimetroCuadrado(value);

    alert(perimetro);
}

function calcularAreaCuadrado() {
    const input = document.getElementById("InputCuadrado");
    const value = input.value;

    const area = areaCuadrado(value);
    
    alert(area);
}



