#!/bin/bash

for pkg
do
    if [ $(grep -P "^$pkg$" ./white_apps.list | wc -l) -gt 0 ]
    then
        echo "Exists: $pkg"
    else
        echo "$pkg" >> ./white_apps.list
        sort -u ./white_apps.list -o ./white_apps.list
        echo "Added: $pkg"
    fi
done

