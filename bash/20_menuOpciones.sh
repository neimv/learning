# !/bin/bash

opcion=0

while :
do
    # Limpiar pantalla
    clear
    #Desplegar el menu de opciones
    echo "_______________________"
    echo "PGUTIL - Programa de Utilidad para Postgres"
    echo "_______________________"
    echo "MENU PRINCIPAL"
    echo "_______________________"
    echo "1. instalar postgres"
    echo "2. desinstalar postgres"
    echo "3. Sacar respaldo"
    echo "4. restar respaldo"
    echo "5. Salir"

    # Leer los datos del usuario
    read -n1 -p "Ingrese una opcion [1-5]: " opcion

    # validar la opcion ingresada
    echo -e "\n"
    case $opcion in
        1) echo "Instalar postgres"
            sleep 3
            ;;
        2) echo "desinstalar postgres"
            sleep 3
            ;;
        3) echo "Sacar respaldo"
            sleep 3
            ;;
        4) echo "restaurar respaldo"
            sleep 3
            ;;
        5) echo "Salir del programa"
            exit 0
            ;;
    esac
done
