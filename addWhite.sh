#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -P "^$s$" ./white.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./white.list
        sort -u ./white.list -o ./white.list
        echo "Added: $s"
    fi
done

