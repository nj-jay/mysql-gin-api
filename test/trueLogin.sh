#!/bin/bash
echo "正在测试用户登录接口,按提示进行操作"
read -p "输入username:" username
read -p "输入password:" password
curl -d "username=$username&password=$password" http://nj-jay.com:8080/api/v2/trueLogin

