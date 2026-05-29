# Go ADK 实战系列配套代码

本仓库托管 Go ADK 实战系列的全部配套代码，每个案例均可独立运行。

## 目录结构

```
01-environment/          # 模块 1：环境与安装
02-quickstart/          # 模块 2：快速入门
03-tools/               # 模块 3：工具使用
04-memory/              # 模块 4：记忆与上下文
05-multi-agent/         # 模块 5：多 Agent 协作
06-streaming/           # 模块 6：流式交互
07-deployment/          # 模块 7：部署与运维
08-a2a/                 # 模块 8：A2A 协议
09-advanced/            # 模块 9：进阶主题
10-realworld/           # 模块 10：实战综合
```

## 使用方法

每个子目录是一个独立的 Go Module，可直接运行：

```bash
cd 02-quickstart/adk-go-quickstart-hello-world
go mod tidy
go run agent.go
```

## 前置要求

- Go 1.24.4+
- ADK Go v0.2.0+
- 环境变量 `GOOGLE_API_KEY`（部分案例需要）

## 获取 API Key

访问 [Google AI Studio](https://aistudio.google.com/app/apikey) 创建 API Key，并写入 `.env` 文件：

```bash
export GOOGLE_API_KEY="your-api-key"
```

## 文章索引

| 模块 | 文章 |
|------|------|
| 模块 0 | [系列介绍与学习路线图](https://rexai.top/tutorials/go-adk/) |
| 模块 1 | [环境与安装](https://rexai.top/tutorials/go-adk/01-environment/) |
| 模块 2 | [快速入门](https://rexai.top/tutorials/go-adk/02-quickstart/) |
| ... | 持续更新中 |

---

*配套文章发布于公众号「梦兽编程」*