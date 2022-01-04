# !/bin/bash

arregloNumeros=(1 2 3 4 5 6)

echo "Iterar en la lista de numeros"

for num in ${arregloNumeros[*]}; do
    echo "numero: $num"
done

echo "Iterar en la lista de cadenas"

for nom in "Marco" "Pedro" "Luis" "Daniela"; do
    echo "Nombres: $nom"
done

echo "Iterar en archivos"

for fil in *; do
    echo "Nombre archivo $fil"
done

echo "Iterar utilizando un comando"
for fil in $(ls); do
    echo "Nombre archivo $fil"
done

echo "Iterando usando el formato tradicional"
for ((i=1; i<10; i++)) {
    echo "Numero: $i"
}
