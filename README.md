# ilicense-lite

[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](LICENSE)

**ilicense-lite** 是一个轻量级的 License 管理与校验解决方案，面向需要为软件产品提供授权控制能力的开发者和团队。

项目提供 **服务端 License 管理能力** 以及 **客户端 SDK**，用于完成 License 的签发、校验与生命周期管理，帮助开发者以较低成本构建清晰、可维护的授权体系。

ilicense-lite 专注于「够用、清晰、可扩展」的设计原则，避免引入复杂的计费、自动升级或 SaaS 绑定逻辑，适合作为中小项目、内部系统或商业产品的 License 基础设施。

---

## 核心特性

- 产品与版本管理
- License 的创建、吊销与状态控制
- 客户与机构维度的授权管理
- 轻量的版本 / 部署包管理（一版本一包）
- 标准化 License 校验接口
- 多语言客户端 SDK（可扩展）

---

## 设计目标

- **简单**：减少不必要的配置和概念
- **清晰**：数据模型与授权流程直观可理解
- **可扩展**：为后续企业版或定制场景保留扩展空间
- **可自托管**：不依赖第三方 SaaS

---

## 非目标（Not Goals）

- 自动升级与灰度发布
- 商业计费系统
- 复杂权限与多租户计费模型
- 云服务绑定

---

## 适用场景

- 独立开发者的软件授权管理
- 中小型商业产品的 License 控制
- 企业内部系统或私有部署软件
- 需要自定义授权策略的项目

---

## 快速开始

### 1. 克隆仓库

```bash
git clone https://github.com/your-username/ilicense-lite.git
cd ilicense-lite
```