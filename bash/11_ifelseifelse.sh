# !/bin/bash

edad=0

read -p "Indique cual es si edad: " edad

if [ $edad -le 18 ]; then
    echo "La persona es adolescente"
elif [ $edad -ge 19 ] && [ $edad -le 64 ]; then
    echo "la persona es adulta"
else
    echo "la persona es adulto mayor"
fi
