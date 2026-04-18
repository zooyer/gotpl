# gotpl
Golang service template



目录介绍

- common: 项目中公共目录
  - conf：解析配置文件
  - consts：定义常量
  - errnos：错误码定义
  - log：全局日志初始化
  - sdk：外部依赖
  - utils：本项目依赖的工具封装
- conf：配置文件，采用两段式命名（环境 + 集群）
  - dev.yaml、test.yaml、pre.hna.yaml、prod.hnb-batch-v.yaml
- dao：数据库初始化层，依赖本项目model结构，使用github.com/zooyer/miskit/micro集成gorm
- doc：当前项目所有相关文档，包括迭代过程中的资源图片，技术架构细节、资源源文件等
- dto：api接口的请求和响应的结构，通常model依赖dto中的子结构，避免重复定义
  - 对应handler结构的请求和响应结构
  - dto.go：公共通用的结构
- handler：gin的接口处理器，通常封装成一个结构体继承github.com/zooyer/miskit/micro中的一些方法
- middler：gin的中间件
- model：定义数据库中的模型字段
- service：服务的具体实现逻辑，尽量屏蔽http协议相关，为了后续支持其他协议
- test：测试目录
- web：一般情况下都是该服务后端管理的前端UI，通常使用React、TypeScript、Antd Pro
- build.sh：编译脚本，支持交叉编译，支持压缩
- start.sh：启动脚本
- main.go：服务主入口
