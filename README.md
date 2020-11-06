# mysql-gin-api

前后端分离

go + gin(web框架) + mysql(数据库) + restful api

## 代码结构

main.go 项目入口

conf    配置文件(mysql参数)

controllers 返回json数据

database    与数据库进行交互

middlewares 中间件

models  声明类

routers 设计路由

service 返回数据给controllers层

view    前端界面
