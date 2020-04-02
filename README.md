# inn
使用go-micro实践微服务之简单IM系统

> *项目开发中*

#### 快速开始

- 首先先安装`docker`和`docker-compose`
- 运行 `make docker-run`
- 访问 `http://localhost:8888`

#### 服务

- user: 提供用户登录注册的服务
- gateway: 提供消息收发的出入口，主要有四块功能：连接保持、协议解析、Session 维护和消息推送
- message: 提供消息业务逻辑处理服务。

#### 链路跟踪

基于 OpenTracing + Jaeger 构建分布式服务追踪系统

jaeger Ui: `http://localhost:16686`