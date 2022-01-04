# !/bin/bash

echo "Loops anidados"

for fil in $(ls); do
    for nombre in {1..4}; do
        if [ $fil = "10_download.sh" ]; then
            break;
        elif [[ $fil == 04* ]]; then
            continue;
        else
            echo "Nombre archivo: $fil _ $nombre"
        fi
    done
done
