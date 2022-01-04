# !/bin/bash

option=1
backupName=""
clave=""

echo "Programa Utilidades postgresq"
# Acepta el ingreso de informacion de solo caracter
read -n1 -p "Ingresa una opcion: " opcion
echo -e "\n"
read -n10 -p "Ingresa el nombre del archivo del backup: " backupName
echo -e "\n"
echo "Opcion: $opcion, backupName: $backupName"
read -s -p "Clave: " clave
echo "Clave: $clave"
