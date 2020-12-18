#!/bin/bash

echo "测试修改数据接口"
read -p "请输入id:" id
read -p "请输入书籍名称:" name
read -p "请输入书籍类型:" types
read -p "请输入作者:" author 
read -p "请输入价格:" price
read -p "请输入上架时间:" addTime

curl -d "name=$name&types=$types&author=$author&price=$price&addTime=$addTime&id=$id" http://nj-jay.com:8080/api/v2/book/update
