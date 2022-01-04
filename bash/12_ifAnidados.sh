# !/bin/bash

notaClase=0
continua=""

echo "Ejemplo sentencia if-else"
read -n1 -p "Indica cual es su nota (1-9): " notaClase
echo -e "\n"

if [ $notaClase -ge 7 ]; then
    echo "El alumno aprueba la materia"
    read -p "Si va a continuar estudiando en el siguiente nivel (s/n): " continua

    if [ $continua = "s" ]; then
        echo "Bienvenido al siguiente nivel"
    else
        echo "gracias por trabajar con nosotros"
    fi
else
    echo "el alumno reprueba la materia"
fi
