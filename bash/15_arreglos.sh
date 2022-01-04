# !/bin/bash

arregloNumeros=(1 2 3 4 5 6)
arregloCadenas=(Marco, Antonio, Pedro, Susana)
arregloRangos=({A..Z} {10..20})

# Imprimir todos los valores
echo "Arreglo de numeros: ${arregloNumeros[*]}"
echo "Arreglo de cadenas: ${arregloCadenas[*]}"
echo "Arreglo de rngos: ${arregloRangos[*]}"

# Imprimir los tamanios de los arreglos
echo "tamanio arreglo de numeros: ${#arregloNumeros[*]}"
echo "tamanio arreglo de cadenas: ${#arregloCadenas[*]}"
echo "tamanio arreglo de rngos: ${#arregloRangos[*]}"

# imprimir la posicion 3 del arreglo
echo "posicion 3 del arreglo de numeros: ${arregloNumeros[3]}"
echo "posicion 3 del arreglo de cadenas: ${arregloCadenas[3]}"
echo "posicion 3 arreglo de rangos: ${arregloRangos[3]}"

# anadir y eliminar valores en un arreglo
arregloNumeros[7]=20
unset arregloNumeros[0]
echo "arreglo de numeros: ${arregloNumeros[*]}"
echo "tamanio arreglo de numeros ${#arregloNumeros[*]}"
