# Go + Ent + Casbin + Postgres 模板

## 目录结构

```yaml
.
├── grphql
│   ├── ent.graphql # 基于 entgql 配置自动生成的 GraphQL 接口
│   └── *.graphql # 自定义 GraphQL 接口
└── backend
    ├── config # 基于 viper 的配置定义与加载
    ├── ent
    │   ├── entc.go # ent 代码生成配置
    │   ├── schema # ent 数据模型定义
    │   └── ... # 其他由代码生成器生成的文件
    ├── graph # graphql 接口 resolver
    │   ├── generated.go # 自动生成
    │   ├── model # 自动生成的 GraphQL 数据模型
    │   ├── reqerr # 请求错误定义
    │   ├── resolver.go # 提供 GraphQL directive 实现以及服务初始化
    │   └── *.resolvers.go # 由 gqlgen 代码生成器生成未实现的新接口/已实现的原有接口，其名称与 graphql 文件名称对应
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
      login(input: { phone: "12345678910", password: "root" }) {
        token
        user {
          id
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

## 指南——创建 CRUD API

为了开发便利，请将 `graphql.introspection` 和 `graphql.playground` 设置为 `true`。

以 PublicKey 为例。（以下命令均在 `backend` 目录下执行）

### 定义数据模式

运行下面的命令，生成 PublicKey 数据模型样板。

```sh
go run -mod=mod entgo.io/ent/cmd/ent new PublicKey
```

会生成在 `ent/schema/` 生成 `publickey.go`，文件内容如下：

```go
package schema

import "entgo.io/ent"

// PublicKey holds the schema definition for the PublicKey entity.
type PublicKey struct {
	ent.Schema
}

// Fields of the PublicKey.
func (PublicKey) Fields() []ent.Field {
	return nil
}

// Edges of the PublicKey.
func (PublicKey) Edges() []ent.Edge {
	return nil
}
```

根据 [ent 文档](https://entgo.io/docs/schema-fields)定义字段和边（关联关系）。PublicKey 具有以下字段，并且属于一个特定的用户。我们还为 `expired_at` 字段创建一个索引。

```go
// import ...

func (PublicKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("name").Optional(),
		field.String("description").Optional(),
		field.String("type").Optional(),
		field.String("status").Optional(),
		field.Time("expired_at").Optional(),
	}
}

func (PublicKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Immutable().
			Annotations(entsql.OnDelete(entsql.Cascade)),
	}
}

func (PublicKey) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("expired_at"),
	}
}
```

注意，本模板项目中包含 ID（xid 适配）与自动更新的“更新创建时间”字段 [mixin](https://entgo.io/docs/schema-mixin)，可以直接引入。这样在数据库表中，就有了 xid 格式的主键 ID 和自动维护的 `created_at` 和 `updated_at` 字段了。

```go
import "github.com/hawa130/serverx/ent/schema/mixinx"

func (PublicKey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixinx.XId{},
		mixinx.Time{},
	}
}
```

接下来运行生成器，生成数据模型相关操作 API 的代码。

```sh
go generate .
```

此时可以在程序中进行数据库的 [CRUD](https://entgo.io/docs/crud) 操作了。

### 接入 GraphQL

#### 查询

ent 具有内置的 GraphQL 生成器支持，我们只需要使用 entgql 提供的 annotation 即可自动生成相关的 graphql schema。

```go
func (PublicKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
```

其中，`RelayConnection()` 表明需要生成 Relay 风格的分页（可选）。`QueryField()` 表明数据模式的字段需要暴露给 GraphQL 的 `query` 以便查询使用，`Mutations(...)` 表明为数据模式生成 GraphQL 的 `input` 以便创建和更新使用。

更新边的定义。

```go
func (PublicKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Immutable().
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput|entgql.SkipMutationCreateInput),
				entsql.OnDelete(entsql.Cascade),
			),
	}
}
```

添加了 `entgql.Skip(...)` annotation，表明在生成 GraphQL 时，创建和更新输入不包含用户字段，以免修改用户信息。

完成后运行生成器 `go generate .`，生成器会更新 `graphql/ent.graphql` 文件，生成新的接口和类型。

并且由于执行了 gqlgen，因此 `backend/graph/ent.resolvers.go` 也会创建一个未实现的接口。（如果没有使用 Relay Connection 接口会不一样）

```go
// PublicKeys is the resolver for the publicKeys field.
func (r *queryResolver) PublicKeys(ctx context.Context, after *entgql.Cursor[xid.ID], first *int, before *entgql.Cursor[xid.ID], last *int, orderBy *ent.PublicKeyOrder, where *ent.PublicKeyWhereInput) (*ent.PublicKeyConnection, error) {
	panic(fmt.Errorf("not implemented: PublicKeys - publicKeys"))
}
```

实现这个接口吧，结合 ent 实现起来非常简单。（[参考](https://entgo.io/docs/tutorial-todo-gql-filter-input#configure-gql)）

```go
func (r *queryResolver) PublicKeys(ctx context.Context, after *entgql.Cursor[xid.ID], first *int, before *entgql.Cursor[xid.ID], last *int, orderBy *ent.PublicKeyOrder, where *ent.PublicKeyWhereInput) (*ent.PublicKeyConnection, error) {
	return r.client.PublicKey.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithPublicKeyOrder(orderBy),
			ent.WithPublicKeyFilter(where.Filter),
		)
}
```

现在可以在 GraphQL playground（默认在 http://localhost:8080/playground）中使用 `publicKeys` 来查询了，可以试试下面的查询。

```graphql
query ListPublicKeys {
  publicKeys(first: 10) {
    edges {
      node {
        key
      }
      cursor
    }
    pageInfo {
      hasPreviousPage
      hasNextPage
    }
    totalCount
  }
}
```

如果不使用 Relay 风格分页的话，查询写起来会更简洁，但是就没有分页了，适合数据量小的情况。

#### 创建、更新与删除

由于业务的多样性，ent 并不会自动生成 CRUD 中的 CUD 的接口，因此需要自己去定义。

在 `graphql/` 目录下创建 `publickey.graphql`。并在里面声明几个接口。其中，`PublicKey` 类型，以及 `CreatePublicKeyInput` 和 `UpdatePublicKeyInput` 输入均在 `ent.graphql` 中已经生成，可以直接拿来用。

```graphql
extend type Mutation {
  """
  为指定用户创建公钥
  """
  createPublicKey(uid: ID!, input: CreatePublicKeyInput!): PublicKey!

  """
  更新公钥
  """
  updatePublicKey(id: ID!, input: UpdatePublicKeyInput!): PublicKey!

  """
  删除公钥
  """
  deletePublicKey(id: ID!): Boolean!
}
```

完成后运行生成器 `go generate .`，gqlgen 会在 `graph/` 目录下生成 `publickey.resolvers.go` 并包含未实现的接口。实现这些接口也很容易。

```go
// CreatePublicKey is the resolver for the createPublicKey field.
func (r *mutationResolver) CreatePublicKey(ctx context.Context, uid xid.ID, input ent.CreatePublicKeyInput) (*ent.PublicKey, error) {
	return ent.FromContext(ctx).PublicKey.Create().SetInput(input).SetUserID(uid).Save(ctx)
}

// UpdatePublicKey is the resolver for the updatePublicKey field.
func (r *mutationResolver) UpdatePublicKey(ctx context.Context, id xid.ID, input ent.UpdatePublicKeyInput) (*ent.PublicKey, error) {
	return ent.FromContext(ctx).PublicKey.UpdateOneID(id).SetInput(input).Save(ctx)
}

// DeletePublicKey is the resolver for the deletePublicKey field.
func (r *mutationResolver) DeletePublicKey(ctx context.Context, id xid.ID) (bool, error) {
	if err := ent.FromContext(ctx).PublicKey.DeleteOneID(id).Exec(ctx); err != nil {
		return false, err
	}
	return true, nil
}
```

现在可以打开 playground 来使用这些接口了。

#### 说明

上面只是一个简单的创建 CRUD 接口的示例。在实际的业务中，由于权限的复杂性，代码会有所不同。仓库中的 PublicKey 是完全版，考虑了 CRUD 的鉴权，可以参考。
