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
    │   ├── generated.go # 自动生成的适配 ent 数据模型的 resolver
    │   ├── model # 自动生成的 GraphQL 数据模型
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

    开发过程中建议将 `graphql.introspection` 和 `graphql.playground` 设置为 `true` 以便于使用 GraphQL Playground 进行调试，默认地址为 http://localhost:8080/playground 。
    
    生成 JWT 所需的密钥对：
    
    ```shell
    bash genkey.sh
    ```

3. 代码生成
   
    如果添加了新的 ent 数据模型 或者 GraphQL 接口定义，需要重新生成代码。
    ```shell
    go generate .
    ```
    
    在运行代码生成时，所有定义在 `graphql` 文件夹下的 `.graphql` 文件都会被自动加载并生成对应的 GraphQL 接口。
    
    相关文档：[ent](https://entgo.io/zh/docs/code-gen)，[gqlgen](https://gqlgen.com/)。
    
4. 运行
   
    ```shell
    go run .
    ```

    默认情况下，修改配置文件并保存后无需重启程序，程序会自动加载新的配置。

    默认超级管理员登录方式为手机号和密码登录，手机号为 `12345678910`，密码为 `root`。可以通过下面的 GraphQL 接口进行登录：

    ```graphql
    mutation Login {
       login(input: {
          phone: "12345678910",
          password: "root",
       }) {
          token
          user {
             id
             username
          }
       }
    }
    ```
   
    接口返回的 `token` 即为 JWT，可以在请求头中添加 `Authorization` 字段进行身份验证（格式为 `Bearer <token>`）。

## 细节

本模板项目包含两个数据模型：`User`（用户）和 `PublicKey`（公钥）。其中一个用户可以有多个公钥，每个用户也可以拥有多个子用户。数据模型均采用的 [xid](https://github.com/rs/xid) 作为全局唯一 ID。

权限采用 RBAC 模型。
- 默认的超级管理用拥有全部权限。
- 创建用户时，通过 hook 实现对用户密码的加密。
- 在用户创建公钥时，通过 hook 授予用户对该公钥的权限。
- 在用户创建子用户时，通过 hook 授予用户对该子用户的权限。

注意：由于 casbin 的字段并没有和其他数据模型有关联，因此在数据删除时，需要使用 hook 手动清理 casbin 数据表中的数据。
