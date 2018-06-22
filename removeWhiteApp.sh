#!/bin/bash

for s
do
    if [ $(grep -e "^$s$" ./white_app.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$s$/d" ./white_app.list
        echo "Removed: $s"
    else
        echo "No: $s"
    fi
done
