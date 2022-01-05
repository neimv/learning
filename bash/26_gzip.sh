# !/bin/bash

echo "empaquetar todos los scripts de la carpeta course"

tar -cvf shellCourse.tar *.sh
# cuando se empaqueta con gzip el empaquetamiento anterior se elimina
gzip shellCourse.tar

echo "empaquetar un solo archivo, con ratio 9"
gzip -9 09_options.sh
