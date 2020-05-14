#!/bin/bash

# check from a address list, domain:port
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
    if [ $(echo $1 | $grep -P '(goog|gstatic|googleusercontent|googleapis|googlevideo|gvt)' | wc -l) -gt 0 ]
    then
        echo "no"
        return
    fi
    ip=$(dig +short $1 @223.6.6.6 | tail -n 1)
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
    echo "no"
}

function getdomain(){
    if [ $(echo $1 | $awk -F . '{print $(NF)}') = "cn" ]
    then
        return
    fi
    dm=$(echo $1 | $awk -F . '{print $(NF-1) "." $(NF)}')
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
    if [ "$is" = "" ]
    then
        continue
    fi
    dm=$(getdomain $host)
    if [ -z $dm ]
    then
        continue
    fi
    if [ "$is" = "yes" ]
    then
        echo $dm > /tmp/china.list
    fi
    if [ "$is" = "no" ]
    then
        if [ $(echo $dm | $grep -P '^(com\.|ip\.sb)$' | wc -l) -ne 0 ]
        then
            continue
        fi
        echo $dm > /tmp/nonchina.list
    fi
done
