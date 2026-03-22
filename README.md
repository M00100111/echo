# 基于Eino的个人知识库RAG项目

本项目是一个基于[Eino](https://github.com/cloudwego/eino)框架构建的智能AI代理系统，采用ReAct模式实现检索增强生成（RAG）功能，支持个人知识库内容检索、联网搜索和笔记迭代等核心能力。

## 核心特性

### 1. ReAct代理模式
系统采用经典的[ReAct (Reasoning + Acting)](https://arxiv.org/abs/2210.03629)框架，通过推理与行动的循环机制实现复杂任务处理。代理在每个步骤中都会进行思考并决定采取何种行动，直到完成最终目标。

- **最大步骤数**: 25步
- **核心实现**: `api/internal/ai/agent/reAct.go`
- **工作流程**: 观察 → 思考 → 行动 → 反馈 → 迭代

### 2. 个人知识库内容检索
系统集成了RAG技术，能够从个人知识库中检索相关信息并结合上下文生成回答。

- **检索工具**: `api/internal/ai/tools/query_internal_docs.go`
- **向量数据库**: Milvus（部署于Docker环境中）
- **检索流程**: 用户查询 → 向量化 → 相似性搜索 → 结果排序 → 内容返回

### 3. 联网搜索能力
除了本地知识库，系统还具备实时联网搜索能力，确保信息的时效性和完整性。

### 4. 笔记迭代功能
系统支持对已有笔记内容进行修改和更新，实现知识的持续演进。

- **功能实现**: `api/internal/ai/tools/modify_md_content.go`
- **操作类型**: 新增、编辑、删除Markdown内容
- **应用场景**: 知识更新、错误修正、内容优化

## 技术架构

### 前端组件
- **框架**: Eino AI框架
- **语言**: Go

### 后端服务
- **向量数据库**: Milvus
- **对象存储**: MinIO
- **配置管理**: etcd

### 部署环境
使用Docker Compose进行容器化部署，包含以下服务：
- Milvus向量数据库
- MinIO对象存储
- etcd配置中心

部署文件: `deploy/docker-compose.yaml`

## 快速开始

```bash
# 启动基础设施
docker-compose -f deploy/docker-compose.yaml up -d

# 运行应用
go run api/agent.go
```

## 配置说明

主要配置文件位于 `api/etc/agent.yaml`，包含模型参数、工具配置和系统设置等。

## 工具集成

系统支持多种工具函数，通过function calling机制动态调用：
- `query_internal_docs`: 查询内部文档
- `modify_md_content`: 修改Markdown内容
- `get_current_time`: 获取当前时间
- `mysql_crud`: MySQL数据库操作

## 开发指南

### 添加新工具
1. 在`api/internal/ai/tools/`目录下创建新的工具文件
2. 实现Tool接口
3. 在agent配置中注册新工具

### 自定义行为
通过修改prompt模板可以调整代理的行为模式，模板文件位于`api/internal/ai/template/template.go`。

## 使用场景

- 个人知识管理
- 智能问答系统
- 文档自动化处理
- 研究辅助工具

## 许可证

MIT License