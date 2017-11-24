#!/bin/bash

echo 'package blackwhite' > _
echo 'var black_list = map[string]byte{' >> _
awk '{print "\"" $1 "\":0,"}' ../black.list >> _
echo '}' >> _
mv _ black_list.go
