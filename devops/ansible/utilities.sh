#!/usr/bin/env bash

fileOutPut="hosts.txt"
counter=0

##########################################################################
# Functions
##########################################################################
handle_errs() {
    exit_need=$1
    failed=$2

    if [ "$exit_need" = "true" ] && [ $failed -eq 1 ]; then
        echo "Command failed and exit of script"
        exit 1
    elif [ "$exit_need" = "false" ] && [ $failed -eq 1 ]; then
        echo "Command failed and continues with the script"
    fi
}

read -s -p "Ingress password of sudo: " password
echo -e "\n"

# Restore original file
sudo cp /etc/hosts.bk /etc/hosts && failed=0 || failed=1
handle_errs "false" $failed

sudo cp /etc/hosts /etc/hosts.bk && failed=0 || failed=1
handle_errs "true" $failed

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

