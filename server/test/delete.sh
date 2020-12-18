#!/bin/bash
echo "测试删除接口"
read -p "请输入id:" id
curl -d "id=$id" http://nj-jay.com:8080/api/v2/book/delete
