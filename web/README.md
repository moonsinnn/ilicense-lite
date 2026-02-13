# ilicense-lite web

Nuxt 前端控制台，负责 License 管理 UI 与后端 API 代理（BFF）。

## Requirements

- Node.js `22+`
- pnpm `10+`

## Setup

```bash
cp .env.example .env
pnpm install
```

## Environment Variables

- `NUXT_API_BASE`: 后端 API 基础地址，默认 `http://127.0.0.1:8080`
- `NUXT_PUBLIC_SITE_URL`: 可选，用于公开链接场景

## Development

```bash
pnpm dev
```

## Quality Checks

```bash
pnpm lint
pnpm typecheck
pnpm build
```

## Notes

- `server/api/*` 中的路由会代理请求到后端。
- 登录后，服务端会透传 `Authorization` 或 `auth_token` Cookie 到后端。
