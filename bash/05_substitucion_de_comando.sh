# !/bin/bash
# Como ejecutar comandos dentro de un programa y almacenarlas dentro de una varible

ubicacionActual=`pwd`
infoKernel=$(uname -a)

echo "la ubicacion actual es la siguiente: $ubicacionActual"
echo "la version de kernel es: $infoKernel"
