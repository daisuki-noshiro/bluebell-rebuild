# Bluebell Rebuild - 项目地图

> 本文件用于记录整个项目的结构和各模块职责。
> 当长时间没有继续学习时，只需要阅读本文件即可快速恢复项目上下文。

---

# 当前项目结构（进行中）

```
bluebell-rebuild
│
├── main.go
│      程序入口
│
├── conf/
│      配置文件(config.yaml)
│
├── setting/
│      读取配置(Viper)
│
├── router/
│      注册所有路由
│
├── controller/
│      接收HTTP请求
│      参数解析
│      返回统一响应
│
├── logic/
│      业务逻辑（待实现）
│
├── dao/
│      数据访问层（待实现）
│
│   ├── mysql/
│   └── redis/
│
├── middlewares/
│      JWT、中间件（待实现）
│
├── models/
│      数据结构（待实现）
│
├── logger/
│      Zap日志（待实现）
│
└── pkg/
       公共工具（Snowflake、JWT等）
```

---

# 当前调用流程

目前已经完成：

```
main
    ↓
setting.Init()
    ↓
router.SetupRouter()
    ↓
GET /ping
    ↓
controller.PingHandler()
    ↓
ResponseSuccess()
    ↓
JSON 返回
```

---

# 各层职责

## main

程序入口。

负责：

* 初始化配置
* 初始化各种组件
* 启动 Gin

原则：

> main 不写业务。

---

## router

负责：

注册所有 HTTP 路由。

例如：

```
GET /ping

↓

controller.PingHandler
```

原则：

> router 只负责"把请求交给谁"。

---

## controller

负责：

* 接收请求
* 参数绑定
* 参数校验
* 调用 logic
* 返回统一响应

原则：

> controller 不直接访问数据库。

---

## logic（待实现）

负责：

真正的业务逻辑。

例如：

登录：

```
controller

↓

logic.Login

↓

dao.Login
```

以后：

* 登录
* 发帖
* 创建社区

都写在这里。

---

## dao（待实现）

负责：

数据库操作。

以后包括：

* MySQL
* Redis

原则：

> dao 不写业务，只负责数据。

---

# 当前完成情况

✓ Gin 服务启动

✓ router

✓ controller

✓ PingHandler

✓ ResponseSuccess

✓ Viper 配置读取

□ logger

□ MySQL

□ Redis

□ JWT

□ middleware

□ 用户模块

□ 社区模块

□ 帖子模块

---

# 最终目标

完整复现 Bluebell 项目。

不是照抄代码，而是理解：

* 为什么这样分层
* 每层负责什么
* 请求如何流转
* 为什么要这样设计
