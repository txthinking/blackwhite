#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -P "^$s$" ./white_cidr.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$s$/d" ./white_cidr.list
        echo "Removed: $s"
    else
        echo "No: $s"
    fi
done

