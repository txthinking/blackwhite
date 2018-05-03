#!/bin/bash

curl -s https://raw.githubusercontent.com/17mon/china_ip_list/master/china_ip_list.txt -o ../china_cidr.list

echo 'package blackwhite' > _
echo 'var white_cidr []string = []string{' >> _
echo '"10.0.0.0/8",' >> _
echo '"100.64.0.0/10",' >> _
echo '"169.254.0.0/16",' >> _
echo '"172.16.0.0/12",' >> _
echo '"192.168.0.0/16",' >> _
awk '{print "\"" $1 "\","}' ../china_cidr.list >> _
echo '}' >> _
mv _ white_cidr.go
