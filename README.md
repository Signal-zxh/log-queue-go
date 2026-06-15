# Log Queue Go

一个基于 Go 和 Kafka 的日志队列示例项目。

## 功能特性

- Kafka 生产者：发送消息到指定主题
- Kafka 消费者：订阅并接收消息
- 完整的消息生产-消费流程演示

## 技术栈

- **语言**: Go 1.20+
- **消息队列**: Apache Kafka
- **客户端库**: `github.com/IBM/sarama`

## 快速开始

### 前置要求

- Go 1.20+
- Docker 和 Docker Compose

### 启动 Kafka

使用 Docker Compose 启动 Kafka 服务：

```bash
docker-compose up -d
```

### 安装依赖

```bash
go mod download
```

### 运行项目

```bash
go run main.go
```

## 项目结构

```
log-queue-go/
├── main.go          # 主程序入口
├── docker-compose.yml # Kafka 容器配置
├── go.mod           # Go 模块依赖
└── go.sum           # 依赖校验文件
```

## 代码说明

`main.go` 实现了以下功能：

1. **创建消费者**：连接 Kafka 并订阅 `test-topic` 主题
2. **创建生产者**：同步生产者发送消息
3. **发送消息**：向 `test-topic` 发送 "hello kafka"
4. **接收消息**：消费者监听并打印接收到的消息

## 使用说明

1. 确保 Kafka 服务已启动（使用 `docker-compose up -d`）
2. 运行 `go run main.go`
3. 程序会自动发送一条消息并接收打印

## 停止服务

```bash
docker-compose down
```
