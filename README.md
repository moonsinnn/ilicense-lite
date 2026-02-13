# ilicense-lite

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

`ilicense-lite` 是一个轻量的 License 管理与校验项目，包含：

- `app/`: Go 后端（License 管理 API、签发与校验逻辑）
- `web/`: Nuxt 前端（管理控制台 + BFF 接口）

## Features

- 产品、客户、签发机构管理
- License 创建、续期、删除、激活校验
- JWT 鉴权
- OpenTelemetry + Prometheus 中间件
- 前后端分离，支持本地开发与自托管部署

## Repository Structure

```text
.
├── app/               # Go backend
│   ├── main.go
│   ├── etc/
│   ├── router/
│   ├── controller/
│   ├── service/
│   ├── dao/
│   └── script/db/schema.sql
└── web/               # Nuxt frontend + server APIs
    ├── app/
    ├── server/
    └── package.json
```

## Prerequisites

- Go `1.24+`
- Node.js `22+`
- pnpm `10+`
- MySQL `8+`

## Quick Start

### 1. Backend

```bash
cd app
cp etc/app.example.yaml etc/app.yaml
```

编辑 `app/etc/app.yaml`，填入数据库连接信息与安全配置：

- `app.jwt_secret`
- `mysql_demo.data_source.*`

初始化数据库（示例）：

```bash
mysql -u <user> -p <database> < script/db/schema.sql
```

启动后端：

```bash
go run . -config etc/app.yaml
```

默认监听：`http://127.0.0.1:8080`

### 2. Frontend

```bash
cd web
cp .env.example .env
pnpm install
pnpm dev
```

默认地址：`http://127.0.0.1:3000`

确保 `web/.env` 中 `NUXT_API_BASE` 指向后端地址（默认 `http://127.0.0.1:8080`）。

## Development

### Backend

```bash
cd app
GOCACHE=/tmp/go-build-cache go test ./...
```

### Frontend

```bash
cd web
pnpm lint
pnpm typecheck
pnpm build
```

## API and Auth

- 前端通过 `web/server/api/*` 代理后端 API。
- 后端使用 `Authorization: Bearer <token>`。
- 未鉴权路由由业务层自行控制；如需强制鉴权，路由中应接入 `RequireAuth()`。

## Configuration

- 后端配置文件：`app/etc/app.yaml`
- 前端环境变量：`web/.env`

生产环境建议：

- 使用强随机 `JWT_SECRET`
- 使用 `LICENSE_MASTER_KEY` 环境变量覆盖默认主密钥
- 关闭调试日志并做好日志轮转和备份

## Open Source Governance

- Contributing: [CONTRIBUTING.md](CONTRIBUTING.md)
- Security Policy: [SECURITY.md](SECURITY.md)
- Code of Conduct: [CODE_OF_CONDUCT.md](CODE_OF_CONDUCT.md)

## License

Apache-2.0. See [LICENSE](LICENSE).
