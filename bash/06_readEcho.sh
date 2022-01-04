# !/bin/bash

option=0
backupName=""

echo "Programa Utilidades postgresql"
echo -n "Ingresar una opcion: "
read 
option=$REPLY
echo -n "Ingresar el nombre del archivo del backup: "
read
backupName=$REPLY

echo "Opcion: $option, backup: $backupName"
