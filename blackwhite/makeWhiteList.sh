#!/bin/bash

echo 'package blackwhite' > _
echo 'var white_list = map[string]byte{' >> _
awk '{print "\"" $1 "\":0,"}' ../white.list >> _
echo '}' >> _
mv _ white_list.go
