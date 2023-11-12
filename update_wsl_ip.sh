#!/bin/bash

# Retrieve the IP address of the WSL 2 instance
WSL_IP=$(ip addr show eth0 | grep 'inet ' | awk '{print $2}' | cut -d/ -f1)

# Set the IP address as an environment variable
export WSL_IP
echo "Running with WSL_IP=$WSL_IP"
# Start your Docker Compose stack
 docker-compose up
