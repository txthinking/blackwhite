#!/bin/bash

for s in $(cat $1)
do
    ./addWhite.sh $s
done
