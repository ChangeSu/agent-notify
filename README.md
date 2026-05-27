<div align="center">

# Agent Notify

<p align="center"> 一个面向 AI Agent 的通知配置工具 </p>

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.25-blue.svg)](https://go.dev/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![Release](https://img.shields.io/github/v/release/hellolib/agent-notify.svg)](https://github.com/hellolib/agent-notify/releases)

</div>

## 项目简介

一个面向 AI Agent 的通知配置工具。支持将 Claude Code、Codex、Qoder CN（通义灵码）等 Agent 的事件通知推送到飞书、企业微信和系统通知。

## ❤️ 赞助

<a href="https://www.ddshub.cc/register?aff=E7N6PDYWW4N5">
<img src="assist/ddshub/logo.png" alt="DDSHub" />
</a>

感谢 **[DDS（呆呆兽）](https://www.ddshub.cc/register?aff=E7N6PDYWW4N5)** 赞助本项目！呆呆兽是一家专注 Claude 和 CodeX 的可靠高效 API 中转站，为个人和企业用户提供极具性价比的国内 Claude / CodeX API 直连加速服务。支持 Claude Haiku / Opus / Sonnet 等满血模型。企业客户更可享受定制化分组和技术支持服务。

访问 [DDS（呆呆兽）](https://www.ddshub.cc/register?aff=E7N6PDYWW4N5) 获取更多详情。

## 功能特性

- 🖥️ **系统通知** - 支持 macOS、Linux、Windows 系统通知
- 📱 **飞书通知** - 支持一键扫码绑定、支持飞书机器人消息推送
- 💬 **企业微信通知** - 支持通过企业微信群机器人 Webhook 推送通知消息
- 🔔 **事件订阅** - Claude Code 支持 4 种事件；Codex 支持 2 种事件

### 支持的事件

| 事件 | 说明 | Claude Code | Codex | Qoder CN |
|------|------|:---:|:---:|:---:|
| `permission_required` | Agent 需要授权（如执行命令） | ✅ | ✅ | ✅ |
| `input_required` | Agent 等待用户输入 | ✅ | — | — |
| `run_completed` | 任务执行完成 | ✅ | ✅ | ✅ |
| `run_failed` | 任务执行失败 | ✅ | — | ✅ |

说明：

- Claude Code 通过 `~/.claude/settings.json` 的 hooks 订阅四个事件（`PermissionRequest`、`Notification`、`Stop`、`PostToolUseFailure`）。
- Codex 通过 `~/.codex/hooks.json` 订阅 `PermissionRequest` 与 `Stop`，分别映射到 `permission_required` 与 `run_completed`。`input_required` 与 `run_failed` Codex 目前没有对应 hook，因此暂不支持。
- Qoder CN 通过 `~/.lingma/settings.json` 的 hooks 订阅 `PreToolUse`、`Stop` 与 `PostToolUseFailure`，分别映射到 `permission_required`、`run_completed` 与 `run_failed`。

### 支持的平台

MacOS 、Linux 、Windows

## 快速开始

```bash
npx agent-notify
```

首次运行会从 GitHub Releases 下载当前 npm 包版本对应平台的二进制文件，并安装到：

- macOS / Linux: `~/.agent-notify/agent-notify`
- Windows: `~/.agent-notify/agent-notify.exe`

之后每次运行都会检查本地二进制版本：不存在则自动下载，版本落后则自动更新，否则直接运行。launcher 不会持久修改 PATH，始终用绝对路径执行。

> **注意**: Codex 通过 `~/.codex/hooks.json` 接入官方 hooks 系统，目前仅订阅 `PermissionRequest`、`Stop` 两个事件。首次安装后请在 codex 内运行 `/hooks` 完成 trust 审核。


## 配置

agent-notify 自身配置位于 `~/.agent-notify/config.yaml`。Agent 集成配置位置：

- Claude Code: `~/.claude/settings.json`（写入 hooks → 命令 `agent-notify handle-claude-hook`）
- Codex: `~/.codex/hooks.json`（写入 hooks → 命令 `agent-notify handle-codex-hook`，需在 codex 内运行 `/hooks` 完成 trust）
- Qoder CN: `~/.lingma/settings.json`（写入 hooks → 命令 `agent-notify handle-qodercn-hook`）

### 企业微信机器人配置

1. **创建单人通知群**：在企业微信中发起群聊（随便拉几个同事），创建成功后**不要在群里发言**，直接将其他人移出，此时该群将变成你的单人通知群；
2. **添加机器人**：「群设置」->「消息推送」->「添加」-> 「自定义消息推送」，命名并保存；
3. **获取 Webhook 地址**：复制生成的地址，格式类似 `https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=xxx`；
4. **绑定配置**：运行 `npx agent-notify`，在配置向导中选择启用企业微信渠道，粘贴 Webhook URL 即可；
> 旧版企业微信添加机器人步骤：「群设置」->「群机器人」->「添加机器人」-> 「新建机器人」，命名并保存

## 工作流程

<p align="center">
  <img src="assist/workflow.png" alt="工作流程图" />
</p>

## 效果图

| |                                                              |
|:---:|:------------------------------------------------------------:|
| <img src="assist/launch-setting.png" alt="软件配置" width="75%"> |  <img src="assist/feishu-bind.png" alt="飞书绑定" width="75%">   |
| **软件配置** |                           **飞书绑定**                           |
| <img src="assist/feishu-notify-phone.png" alt="飞书通知" width="50%"> | <img src="assist/wecom-notify.jpg" alt="企业微信通知" width="55%"> |
| **飞书通知** |                          **企业微信通知**                          |
| ![系统通知](assist/system-notify.png) |                                                              |
| **系统通知** |                                                              |

## Friendship Link

Thanks for the support and feedback from the friends at [LINUX DO](https://linux.do/).
