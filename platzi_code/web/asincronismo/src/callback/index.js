/**
 * @class       : index
 * @author      : neimv (neimv@dark-universe)
 * @created     : lunes jun 14, 2021 09:46:09 CDT
 * @description : index
 */

function sum(num1, num2) {
    return num1 + num2;
}

function calc(num1, num2, callback) {
    return callback(num1, num2);
}

console.log(calc(2, 2, sum));

function date(callback) {
    console.log(new Date);
    setTimeout(function () {
        let date = new Date;
        callback(date);
    }, 3000)
}

function printDate(dateNow) {
    console.log(dateNow);
}

date(printDate);

