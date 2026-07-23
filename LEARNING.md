# Bluebell Rebuild - 学习记录

> 本文件记录整个复现过程。
> 每完成一个阶段就更新一次。

---

# 第一阶段：最小 Gin 服务

目标：

理解 Gin 如何启动。

完成：

* 创建 main.go
* 创建 router
* 注册 GET /ping
* 返回 JSON

掌握知识：

* gin.Default()
* r.GET()
* c.JSON()
* r.Run()

状态：

✅ 已完成

---

# 第二阶段：Controller 分层

目标：

理解为什么项目需要分层。

完成：

* 创建 controller
* 编写 PingHandler
* router 调用 controller

掌握知识：

router：

负责：

"收到请求以后交给谁。"

controller：

负责：

"真正处理请求。"

状态：

✅ 已完成

---

# 第三阶段：统一响应

目标：

以后所有接口统一返回格式。

完成：

创建：

ResponseData

ResponseSuccess()

掌握知识：

统一返回：

```
{
    code,
    msg,
    data
}
```

为什么要统一：

方便前端统一处理。

状态：

✅ 已完成

---

# 第四阶段：配置读取

目标：

把程序配置放进 YAML。

完成：

* config.yaml
* setting.Init()
* Viper
* main 读取配置启动

掌握知识：

程序启动流程：

```
main

↓

Init()

↓

读取 config.yaml

↓

获得配置

↓

启动服务器
```

状态：

✅ 已完成

---

# 当前项目进度

████░░░░░░░░░░░░░

约 20%

---

# 下一阶段

准备进入：

日志系统（logger）

随后开始：

用户注册

↓

用户登录

↓

JWT

↓

中间件

↓

MySQL

↓

帖子

↓

社区

---

# 当前已经理解的知识

✓ Gin 基本使用

✓ 路由

✓ Controller

✓ 统一响应

✓ Viper

✓ YAML 配置

---

# 当前容易忘记的地方

1.

router 不处理业务。

2.

controller 不访问数据库。

3.

logic 才是真正的业务层。

4.

dao 只负责数据。

5.

main 只负责启动。

---

# 每完成一节课后

请更新：

* 已完成内容
* 当前项目进度
* 下一阶段
* 新学到的重要思想
