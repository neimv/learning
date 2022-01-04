# !/bin/bash

option=0
backupName=""

echo "Programa Utilidades postgresql"
read -p "Ingresar una opcion: " option
read -p "Ingresar el nombre del archivo del backup: " backupName

echo "Opcion: $option, backup: $backupName"
