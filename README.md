# mysql-gin-api

## 简介

前后端分离项目,关于书籍管理的一个系统

构建一个云原生的基于mysql和go的微服务, 使用resful api与前端进行交互

该项目已经部署在linux服务器上 [地址](http://localhost/test)

## api文档

### 获取所有用户信息及登录和注册

* [GET]  http://localhost:8080/api/v2/login

  > 获取所有用户的登录信息
  >
  > 使用curl模拟请求,即可得到json数据
  >
  > `curl http://localhost:8080/api/v2/login`

* [POST] http://localhost:8080/api/v2/trueLogin

  > 两个参数: username password
  >
  > 通过前端接收到用户登录信息后,返回success or error
  >
  > 使用curl模拟请求
  >
  > `curl -d "username=linux&password=linux" http://localhost:8080/api/v2/trueLogin`
  >
  > 若用户名存在在数据库并且密码正确,则返回success

* [POST] http://localhost:8080/api/v2/login/add

  > 两个参数: username, password
  >
  > 通过前端接收到注册的用户信息, 返回"success add user" or "err add user"
  >
  > 使用curl模拟请求
  >
  > `curl -d "username=linux&password=linux" http://localhost:8080/api/v2/login/add`
  >
  > 若数据库中无相同的username,则返回"success add user"
  >
  > 否则返回"err add user"

### 数据信息的增删改查

* [GET] http://localhost:8080/api/v2/books/:page

  > 参数: page
  >
  > 通过前端的page值返回对应的数据,每页8个数据
  >
  > 使用curl模拟请求
  >
  > `curl http://localhost:8080/api/v2/books/1`
  >
  > page为1则返回第一页的数据
  >
  > 当page值到了一定值时,数据库查询不到相关的数据返回空[]

* [POST] http://localhost:8080/api/v2/book/search

  > 参数: name
  >
  > 进行模糊匹配,可以匹配的字段为书籍名称,书籍类型,书籍作者
  >
  > 若在数据库中查询到相关数据则返回该书籍的所有信息
  >
  > 使用curl模拟请求
  >
  > `curl -d "name=鲁迅" http://localhost:8080/api/v2/book/search`
  >
  > 若找到则返回所有鲁迅的书籍的信息
  >
  > 否则返回空[]

* [POST] http://localhost:8080/api/v2/book/add

  > 参数: name, types, author, price, addTime
  >
  > 通过前端提交的数据返回到后端,添加数据到数据库中
  >
  > 使用curl模拟请求
  >
  > `curl  -d "name=test&types=test&author=test&price=1000&addTime=0 http://localhost:8080/api/v2/book/add`

* [POST] http://localhost:8080/api/v2/book/delete

  > 参数: id
  >
  > 通过主键删除信息
  >
  > 使用curl模拟请求
  >
  > `curl -d "id=10" http://localhost:8080/api/v2/book/delete`

* [POST] http://localhost:8080/api/v2/book/update

  > 参数:name, types, author, price, addTime, id
  >
  > 通过id修改数据
  >
  > 使用curl模拟请求
  >
  > `curl -d "name=test&types=test&author=test&price=1000&addTime=0&id=1 http://localhost:8080/api/v2/book/update  ` 

## TODO

- [x] 路由拆分

- [x] jwt鉴权

- [x] 使用gorm

- [x]  wagger api文档