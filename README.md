# agent-notify

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.25-blue.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/hellolib/agent-notify.svg)](https://github.com/hellolib/agent-notify/releases)

一个面向 AI Agent 的通知配置工具。支持将 Claude Code、Codex 等 Agent 的事件通知推送到飞书和系统通知。

## 工作流程

<p align="center">
  <img src="assist/workflow.png" alt="工作流程图" />
</p>

## 功能特性

- 🖥️ **系统通知** - 支持 macOS、Linux、Windows 系统通知
- 📱 **飞书通知** - 支持飞书机器人消息推送
- 🔔 **事件订阅** - 支持多种事件类型：授权请求、等待输入、任务完成、任务失败
- 🔄 **自动安装** - 自动配置 Claude Code hooks 与 Codex notify
- 📦 **NPX 一键安装** - 无需手动下载，一条命令完成安装

## 安装

### 通过 NPX 运行（推荐）

```bash
npx agent-notify
```

首次运行会从 GitHub Releases 下载当前 npm 包版本对应平台的二进制文件，并安装到：

- macOS / Linux: `~/.agent-notify/agent-notify`
- Windows: `~/.agent-notify/agent-notify.exe`

之后每次运行 `npx agent-notify` 时都会检查本地二进制版本：

- 本地不存在：自动下载
- 本地版本落后：自动更新
- 本地版本不落后：直接运行本地二进制

launcher 不会持久修改你的 PATH，而是始终用绝对路径执行已安装的真实二进制。

> **注意**: Codex 支持目前为实验性功能，当前通过 `~/.codex/config.toml` 中的 `notify = ["/path/to/agent-notify", "handle-codex-notify"]` 接入，事件覆盖范围可能少于 Claude Code。

### 支持的平台

- macOS amd64
- macOS arm64
- Linux amd64
- Linux arm64
- Windows amd64
- Windows arm64

### 通过 npm 全局安装

```bash
npm install -g agent-notify
agent-notify
```

全局安装 npm 包后，运行方式与 `npx agent-notify` 一致：launcher 仍会把真实二进制安装到 `~/.agent-notify/` 并按版本检查后执行。

### 手动下载

从 [GitHub Releases](https://github.com/hellolib/agent-notify/releases) 下载对应平台的二进制文件。

## 快速开始

### 通过 launcher 使用

#### NPX

```bash
npx agent-notify
npx agent-notify init
npx agent-notify test system
npx agent-notify test feishu
npx agent-notify doctor
```

#### npm 全局安装后的 launcher 入口

```bash
npm install -g agent-notify
agent-notify
agent-notify init
agent-notify test system
agent-notify test feishu
agent-notify doctor
```

无论是 `npx agent-notify` 还是全局安装后的 `agent-notify`，入口仍然是 launcher；它会把真实二进制放在 `~/.agent-notify/` 下并按版本检查后执行。

### 通过手动下载的二进制使用

如果你下载了 release 二进制，并直接执行该文件，则可以使用：

```bash
./agent-notify
./agent-notify init
./agent-notify test system
./agent-notify test feishu
./agent-notify doctor
```

按提示选择：
1. 选择要配置的 Agent（Claude Code / Codex，单选）
2. 选择通知渠道（飞书和系统通知默认全选）
3. 如果选择 Claude Code：选择要接收通知的事件类型（默认全选4种事件）；Codex 无事件选择

## 支持的事件

| 事件 | 说明 |
|------|------|
| `permission_required` | Agent 需要授权（如执行命令） |
| `input_required` | Agent 等待用户输入 |
| `run_completed` | 任务执行完成 |
| `run_failed` | 任务执行失败 |

说明：Claude Code 当前支持完整的 hooks 事件映射，可通过 init 流程选择事件类型；Codex 当前走 notify 集成，不支持事件选择，通知触发时机以 Codex CLI 官方 notify 能力为准。

## 配置文件

agent-notify 自身配置位于 `~/.agent-notify/config.yaml`。

Agent 集成配置位置：

- Claude Code: `~/.claude/settings.json`（通过 hooks 接入）
- Codex: `~/.codex/config.toml`（通过 `notify = ["/path/to/agent-notify", "handle-codex-notify"]` 接入）

其中 Codex 当前只接入官方公开的 notify 能力，不再复用 Claude Code 的 hooks JSON 配置。

```yaml
version: 1
agent:
  claude_code:
    enabled: true
    install_scope: user
  codex:
    enabled: false  # 实验性功能
    install_scope: user
notify:
  claude_code:
    system:
      enabled: true
      events:
        - permission_required
        - input_required
        - run_completed
        - run_failed
    feishu:
      enabled: false
      events: []
  codex:
    system:
      enabled: false
      events: []  # Codex 不支持事件选择
    feishu:
      enabled: false
      events: []  # Codex 不支持事件选择
behavior:
  dedupe_seconds: 60
  send_timeout_seconds: 5
  locale: zh-CN
```

## 命令参考

### launcher 方式

```bash
npx agent-notify              # 通过 npx 进入交互式主菜单
npx agent-notify init         # 通过 npx 初始化配置
npx agent-notify test system  # 通过 npx 测试系统通知
npx agent-notify test feishu  # 通过 npx 测试飞书通知
npx agent-notify doctor       # 通过 npx 做环境诊断
npx agent-notify --help       # 通过 npx 查看帮助

agent-notify                  # npm 全局安装后的 launcher 入口
agent-notify init
agent-notify test system
agent-notify test feishu
agent-notify doctor
agent-notify --help
```

### 手动下载的二进制方式

```bash
./agent-notify              # 直接执行下载后的二进制
./agent-notify init
./agent-notify test system
./agent-notify test feishu
./agent-notify doctor
./agent-notify --help
```

## 飞书配置

首次使用飞书通知时，会自动引导你完成飞书 CLI 的初始化：
1. 扫码登录
2. 选择要使用的机器人

飞书的通知目标为机器人的所有者；

## 开发

### 使用 Makefile

```bash
# 查看所有可用命令
make help

# 构建二进制文件
make build

# 构建所有平台
make build-all

# 运行测试
make test

# 运行测试并生成覆盖率报告
make test-coverage

# 本地开发运行
make run

# 清理构建产物
make clean

# 安装到 GOPATH/bin
make install

# 代码格式化
make fmt

# 运行 go vet
make vet

# 运行 linter（需要安装 golangci-lint）
make lint

# 整理 go modules
make mod-tidy

# 运行 doctor 诊断
make doctor
```

### 直接使用 Go 命令

```bash
# 运行测试
go test ./...

# 本地开发运行
go run ./cmd/agent-notify

# 构建
go build -o agent-notify ./cmd/agent-notify
```
