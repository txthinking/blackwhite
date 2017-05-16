#!/bin/bash

curl -s https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt -o ../china_cidr.list

echo 'package blackwhite' > _
echo 'var china_cidr []string = []string{' >> _
awk '{print "\"" $1 "\","}' ../china_cidr.list >> _
echo '}' >> _
mv _ china_cidr.go
