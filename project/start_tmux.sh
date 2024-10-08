#!/bin/bash

# Name of the tmux session
SESSION_NAME="dev_session"

# Start a new tmux session
tmux new-session -d -s $SESSION_NAME

# Create a window for the editor
tmux rename-window -t $SESSION_NAME:0 'Project'
tmux send-keys -t $SESSION_NAME:0 'ls' C-m

# Create a new window for the server
tmux new-window -t $SESSION_NAME:1 -n 'dashboard'
tmux send-keys -t $SESSION_NAME:1 'minikube dashboard' C-m

# Create a new window for logs
tmux new-window -t $SESSION_NAME:2 -n 'Helpers'
tmux send-keys -t $SESSION_NAME:2 'cd ../../microservices-helper/' C-m

# Create a new window for database
tmux new-window -t $SESSION_NAME:3 -n 'Proto'
tmux send-keys -t $SESSION_NAME:3 'cd ../../microservices-proto/' C-m
# Create a new window for database
tmux new-window -t $SESSION_NAME:4 -n 'Port FWD'
tmux send-keys -t $SESSION_NAME:4 'kubectl port-forward service/broker 8082:8082' C-m


# Attach to the session
tmux attach -t $SESSION_NAME