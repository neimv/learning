# !/bin/bash

echo "Empeaquetar todos los scripts y transferirlo por la red a otro equipo utilizando el comando rsync"

read -p "Ingresa el host: " host
read -p "Ingresar el usuario: " usuario
echo -e "\nEn este momento se procedera a empaquetar ela carpeta y transferir segun lso datos ingresados"

rsync -avz $(pwd) $usuario@$host:/home/neimv/Documents
