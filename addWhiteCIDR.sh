#!/bin/bash
#
# https://github.com/txthinking/pac
#

for s
do
    if [ $(grep -e "^$s$" ./white_cidr.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./white_cidr.list
        sort -u ./white_cidr.list -o ./white_cidr.list
        echo "Added: $s"
    fi
done

