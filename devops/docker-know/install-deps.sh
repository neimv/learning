#!/usr/bin/bash

apt update && apt install openssh-server

ufw allow ssh
systemctl enable ssh

