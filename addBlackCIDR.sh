#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -P "^$s$" ./black_cidr.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./black_cidr.list
        sort -u ./black_cidr.list -o ./black_cidr.list
        echo "Added: $s"
    fi
done

