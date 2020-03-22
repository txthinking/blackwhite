#!/bin/bash

# !!! This require GUN tools

grep="grep"
if [ $(uname) = "Darwin" ]
then
    command -v gawk >/dev/null 2>&1
    if [ $? -eq 0 ]
    then
        grep="ggrep"
    else
        echo "Please install GUN grep"
        exit
    fi
fi

awk="awk"
if [ $(uname) = "Darwin" ]
then
    command -v gawk >/dev/null 2>&1
    if [ $? -eq 0 ]
    then
        awk="gawk"
    else
        echo "Please install GUN awk"
        exit
    fi
fi

sed="sed"
if [ $(uname) = "Darwin" ]
then
    command -v gsed >/dev/null 2>&1
    if [ $? -eq 0 ]
    then
        sed="gsed"
    else
        echo "Please install GUN sed"
        exit
    fi
fi

function gethost(){
    if [ $(echo $1 | $grep -P '\[' | wc -l) -ne 0 ]
    then
        return
    fi
    host=$(echo $1 | $awk -F : '{print $1}')
    if [ $(echo $host | $grep -P '[A-Za-z]' | wc -l) -eq 0 ]
    then
        return
    fi
    echo $host
}

function ischina(){
    ip=$(dig +short $1 @114.114.114.114 | tail -n 1)
    if [ $(echo $ip | $grep -P '\d' | wc -l) -eq 0 ]
    then
        return
    fi
    if [ $(mmdblookup --file ./GeoLite2-Country.mmdb --ip $ip country iso_code | $grep -P 'CN' | wc -l) -eq 1 ]
    then
        echo "yes"
        return
    fi
    if [ $(mmdblookup --file ./GeoLite2-Country.mmdb --ip $ip registered_country iso_code | $grep -P 'CN' | wc -l) -eq 1 ]
    then
        echo "yes"
        return
    fi
}

function getdomain(){
    if [ $(echo $1 | $awk -F . '{print $(NF)}') = "cn" ]
    then
        return
    fi
    dm=$(echo $1 | $awk -F . '{print $(NF-1) "." $(NF)}')
    if [ $(echo $dm | wc -L) -eq 5 ]
    then
        echo "please review $dm"
        echo $dm >> /tmp/review.5.list
        return
    fi
    if [ $(echo $dm | wc -L) -eq 6 ]
    then
        echo "please review $dm"
        echo $dm >> /tmp/review.6.list
        return
    fi
    echo $dm
}

for s in $(cat $1 | sort | uniq)
do
    host=$(gethost $s)
    if [ -z $host ]
    then
        continue
    fi
    is=$(ischina $host)
    if [ "$is" != "yes" ]
    then
        continue
    fi
    dm=$(getdomain $host)
    if [ -z $dm ]
    then
        continue
    fi
    ../../addWhite.sh $dm
done
