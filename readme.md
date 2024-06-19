**校园抖音开发记录**

**开发模式：**mvc

![img](C:\Users\h2629\blog\source\images\clip_image002.jpg)

**目录结构：**

- app

- - controllers // 控制器

  - - auth_controller.go // 处理用户登录/注册相关逻辑
    - user_controller.go // 处理获取用户信息相关逻辑
    - video_controller.go // 处理视频发布/获取视频相关逻辑
    - like_controller.go // 处理视频点赞相关逻辑
    - comment_controller.go // 处理视频评论相关逻辑
    - follow_controller.go // 处理关注/获取关注列表/粉丝列表/好友列表相关逻辑
    - message_controller.go // 处理发送消息/获取聊天记录相关逻辑

  - models // 模型

  - - user.go // 定义用户模型，包括注册、登录、获取用户信息等操作
    - video.go // 定义视频模型，包括发布、获取视频等操作
    - like.go // 定义点赞模型，包括点赞、取消点赞等操作
    - comment.go // 定义评论模型，包括发布、获取评论等操作
    - follow.go // 定义关注模型，包括关注、取消关注等操作
    - message.go // 定义消息模型，包括发送消息、获取聊天记录等操作

  - services // 服务

  - - auth_service.go // 处理用户注册、登录逻辑
    - user_service.go // 处理获取用户信息逻辑
    - video_service.go // 处理视频发布、获取视频逻辑
    - like_service.go // 处理视频点赞逻辑
    - comment_service.go // 处理视频评论逻辑
    - follow_service.go // 处理关注/获取关注列表/粉丝列表/好友列表相关逻辑
    - message_service.go // 处理发送消息/获取聊天记录相关逻辑

- database // 数据库

- - mysql // 数据库初始化以及配置
  - redis // 缓存初始化以及配置
  - viper // viper初始化以及配置

- public // 公共资源 存放视频文件,以及视频封面

- routes // 路由

- - api.go // 定义 API 路由

- utils // 工具

- - ffmpeg // 视频处理文件
  - response.go // 定义 API 响应结构体和函数
  - jwt.go // token处理
  - videos.go // 视频处理的相应方法

- config.yml // 配置文件

- go.mod // 工作区以及依赖

- - go.sum

- mian.go // 入口文件

**项目依赖：**

Gin框架: go get github.com/gin-gonic/gin
 viper: go get -u github.com/spf13/viper
 gorm: go get -u gorm.io/gorm
 mysql: go get -u gorm.io/driver/mysql
 jwt: go get github.com/golang-jwt/jwt/v4
 redis: go get -u github.com/go-redis/redis/v8

**建立数据库模型：**

![img](C:\Users\h2629\blog\source\images\clip_image004.jpg)

**运行说明**

此源码为服务端源码，且已经在线上运行，如果需要本地运行请在yml文件里更改对应数据库地址，redis地址，以及本地IPv4的地址，然后可以在main文件下直接编译运行，也可以在终端里使用 go build mian.go 编译运行，然后去前端修改对应的BaseUrl: