# Go + Ent + Casbin + Postgres 模板

## 目录结构

```yaml
.
├── grphql
│   ├── ent.graphql # 自动生成
│   └── schema.graphql # 自定义 GraphQL 接口
└── backend
    ├── config # 基于 viper 的配置定义与加载
    ├── ent
    │   ├── entc.go # ent 代码生成配置
    │   ├── schema # ent 数据模型定义
    │   └── ... # 其他由代码生成器生成的文件
    ├── graph # graphql 接口 resolver
    │   ├── generated.go # 自动生成
    │   ├── model # 自动生成
    │   ├── reqerr # 请求错误定义
    │   ├── resolver.go # 提供 GraphQL directive 实现以及服务初始化
    │   └── *.resolvers.go # 自定义 resolver，由 gqlgen 代码生成器生成的未实现接口（或者自己已经实现）
    ├── internal # 内部服务
    │   ├── adapter # casbin ent 适配器
    │   ├── auth # 认证服务，提供 JWT 生成、验证，用户密码加密、验证，用户权限验证等工具
    │   ├── database # 数据库客户端与seed
    │   ├── hookx # ent hook 扩展
    │   ├── logger # 基于 zap 的日志服务
    │   ├── perm # 权限管理
    │   ├── rule # ent 隐私层规则定义
    │   └── xidgql # xid graphql 类型适配
    ├── main.go # 入口
    ├── server.go # 服务相关
    ├── gqlgen.yml # gqlgen 配置
    ├── perm-model.conf # casbin 权限模型定义
    └── generate.go # 定义代码生成
```

## 使用

以下操作均在 `backend` 目录下进行。

1. 安装依赖
    
    ```shell
    go mod tidy
    ```

2. 初始化配置
    
    ```shell
    cp config.default.toml config.toml
    ```
    
    并根据实际情况修改 `config.toml`。
    
    生成 JWT 所需的密钥对：
    
    ```shell
    bash genkey.sh
    ```

3. 代码生成
    
    如果添加了新的 ent 数据模型 或者 GraphQL 接口定义，需要重新生成代码。
        
    ```shell
    go generate .
    ```

4. 运行
    
    ```shell
    go run .
    ```
