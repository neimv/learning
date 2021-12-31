# !/bin/bash
# Programa para revisar la declaracion de variables

opcion=0
nombre=Neimv

echo "Opcion: $opcion y Nombre: $nombre"

# Exportar la variable nombre, para que este disponible a los demas procesos
export nombre

# llamar al siguiente script para recuperar la variable
./02_variables_2.sh
