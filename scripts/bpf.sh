################################################
# eBPF setup for Containment program
# Author: @christophersinclair

################################################

#!/bin/bash
clang -O2 -target bpf -c ../ebpf/execvemon.c -o ../ebpf/execvemon.o

clang -O2 -g -Wall -I/usr/include -I/usr/include/bpf -o load ../ebpf/loader.c -lbpfsudo ./load
