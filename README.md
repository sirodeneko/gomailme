# gomailme

gomailme: A Simple mail Golang Web Service

gomailme 实现简单的api,直接部署后即可通过简单的url发送邮件

## 目的

本项目采用了Singo框架开发web服务，可以以本项目为基础快速拓展自己的邮件需求

## 特色

locales内放置翻译相关的配置文件

## Godotenv

项目在启动的时候依赖以下环境变量，但是在也可以在项目根目录创建.env文件设置环境变量便于使用(建议开发环境使用)

```shell
MYSQL_DSN="db_user:db_password@/db_name?charset=utf8&parseTime=True&loc=Local" # Mysql连接地址

#以下为singo自带参数可不配置
REDIS_ADDR="127.0.0.1:6379" # Redis端口和地址
REDIS_PW="" # Redis连接密码
REDIS_DB="" # Redis库从0到10
SESSION_SECRET="setOnProducation" # Seesion密钥，必须设置而且不要泄露
GIN_MODE="debug"
```

## Go Mod

本项目使用[Go Mod](https://github.com/golang/go/wiki/Modules)管理依赖。

```shell
go mod init go-crud
export GOPROXY=http://mirrors.aliyun.com/goproxy/
go run main.go // 自动安装
```

## 运行

```shell
go run main.go
```

项目运行后启动在3000端口（可以修改，参考gin文档)

##使用方法
本项目依赖mysql,请确保有mysql数据库正确配置.env(.env由.env.example复制后重命名)。
设置账号，授权码（也可直接编辑数据库相应字段）
(本地运行时情况,在服务器上请改变域名)
（post) http://localhost:3000/mail/set 
提供
user（账号） 
pass(授权码，不是密码，出于邮箱安全的考虑，很多邮箱缺省是关闭 POP3/SMTP 服务的，需要登录邮箱设置后开启。以 QQ 邮箱为例，进入邮箱“设置”，在“帐户”项里就可找到“POP3/SMTP服务”的设置项，进行开启。 可获得授权码)
可设置账号密码

（post) http://localhost:3000/mail/send 
提供
to(收件人，多个收件人用 "," 连接)
body(邮件内容)
可发送邮件

（get) http://localhost:3000/mail/send/:to/:body
直接通过url访问
提供 to , body 例 http://localhost:3000/mail/send/xxx@qq.com/大笨蛋 可发送大笨蛋到xxx@qq,com

本项目基于qq邮箱，如果是其他邮箱，请在 service\send_service.go修改41，56行（日后可能因为代码电话改变行数）

