#!/bin/bash
if [ -z "$1" ]; then
    echo "Monitor number of threads a process is using"
    echo "usage $0 <pid>"
    exit 1
fi

watch -n 0.1 "ps -o nlwp $1"
