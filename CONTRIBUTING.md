# Contributing Guide

感谢你参与 `ilicense-lite`。

## Before You Start

- 先阅读 `README.md` 并完成本地启动。
- 提交前确保测试和静态检查通过。
- 变更尽量保持小而聚焦，避免一个 PR 混入无关修改。

## Branch and Commit

- 分支命名建议：`feat/<topic>`、`fix/<topic>`、`docs/<topic>`、`refactor/<topic>`
- 提交信息建议使用简洁前缀：
  - `feat:`
  - `fix:`
  - `docs:`
  - `refactor:`
  - `test:`
  - `chore:`

示例：

```text
feat: add license renew validation
fix: avoid logging large request body
docs: update backend setup section
```

## Local Checks

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

## Pull Request Checklist

- [ ] 只包含与目标相关的改动
- [ ] 本地检查全部通过
- [ ] 必要文档已更新（README/配置说明/API 行为）
- [ ] 若影响行为，已补充或更新测试
- [ ] PR 描述包含变更动机、实现方式、验证结果

## Code Style

- Go: 运行 `gofmt`
- TypeScript/Vue: 通过 ESLint + typecheck
- 避免提交调试日志和敏感信息

## Reporting Issues

- Bug 或建议请使用仓库中的 issue 模板提交。
- 安全问题请不要公开提交 issue，见 `SECURITY.md`。
