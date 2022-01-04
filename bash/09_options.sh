# !/bin/bash

echo "Programa opciones"
echo "Opcion 1 enviada: $1"
echo "Opcion 2 enviada: $2"
echo "Opciones enviadas: $*"
echo -e "\n"
echo "Recuperar los valores"

while [ -n "$1" ]
do
    case "$1" in
        -a) echo "-a option utilizada";;
        -b) echo "-b option utilizada";;
        -c) echo "-c option utilizada";;
        *) echo "$1 no es una opcion";;
    esac
    shift
done

