# iam
统一资源调度系统

[参考项目地址](https://github.com/marmotedu/iam)

[文档地址](https://www.apifox.cn/apidoc/shared-05a28f7d-1714-41f6-94fa-566e82e0bfcf)
`里面提供的案例参数没有太多参考意义，因为他强依赖于数据库环境。详细说明请看接口文档的参数说明。`

#### 项目部署地址
http://43.142.147.229:12345

`可以参照这个根地址发送api请求`

## 第一阶段
### 功能特性
对用户信息的基本操作

**无需登录的操作**：
- 注册（需要绑定邮箱，会发送验证码）
- 登录，输入账户密码进行验证
- 获取用户信息（通过用户id）
- 获取用户信息列表（所有用户）
- 获取验证码（传入邮箱）


**需要登录验证**:
- 退出登录
- 修改密码
- 修改邮箱
- 修改其他用户信息（昵称，github账户，博客地址，公司，学校等）


### 代码设计
- 包含常用集合，如mysql、log、redis等组件的配置，配置文件采用viper进行读取，对于这些常用包的初始化我放在了pkg文件中，并统一提供了option用于初始化，初始化后会留有一份单例保存在改文件下，用于后续封装函数使用。
- service, store, cache层都采用接口的形式进行设计，上层访问下层接口，方便依赖注入，如cache支持 gocache 和 redis 两种形式取实现
- 错误处理，设计了错误码，底层报错统一包装往上抛，过程中不打印error日志，只有在返回给客户端的函数中打印错误及其错误链路，避免错误日志冗余。
- 三层之间的都会传递一个context对象，用于处理超时和 request id 的传参。

更多关于我的设计思路可以查看[develop](docs/develop/basic.md)文件


### 快速部署
`前后端分离，所以要额外部署前端项目`

`但前端还没写好`

后端部署
修改[配置文件](config/config.example.yaml)重命名为`config.yaml`, 填写相关配置.
2. 普通部署

```shell
go mod tidy # 下载依赖

# 检查配置文件，启动mysql容器或服务
# 启动mysql镜像 
docker run -itd --name mysql -p3306:3306 -e MYSQL_ROOT_PASSWORD=[your_password] mysql:latest 
# 启动后需要将 deploy 目录下的 iam.sql 文件导入到mysql中生成数据表

# 默认方式启动服务
go run cmd/apiserver/main.go 

# 二进制文件启动
# 内置了三个参数 port, mode, config 分别表示服务端口，服务模式（debug、release），服务配置文件位置
go build cmd/apiserver/main.go
./main  [options]

```

2. Docker 部署

windows 命令行下面的换行可能会报错，需要去掉换行符，将所有指令放到同一行

```shell
# 1.构建运行mysql容器
docker run --name iam-mysql -itd -p 3306:3306 \
-e MYSQL_ROOT_PASSWORD=123456 \
-v $PWD/storage/mysql/data:/var/lib/mysql \
--restart=always \
mysql \
--character-set-server=utf8mb4 \
--collation-server=utf8mb4_general_ci 


# windows
docker run --name iam-mysql -itd -p 3306:3306 -e MYSQL_ROOT_PASSWORD=123456 -v $PWD/storage/mysql/data:/var/lib/mysql --restart=always  mysql --character-set-server=utf8mb4 --collation-server=utf8mb4_general_ci 

# 2. 读取数据库表
# 连接进mysql容器，执行deploy中的sql文件

# 3. 创建iam镜像
docker build . -t iam-server

# 4. 运行iam容器
docker run --name iam-server \
-itd -p 123456:123456 \
--net=host \
-v $PWD/config:/config \
-v $PWD/storage/log:/storage \
iam-server


# windows
docker run --name iam-server  -itd -p 12345:12345 --net=host -v $PWD/config:/config -v $PWD/storage/log:/storage iam-server

```

`docker-compose 还没写好`


## 关于作者
    YeeBing Chen yeebingchen@qq.com
## 许可证
IAM is licensed under the MIT. See [LICENSE](LICENSE) for the full license text.
