<div align="center">

# Ginblog（项目已完成，欢迎使用)

</div>

<div align="center">

[![Go Report Card](https://goreportcard.com/badge/github.com/wejectchen/ginblog)](https://goreportcard.com/report/github.com/wejectchen/ginblog)
![](https://img.shields.io/badge/Go-Package-blue)</a>

</div>

<div align="center">
<img  src="https://s1.328888.xyz/2022/08/13/Tl95d.jpg" width="600" height="350"/>
</div>


## 介绍

gin+vue 全栈制作一个博客。

这是一个分享全栈制作过程的项目，旨在为有兴趣接触 golang web 开发的朋友分享一些制作经验。

你可以前往 [B 站(https://www.bilibili.com/video/BV1oe411u7Up/?vd_source=3ad57411ea01c62c71e224b64b1499fd)](https://www.bilibili.com/video/BV1oe411u7Up/?vd_source=3ad57411ea01c62c71e224b64b1499fd)
观看部署过程

## 项目预览

- www.qqoc.co

## 感谢
该项目基于 https://github.com/wejectchen/Ginblog 二次开发

采用面向对象的方式，重构了Model层。文章采用markdown储存

UI来自于 https://github.com/theme-nexmoe/hexo-theme-nexmoe

采用vue3+pinia+vite+ts重写

## 目录结构

```shell
├─  .gitignore
│  go.mod // 项目依赖
│  go.sum
│  LICENSE
│  main.go //主程序
│  README.md
│  tree.txt
│          
├─api         
├─config // 项目配置入口   
├─database  // 数据库备份文件（初始化）
├─log  // 项目日志
├─middleware  // 中间件
├─model // 数据模型层
├─routes
│      router.go // 路由入口    
├─static // 打包静态文件
│  ├─admin  // 后台管理页面 (已废弃，打包静态文件在web/admin/dist下)         
│  └─front  // 前端展示页面 (已废弃，打包静态文件在web/front/dist下) 
├─upload   
├─utils // 项目公用工具库
│  │  setting.go 
│  ├─errmsg   
│  └─validator         
└─web // 前端开发源码（VUECLI项目源文件)
    ├─admin             
    └─front
```

## 运行&&部署

1. 克隆项目
```shell
git clone https://github.com/sjtuli/Vue-Gin-Blog.git
```

2. 转到下面文件夹下

```shell
cd yourPath/Vue-Gin-Blog
```

3. 安装依赖

```
go mod tidy
```

4. 初始化项目配置config.ini

```ini
./config/config.ini

[server]
AppMode = debug # debug 开发模式，release 生产模式
HttpPort = :3000 # 项目端口
JwtKey = 89js82js72 #JWT密钥，随机字符串即可

[database]
Db = mysql #数据库类型，不能变更为其他形式
DbHost = 127.0.0.1 # 数据库地址
DbPort = 3306 # 数据库端口
DbUser = ginblog # 数据库用户名
DbPassWord = admin123 # 数据库用户密码
DbName = ginblog # 数据库名

[qiniu]
# 七牛储存信息
Zone = 1 # 1:华东;2:华北;3:华南,不填默认华北。境外服务器特殊使用环境自行配置
AccessKey =
SecretKey =
Bucket =
QiniuSever =

[tencentyun]
# 腾讯云COS储存信息
AccessKey =
SecretKey =
Bucket =
Sever =


[admin]
Username = admin
Password =admin
```

5. 前端部署
```shell
cd web/front
yarn
yarn build
```

6. 启动项目

```shell
 go run main.go
```

此时，项目启动，你可以访问页面

```shell
首页
http://localhost:8848
后台管理页面
http://localhost:8848/admin

默认管理员:admin  密码:admin
```

enjoy~~~~

#### ==使用、二开过程中，发现问题或者有功能需求欢迎提交 `Iusse` 或者直接 `PR`==

## 实现功能

- [x] 简单的用户管理权限设置
- [x] 用户密码加密存储
- [x]  文章分类自定义
- [x]  列表分页
- [x] 图片上传七牛云、腾讯云
- [x]  JWT 认证
- [x] 自定义日志功能
- [x]  跨域 cors 设置
-[x] 前端响应式适配移动端
- [x] 页面dark模式

## 技术栈

- golang
    - Gin web framework
    - gorm(v1 && v2)
    - jwt-go
    - scrypt
    - logrus
    - gin-contrib/cors
    - go-playground/validator/v10
    - go-ini
- JavaScript
    - vue3
    - vue cli
    - vue router
    - ant design vue
    - element UI 
    - vuetify
    - axios
    - tinymce
    - moment
    - pinia
- MySQL version:8.0.21

## TODO
- [ ] 微信小程序
- [ ]  文章评论功能
-[ ]   Docker支持