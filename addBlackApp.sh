#!/bin/bash

export LC_COLLATE=C
for s
do
    if [ $(grep -e "^$s$" ./black_app.list | wc -l) -gt 0 ]
    then
        echo "Exists: $s"
    else
        echo "$s" >> ./black_app.list
        sort -u ./black_app.list -o ./black_app.list
        echo "Added: $s"
    fi
done

