# iam
统一资源调度系统

`参考项目地址:https://github.com/marmotedu/iam` 

## 目录结构



## 第一阶段
### 功能特性
- 对用户信息的基本操作
- 注册（需要邮箱绑定）、登录、修改密码、找回密码、修改用户其他信息

- 搭建一套web项目常用的系统体系，避免重复造轮子。
> 包含常用集合，如mysql、log、redis等组件的配置，命令行读取等。


### 快速部署
`前后端分离，所以要额外部署前端项目` 

`但前端还没写好`

后端部署
1. 普通部署

```shell
go mod tidy # 下载依赖

# 检查配置文件，启动mysql容器或服务

# 默认方式启动服务
go run cmd/apiserver/main.go 

# 二进制文件启动
# 内置了三个参数 port, mode, config 分别表示服务端口，服务模式（debug、release），服务配置文件位置
go build cmd/apiserver/main.go
./main  [options]

```

2. Docker 部署


## 关于作者
    YeeBing Chen yeebingchen@qq.com
## 许可证
IAM is licensed under the MIT. See [LICENSE](LICENSE) for the full license text.