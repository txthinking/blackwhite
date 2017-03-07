#!/bin/bash

for pkg
do
    if [ $(grep -P "^$pkg$" ./white_apps.list | wc -l) -gt 0 ]
    then
        sed -r -i "/^$pkg$/d" ./white_apps.list
        echo "Removed: $pkg"
    else
        echo "No: $pkg"
    fi
done

