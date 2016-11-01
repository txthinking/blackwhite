#!/bin/bash
#
# https://github.com/txthinking/pac
#

for dm
do
    if [ $(grep -P "^$dm$" ./black.list | wc -l) -gt 0 ]
    then
        echo "Exists: $dm"
    else
        sort -u ./black.list -o ./black.list
        echo "Added: $dm"
    fi
done

