#!/bin/bash
echo 测试查找接口

read -p "请输入书籍类型或者书籍名称或者作者:" name

curl -d "name=$name" http://nj-jay.com:8080/api/v2/book/search
