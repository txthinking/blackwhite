#!/bin/bash

for s
do
    if [ $(grep -e "^$s$" ./white_app.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./white_app.list
        sort -u ./white_app.list -o ./white_app.list
        echo "Added: $s"
    fi
done

