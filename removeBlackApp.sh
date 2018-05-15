#!/bin/bash

for s
do
    if [ $(grep -P "^$s$" ./black_app.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$s$/d" ./black_app.list
        echo "Removed: $s"
    else
        echo "No: $s"
    fi
done
