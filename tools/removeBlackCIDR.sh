#!/bin/bash

for s
do
    if [ $(grep -e "^$s$" ../black_cidr.list | wc -l) -gt 0 ]
    then
        sed -r -i "\|^$s$|d" ../black_cidr.list
        echo "Removed: $s"
    else
        echo "No: $s"
    fi
done

