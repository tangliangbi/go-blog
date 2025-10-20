# Blog Backend

这是一个使用 GORM 和 Gin 框架构建的博客后端项目，使用 MySQL 作为数据库。

## 目录结构

```
.
├── README.md
├── go.mod
├── main.go
├── api/
│   ├── routers/
│   │   └── routers.go         # 路由定义
│   └── v1/
│       ├── comments_controller.go
│       ├── posts_controller.go
│       └── user_controller.go
├── cmd/
│   └── server.go              # 服务器初始化和启动
├── config/
│   ├── config.go              # 配置文件解析
│   └── config.yaml            # 配置文件
├── middleware/
│   ├── cors.go                # 跨域处理
│   └── jwt.go                 # JWT认证中间件
├── repository/
│   ├── db/
│   │   ├── dao/               # 数据访问对象
│   │   │   ├── comments_dao.go
│   │   │   ├── errors.go
│   │   │   ├── posts_dao.go
│   │   │   └── user_dao.go
│   │   ├── database.go        # 数据库连接
│   │   └── db.go
│   └── model/
│       ├── migrate.go         # 数据库迁移
│       └── models.go          # 数据模型定义
├── service/
│   ├── comments/
│   │   └── comments_serivce.go
│   ├── posts/
│   │   └── posts_service.go
│   └── users/
│       └── user_service.go
└── utils/
    ├── assert/
    ├── logger/
    └── response/
        └── response.go        # 统一响应格式
```

## 数据模型

### User (用户)
- ID: 用户唯一标识
- Username: 用户名（唯一）
- Password: 密码（加密存储）
- Email: 邮箱（唯一）

### Post (文章)
- ID: 文章唯一标识
- Title: 文章标题
- Content: 文章内容
- UserID: 关联的用户ID

### Comment (评论)
- ID: 评论唯一标识
- Content: 评论内容
- PostID: 关联的文章ID
- UserID: 关联的用户ID

## 功能特性

1. 用户注册与登录（JWT认证）
2. 文章管理（创建、读取、更新、删除）
3. 评论功能（创建、读取、更新、删除）
4. 错误处理与日志记录
5. 数据库迁移

## 技术栈

- [Gin](https://github.com/gin-gonic/gin) - HTTP web 框架
- [GORM](https://gorm.io/) - ORM 库
- [MySQL Driver](https://github.com/go-sql-driver/mysql) - MySQL 驱动
- [jwt-go](https://github.com/dgrijalva/jwt-go) - JWT 实现

## 快速开始

### 环境要求

- Go 1.23+
- MySQL 5.7+

### 安装步骤

1. 克隆项目到本地：
   ```bash
   git clone <repository-url>
   cd blog-backend
   ```

2. 安装依赖：
   ```bash
   go mod tidy
   ```

3. 配置数据库：
   在 `config/config.yaml` 中修改数据库配置：
   ```yaml
   mysql:
     default:
       dialect: "mysql"
       dbHost: "127.0.0.1"     # 数据库地址
       dbPort: "3306"          # 数据库端口
       dbName: "blog"          # 数据库名称
       userName: "root"        # 用户名
       password: "root"        # 密码
       charset: "utf8mb4"      # 字符集
   ```

4. 启动项目：
   ```bash
   go run main.go
   ```

## API 路由

### 用户相关
- `POST /api/v1/register` - 用户注册
- `POST /api/v1/login` - 用户登录

### 文章相关
- `POST /api/v1/posts` - 创建文章
- `GET /api/v1/posts` - 获取文章列表
- `GET /api/v1/posts/:postId` - 获取指定文章
- `PUT /api/v1/posts/:postId` - 更新指定文章
- `DELETE /api/v1/posts/:postId` - 删除指定文章

### 评论相关
- `POST /api/v1/comments` - 创建评论
- `GET /api/v1/posts/:postId/comments` - 获取指定文章的评论列表
- `PUT /api/v1/comments/:id` - 更新指定评论
- `DELETE /api/v1/comments/:id` - 删除指定评论

## 项目配置

项目配置文件位于 `config/config.yaml`：

```yaml
system:
  domain: mall
  version: 1.0
  env: "dev"
  HttpPort: ":8080"           # 服务端口
  Host: "localhost"           # 服务主机
  UploadModel: "local"

mysql:
  default:
    dialect: "mysql"
    dbHost: "127.0.0.1"
    dbPort: "3306"
    dbName: "blog"
    userName: "root"
    password: "root"
    charset: "utf8mb4"
```

## 数据库迁移

项目会在启动时自动执行数据库迁移，创建所需的表结构。

## 日志

日志文件默认保存在 `logs/` 目录下。