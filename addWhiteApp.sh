#!/bin/bash

for pkg
do
    if [ $(grep -P "^$pkg$" ./white_app.list | wc -l) -gt 0 ]
    then
        echo "Exists: $pkg"
    else
        echo "$pkg" >> ./white_app.list
        sort -u ./white_app.list -o ./white_app.list
        echo "Added: $pkg"
    fi
done

