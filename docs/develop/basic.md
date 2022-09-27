# 开发历程

目标：从根基入手，搭建一套通用的web开发框架，避免以后重复造轮子。

## 基础 web 框架设计
采用 controller service store 分层思想，部分目录参考了 (project-layout)[https://github.com/golang-standards/project-layout]

## 项目基础设施(代码开发)

golangci-lint 配置, 参考和修改了他(官网)[golangci-lint 配置, 参考和修改了他)上的配置]上的配置 

git action 待加入（还没怎么用过）

## 项目组件设计(业务相关)

- 这些组件本身以前写代码做过了一点，重新封装了一下，不重复造轮子了，如zap包的初始化
- 参考了iam对于组件的设计与封装，但是他的很多设计过于庞大（如修改pkg/errors包，增加code字段，或者根据zap重构了一个日志系统）
我简单参考了他的设计，最小化完成封装
- 剩下的是参考网上现有的设计，如 gin middleware， 这些可以直接使用现成封装好的库，也可以根据自己需要修改
我的做法就是简单封装一层，对外暴露一些简单的参数设计，对于其他参数我就内部默认设计了，简化使用者的操作。如 gorm-mysql，go-redis等等

### 错误码设计
> 针对错误码，封装了相应的错误包，并写了相应的基准测试，保证错误包的正确性
>
> 重新封装了 Wrap、Unwrap、Is 函数（Is 函数有点问题）。

### zap 日志打印封装
支持 error 和 info 日志分开存储，日志文件切分，debug模式同时以console color形式打印到stdout中

### gorm-mysql redis 外部依赖初始化


### 添加命令行配置
对cobra、viper、pflag 进行封装、使用（参考iam的设计）
可以复用其他组件的option和valid，减少外部依赖。

添加命令行程序接口 命令行参数解析 配置文件解析 (封装使用 cobra pflag viper 三剑客)
> 命令行程序：高级功能 help查看命令， 命令行自动补全，查看命令行参数
> 
> 命令行参数解析
> 
> 配置文件解析：解析不同格式配置文件

### gin 中间件设计 (web服务基础中间件)
> request id 用于打印日志时可以区分同一条请求，如果设置成用户的信息，就可以获取到用户最近的操作
> 
> jwt 用于登录的鉴权，优点:去中心化
> 
> 重新包装了recovery 和 logger 中间件，支持写入到文件，同时有回调函数
> 
> cors 用于解决前后端跨域问题
> 
> limit 使用令牌桶对服务器流量进行限制， 防止服务器过载

## 功能实现
### 用户注册

### 用户登录

用户鉴权
>初步考虑使用jwt来实现，但是jwt对于退出登录不友好(要么前端删除token，要么后端生成黑名单)
>所以我又重新选择redis+token的方式来实现退出登录。

### 用户退出
 
### 用户查询