# mysql-gin-api

## 简介

前后端分离, 关于图书管理的一个系统, 包括权限管理,  对书籍的增删改查接口, 以及登录和注册接口

* 注册
* 登录接口拿到token
* 通过token调用增删改查接口

仅作为学习使用

## 主要技术栈

* gin(go web框架)
* mysql(数据库)
* jwt(json web token权限认证)
* swagger(api文档)

## api文档

http://localhost:8080:/swagger/index.html

注:

* 本文档未写完整
* token有效时间为2小时

## api接口

| 接口                                      | 功能               | 参数                                                   | type |
| ----------------------------------------- | ------------------ | ------------------------------------------------------ | ---- |
| http://localhost:8080/api/v2/login/add    | 注册               | username, password                                     | POST |
| http://localhost:8080/api/v2/trueLogin    | 登录,并生成token   | username, password                                     | POST |
| http://localhost:8080/api/v2/login        | 获取所有的用户信息 | Header中传入token【Authorization: Bearer token】       | GET  |
| http://localhost:8080/api/v2/books/:page  | 查询第page页       | page, Header中传入token【Authorization: Bearer token】 | GET  |
| http://localhost:8080/api/v2/book/search  | 查询指定的书籍     | name                                                   | POST |
| http://localhost:8080/api/v2/book/add     | 增加书籍           | name, types, author, price, addTime                    | POST |
| http://localhost:8080/api/v2/books/update | 更新书籍           | id, name, types, author, price, addTime                | POST |
| http://localhost:8080/api/v2/books/delete | 删除书籍           | id                                                     | POST |

## api使用

```shell
# 注册
curl -d "username=linux&password=linux" http://localhost:8080/api/v2/login/add
# 登录,若账号与密码均正确则返回生成的token,有效期为2小时
curl -d "username=linux&password=linux" http://localhost:8080/api/v2/trueLogin
# 返回的结果为
{
    "code": 200,
    "data": {
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxpbnV4IiwicGFzc3dvcmQiOiIkMmEkMDQkTWovMi9mOWFGOVJaaDA3SnVmeUx2T1p6Vk8vRURmSnNjOXhockVhLzlibEJqNkZaWTZxU3EiLCJleHAiOjE2MDgxNDY4NjUsImlzcyI6Im5qLWpheSJ9.lBeo1NWJH4flS7_ZsGKFjhtFYfeAdIuv6TMwFvpmNAc"
    },
    "msg": "success"
}

# 在header使用该token去调用其他接口,否则无法进行调用
# 查询第一页的书籍
curl  http://localhost:8080/api/v2/books/1 -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxpbnV4IiwicGFzc3dvcmQiOiIkMmEkMDQkTWovMi9mOWFGOVJaaDA3SnVmeUx2T1p6Vk8vRURmSnNjOXhockVhLzlibEJqNkZaWTZxU3EiLCJleHAiOjE2MDgxNDY4NjUsImlzcyI6Im5qLWpheSJ9.lBeo1NWJH4flS7_ZsGKFjhtFYfeAdIuv6TMwFvpmNAc"

# 查询name=python的书籍
curl -d "name=python" http://localhost:8080/api/v2/book/search -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxpbnV4IiwicGFzc3dvcmQiOiIkMmEkMDQkTWovMi9mOWFGOVJaaDA3SnVmeUx2T1p6Vk8vRURmSnNjOXhockVhLzlibEJqNkZaWTZxU3EiLCJleHAiOjE2MDgxNDY4NjUsImlzcyI6Im5qLWpheSJ9.lBeo1NWJH4flS7_ZsGKFjhtFYfeAdIuv6TMwFvpmNAc"
```

## TODO

- [x] 使用gorm
- [x] docker部署
