#!/bin/bash
for ((i=1; i <= 25; i++))

do
    echo -e "第$i页内容\n"
    curl http://nj-jay.com:8080/api/v2/books/$i
    echo -e "\n"

done




