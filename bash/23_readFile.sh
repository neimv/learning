# !/bin/bash

echo "Leer en un archivo"
cat $1
echo -e "\nAlmancenar los valores en una variable"
valorCat=`cat $1`
echo "$valorCat"

# Se utiliza la variable IFS (Internal Field Separator) para evitar que los espacios en blanco se recorten al incioio o al final
echo -e "\nLeer archivos linea por linea"

while IFS= read linea
do
    echo "$linea"
done < $1
