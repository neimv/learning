# !/bin/bash

notaClase=0
edad=0

echo "Ejemplo sentencia if-else"
read -n1 -p "Indica cual es su nota (1-9): " notaClase
echo -e "\n"

if (( $notaClase >= 7 )); then
    echo "El alumno aprueba la materia"
else
    echo "el alumno reprueba la materia"
fi

read -p "Indique cual es si edad: " edad

if [ $edad -le 18 ]; then
    echo "La persona no puede votar"
else
    echo "la persona puede votar"
fi
