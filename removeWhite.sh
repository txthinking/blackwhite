#!/bin/bash
#
# https://github.com/txthinking/pac
#

for dm
do
    if [ $(grep -P "^$dm$" ./white.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$dm$/d" ./white.list
        echo "Removed: $dm"
    else
        echo "No: $dm"
    fi
done

