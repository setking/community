# Community

一个基于 Go + Gin + GORM 后端和 Vue.js 前端的现代化社区应用平台。

## 项目描述

Community 是一个功能完整的社区论坛系统，支持用户注册、登录、发帖、评论、投票等核心功能。后端采用 Go 语言开发，使用 Gin 框架提供 RESTful API，GORM 作为 ORM 工具连接 MySQL 数据库，并集成 Redis 缓存和 JWT 认证。前端基于 Vue 3 + TypeScript 构建，使用 Vuetify 组件库提供美观的 Material Design 界面。

## 功能特性

- **用户管理**: 用户注册、登录、JWT 认证、权限控制
- **社区管理**: 创建和管理社区板块
- **帖子系统**: 发布、编辑、删除帖子，支持分页浏览
- **评论系统**: 多级评论、回复功能
- **投票系统**: 对帖子和评论进行点赞/踩
- **API 文档**: 集成 Swagger 自动生成 API 文档
- **日志系统**: 结构化日志记录，支持多种输出格式
- **配置管理**: 灵活的配置文件支持，支持环境变量覆盖
- **数据库迁移**: 自动生成数据库表结构

## 技术栈

### 后端 (backed/)

- **语言**: Go 1.25.0
- **框架**: Gin (HTTP 框架)
- **ORM**: GORM (数据库操作)
- **数据库**: MySQL
- **缓存**: Redis
- **认证**: JWT
- **配置**: Viper
- **日志**: Zap + Lumberjack
- **ID 生成**: Sonyflake (分布式 ID)
- **验证**: Go Playground Validator
- **密码加密**: go-password-encoder

### 前端 (front/app/)

- **框架**: Vue 3 + TypeScript
- **构建工具**: Vite
- **UI 库**: Vuetify 3 (Material Design)
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **HTTP 客户端**: Axios
- **表单验证**: Vee-Validate
- **样式**: Sass
- **包管理**: pnpm

## 快速开始

### 前提条件

- Go 1.25.0 或更高版本
- Node.js 18+ 和 pnpm
- MySQL 8.0+
- Redis 6.0+

### 安装

1. **克隆项目**

   ```bash
   git clone https://github.com/setking/community.git
   cd community
   ```

2. **后端设置**

   ```bash
   cd backed
   # 安装依赖
   go mod download
   # 配置数据库和 Redis (编辑 config.yaml)
   cp config.yaml.example config.yaml
   ```

3. **前端设置**
   ```bash
   cd ../front/app
   # 安装依赖
   pnpm install
   ```

### 运行

1. **启动后端服务**

   ```bash
   cd backed
   go run main.go
   ```

   服务将在 `http://localhost:8080` 启动

2. **启动前端开发服务器**
   ```bash
   cd front/app
   pnpm dev
   ```
   前端将在 `http://localhost:5173` 启动

### 构建生产版本

1. **后端构建**

   ```bash
   cd backed
   go build -o community main.go
   ```

2. **前端构建**
   ```bash
   cd front/app
   pnpm build
   ```

## API 文档

项目集成了 Swagger UI 用于 API 文档查看。启动后端服务后，访问 `http://localhost:8080/swagger/index.html` 查看完整的 API 文档。

主要 API 端点包括：

- 用户相关: `/user/*`
- 社区相关: `/community/*`
- 帖子相关: `/post/*`
- 评论相关: `/comment/*`
- 投票相关: `/vote/*`

## 项目结构

```
community/
├── backed/                          # 后端 Go 服务
│   ├── app/                         # 应用入口
│   ├── config/                      # 配置管理
│   ├── controller/                  # 控制器层
│   │   ├── comment/                 # 评论控制器
│   │   ├── community/               # 社区控制器
│   │   ├── post/                    # 帖子控制器
│   │   ├── user/                    # 用户控制器
│   │   └── vote/                    # 投票控制器
│   ├── dao/                         # 数据访问层
│   ├── forms/                       # 请求表单结构
│   ├── global/                      # 全局变量
│   ├── initialize/                  # 初始化逻辑
│   ├── middlewares/                 # 中间件
│   ├── models/                      # 数据模型
│   ├── pkg/                         # 公共包
│   ├── routes/                      # 路由定义
│   ├── swagger/                     # Swagger 文档
│   ├── utils/                       # 工具函数
│   ├── validators/                  # 验证器
│   ├── config.yaml                  # 配置文件
│   ├── go.mod                       # Go 模块文件
│   ├── main.go                      # 主入口文件
│   ├── Makefile                     # 构建脚本
│   └── README.md                    # 后端说明
├── front/                           # 前端 Vue.js 应用
│   └── app/                         # 前端源码
│       ├── public/                  # 静态资源
│       ├── src/                     # 源码目录
│       │   ├── components/          # Vue 组件
│       │   ├── layouts/             # 布局组件
│       │   ├── pages/               # 页面组件
│       │   ├── router/              # 路由配置
│       │   ├── stores/              # Pinia 状态管理
│       │   ├── styles/              # 样式文件
│       │   ├── types/               # TypeScript 类型定义
│       │   └── utils/               # 工具函数
│       ├── package.json             # 项目配置
│       ├── tsconfig.json            # TypeScript 配置
│       ├── vite.config.mts          # Vite 配置
│       └── index.html               # HTML 模板
└── README.md                        # 项目总 README
```

## 配置说明

### 后端配置 (backed/config.yaml)

```yaml
mysql:
  host: localhost
  port: 3306
  user: root
  password: your_password
  dbname: community

redis:
  host: localhost
  port: 6379
  password: ''
  db: 0

jwt:
  secret: your_jwt_secret
  expire: 3600

server:
  port: 8080
  mode: debug
```

### 环境变量

支持通过环境变量覆盖配置：

- `MYSQL_HOST`, `MYSQL_PORT`, `MYSQL_USER`, `MYSQL_PASSWORD`, `MYSQL_DBNAME`
- `REDIS_HOST`, `REDIS_PORT`, `REDIS_PASSWORD`, `REDIS_DB`
- `JWT_SECRET`, `JWT_EXPIRE`
- `SERVER_PORT`, `SERVER_MODE`

## 测试

### 后端测试

```bash
cd backed
go test ./...
```

### 前端测试

```bash
cd front/app
pnpm lint
```

## 部署

### Docker 部署 (推荐)

1. 构建后端镜像

```bash
cd backed
docker build -t community-backend .
```

2. 构建前端镜像

```bash
cd front/app
docker build -t community-frontend .
```

3. 使用 Docker Compose 启动

```bash
docker-compose up -d
```

### 传统部署

1. 配置生产环境数据库和 Redis
2. 构建后端可执行文件
3. 构建前端静态文件
4. 配置反向代理 (Nginx/Apache)
5. 启动服务

## 贡献指南

欢迎贡献代码！请遵循以下步骤：

1. Fork 本项目
2. 创建特性分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 创建 Pull Request

### 代码规范

- 后端: 遵循 Go 官方代码规范，使用 `gofmt` 格式化代码
- 前端: 使用 ESLint 和 Prettier 进行代码检查和格式化
- 提交信息: 使用清晰的英文描述，遵循 Conventional Commits 规范

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。

## 联系方式

- 项目维护者: [setking](https://github.com/setking)
- 问题反馈: [Issues](https://github.com/setking/community/issues)

---

**注意**: 这是一个学习和演示项目，不建议直接用于生产环境。如需生产部署，请进行适当的安全审计和性能优化。
