################################################
# System setup for Containment program
# Author: @christophersinclair

# Creates necessary systemd configuration directory
# Attaches eBPF programs to the kernel
# Deploys Containment software
# Configures systemd service
# Starts Containment service
################################################

#!/bin/bash

# Configurations
mkdir -p /etc/containment.d/
cp systemd/containment.d/* /etc/containment.d/

# Systemd service
cp systemd/containment.service /etc/systemd/system/

# Deploy Containment software
if [[ ! -f containment ]]; then
    go build -o containment
fi
cp containment /usr/local/bin/
chmod +x /usr/local/bin/containment

# Start containment service
sudo systemctl daemon-reload
sudo systemctl enable containment
sudo systemctl start containment