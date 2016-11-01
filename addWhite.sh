#!/bin/bash
#
# https://github.com/txthinking/pac
#

for dm
do
    if [ $(grep -P "^$dm$" ./white.list | wc -l) -gt 0 ]
    then
        echo "Exists: $dm"
    else
        echo "$dm" >> ./white.list
        sort -u ./white.list -o ./white.list
        echo "Added: $dm"
    fi
done

