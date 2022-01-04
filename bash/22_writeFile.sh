# !/bin/bash

echo "Escribir en un archivo"

echo "VAlores escritos con el comando echo" >> $1

# Adicion multilinea
cat <<EOM >> $1
$2
EOM
