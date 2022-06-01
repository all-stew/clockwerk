# clockwerk

## 目录结构

```bash
.
├── config 配置文件对应的结构体定义
├── controller 业务层
├── dao 操作数据库，给controller提供数据
├── entitys 实体类
│   ├── models 数据库对象
│   └── views 试图对象
├── globals 全局便利
├── initialize 服务初始化
├── logs 日志存储
├── middlewares 中间件
├── util 工具
│   └── response 封装response
├── router 路由
├── main.go 服务启动文件
└── settings-dev.yaml 配置文件
```