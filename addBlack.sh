#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -P "^$s$" ./black.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./black.list
        sort -u ./black.list -o ./black.list
        echo "Added: $s"
    fi
done

