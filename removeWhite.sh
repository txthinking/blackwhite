#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -e "^$s$" ./white.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$s$/d" ./white.list
        echo "Removed: $s"
    else
        echo "No: $s"
    fi
done

