#!/bin/bash

export LC_COLLATE=C
for s
do
    if [ $(grep -e "^$s$" ../white.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ../white.list
        sort -u ../white.list -o ../white.list
        echo "Added: $s"
    fi
done

