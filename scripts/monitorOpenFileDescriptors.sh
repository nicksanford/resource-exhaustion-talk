#!/bin/bash
if [ -z "$1" ]; then
    echo "Monitor number of open file descriptors a process is using"
    echo "usage $0 <pid>"
    exit 1
fi

watch -n 0.1 "sudo lsof -p $1 | wc -l"

