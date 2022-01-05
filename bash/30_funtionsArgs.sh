# !/bin/bash

opcion=0
fechaActual=`date +%Y%m%d`

# Funcion para instalar postgresql
instalar_postgres() {
    echo "Instalar postgres"
    echo "Verificar instalacion de postgres"

    verifyInstall=$(which psql)

    if [ $? -eq 0 ]; then
        echo -e "\n Postgreq ya esta instalado en el equipo"
    else
        read -s -p "ingresar contrase;a de sudo: " password
        read -s -p "ingresar contrase;a a utilizar en postgres: " passwordPostgres

        echo "$password" | sudo -S apt update
        echo "$password" | sudo -S apt -y postgresql postgresql-contrib
        sudo -u postgres psql -c "ALTER USER postgres WITH PASSWORD '{$passwordPost}';"
        echo "$password" | sudo -S systemctl enable postgresql.service
        echo "$password" | sudo -S systemctl start postgresql.service
    fi

    read -n 1 -s -r -p "Presione [ENTER] para continuar..."
}

# Funcion para desinstalar postgresq
desinstalar_postgres() {
    echo "Desinstalar postgres"

    read -s -p "ingresar contrase;a de sudo: " password
    echo -e "\n"
    echo "$password" | sudo -S systemctl stop postgresql.service
    echo "$password" | sudo -S apt -y --purge postgresql\*
    echo "$password" | sudo -S rm -r /etc/postgresql
    echo "$password" | sudo -S rm -r /etc/postgresql-common
    echo "$password" | sudo -S rm -r /var/lib/postgresql
    echo "$password" | sudo -S userdel -r postgres
    echo "$password" | sudo -S groupdel postgres

    read -n 1 -s -r -p "Presione [ENTER] para continuar..."
}

# Funcion para obtener respaldo
sacar_respaldo() {
    echo "Sacar respaldo"
    echo "Directorio respaldo" $1

    echo "Listar bases de datos"
    sudo -u postgres psql -c "\l"
    read -p "Elegir la base de datos a respaldar: " bddRespaldo
    echo -e "\n"

    if [ -d "$1" ]; then
        echo "Establecer permisos al directorio"
        echo "$password" | sudo -S chmod 755 $1
        echo "Obteniendo respaldo"
        sudo -u postgres pg_dump -Fc $bddRespaldo > "$1/bddRespaldo$fechaActual.back"
        echo "Respaldo realizado correctamente en la ubicacion: $1/bddRespaldo$fechaActual.back"
    else
        echo "el directorio $1 no existe"
    fi

    read -n 1 -s -r -p "Presione [ENTER] para continuar..."
}

# Funcion para restaurar respaldo
restaurar_respaldo() {
    echo "Restaurar respaldo"
    echo "Directorio de respaldo" $1

    echo "Listar respaldos"
    read -p "Ingresar el directorio donde estan los respaldos: " directorioBackup
    ls -la $directorioBackup
    read -p "Elegir el respaldo a restaurar" respaldoRestaurar
    echo -e "\n"
    read -p "Ingrese el nombre de la base de datos destino: " bddDestino

    # Verificar si la bdd existe
    verifyBdd=$(sudo -u postgres psql -lqt | cut -d \| -f 1 | grep -wq $bddDestino)

    if [ $? -eq 0 ]; then
        echo "Restaurando en la bdd destino: $bddDestino"
    else
        sudo -u postgres psql -c "create database $bddDestino"
    fi

    if [ -f "$respaldoRestaurar" ]; then
        echo "Restaurando respaldo"
        su -u postgres pg_restore -Fc -d $bddDestino "$directorioBackup/$respaldoRestaurar"
        echo "listar la base de datos"
        sudo -u postgres psql -c "\l"
    else
        echo "el respaldo $respaldoRestaurar no existe"
    fi

    read -n 1 -s -r -p "Presione [ENTER] para continuar..."
}

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
        1) instalar_postgres
            sleep 3
            ;;
        2) desinstalar_postgres
            sleep 3
            ;;
        3) 
            read -p "Directorio backup: " directorioBackup
            sacar_respaldo $directorioBackup
            sleep 3
            ;;
        4) 
            read -p "Directorio backup de respaldos: " directorioRespaldos
            restaurar_respaldo $directorioRespaldos
            sleep 3
            ;;
        5) echo "Salir del programa"
            exit 0
            ;;
    esac
done
