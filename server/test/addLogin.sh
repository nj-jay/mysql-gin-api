#!/bin/bash
echo 测试注册

read -p "请输入username:" username
read -p "请输入password:" password

echo 正在测试注册接口

curl -d "username=$username&password=$password" http://nj-jay.com:8080/api/v2/login/add

