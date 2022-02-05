#!/usr/bin/env bash

fileOutPut="hosts.txt"
counter=0

cp /etc/hosts.bk /etc/hosts && failed=0 || failed=1

echo $?
echo $failed

if [ $failed -eq 1 ]; then
    exit 1
fi

# restore=$(cp /etc/hosts /etc/hosts.bk)
echo "conts"

# docker ps -q | xargs -n 1 docker inspect --format "{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}} {{ .Name }}" | sed 's/ \// /' | grep os_* | sort > $fileOutPut

# # creating file inventory
# touch inventory
# echo "[servers]" | tee inventory

# while IFS= read -r line
# do
#     # echo "$counter"
#     # echo "$line"

#     echo "******************************************************"
#     arrayHosts=($line)
#     echo ${#arrayHosts[*]}
#     echo ${arrayHosts[1]}
#     echo -e ${arrayHosts[0]} "\t" ${arrayHosts[1]}
#     echo "******************************************************"

#     ((counter++))
# done < "$fileOutPut"

