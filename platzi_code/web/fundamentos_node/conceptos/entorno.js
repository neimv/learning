/**
 * @class       : entorno
 * @author      : neimv (neimv@dark-universe)
 * @created     : viernes sep 03, 2021 10:16:02 CDT
 * @description : entorno
 */

let nombre = process.env.NOMBRE || 'sin nombre';
let web    = process.env.WEB || 'no tengo web'

console.log('Hola ' + nombre);
console.log('Mi web es ' + web);

