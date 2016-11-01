#!/bin/bash
#
# https://github.com/txthinking/pac
#

for dm
do
    if [ $(grep -P "^$dm$" ./black.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$dm$/d" ./black.list
        echo "Removed: $dm"
    else
        echo "No: $dm"
    fi
done

