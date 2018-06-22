#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -e "^$s$" ./black.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$s$/d" ./black.list
        echo "Removed: $s"
    else
        echo "No: $s"
    fi
done

