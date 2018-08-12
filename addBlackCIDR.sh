#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -e "^$s$" ./black_cidr.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./black_cidr.list
        sort -u -t. -k 1,1n -k 2,2n -k 3,3n -k 4,4n ./black_cidr.list -o ./black_cidr.list
        echo "Added: $s"
    fi
done

