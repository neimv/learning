# !/bin/bash

option=""

echo "Ejemplo sentencia case"
read -n1 -p "Ingrese una opcion de la A - Z: " option
echo -e "\n"

case $option in
    "A") echo -e "\n Operacion guardar archivo";;
    "B") echo "Operacion eliminar archivo";;
    [C-E]) echo "No esta implementada la operacion";;
    *) echo "Opcion incorrecta";;
esac
