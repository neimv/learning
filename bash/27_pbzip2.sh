# !/bin/bash

echo "empaquetar todos los scripts de la carpeta course"

tar -cvf shellCourse.tar *.sh
pbzip2 -f shellCourse.tar

echo "empaquetar un directorio con tar y pbzip2"
tar -cf *.sh -c > shellCourseDos.tar.bz2
