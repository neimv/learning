#!/usr/bin/env sh

######################################################################
# @author      : neimv (neimv@dark-universe)
# @file        : exec
# @created     : domingo feb 05, 2023 14:56:57 CST
#
# @description : 
######################################################################

tmux new-session -d 'vim'
tmux split-window -v 'ipython'
tmux split-window -h
tmux new-window 'mutt'
tmux -2 attach-session -d

