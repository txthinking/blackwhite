#!/bin/bash

for s in $(cat $1)
do
    ./addBlack.sh $s
done
