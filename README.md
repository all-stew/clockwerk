# clockwerk

## 目录结构

```bash
.
├── Dockerfile
├── LICENSE
├── Makefile
├── README.md
├── go.mod
├── go.sum
├── logs
├── main.go
├── pkg
└── app
    ├── api 业务层
    ├── common
    ├── config 配置文件对应的结构体定义
    ├── external_api
    ├── global 全局变量
    ├── initialize 服务初始化
    ├── middlewares 中间件
    ├── models 数据库对象
    ├── repository 操作数据库，提供数据
    ├── routes 路由
    ├── task
    └── views 视图对象
```