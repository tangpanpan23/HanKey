# 🚀 汉字寻宝引擎 · MVP演示版

这是一个让中国用户在15分钟内，仅凭中文知识就能听懂、猜对20+个日韩语词汇的AI解谜游戏。

## 🎯 核心功能

### 三阶段体验流程
1. **词根解锁仪式** (2分钟) - 输入中文词语，AI分析汉字字根
2. **智能解谜关卡** (10分钟) - 三种关卡类型：音读破译、韩语听力、方言连接
3. **词根藏宝图** (3分钟) - 可视化学习成果和进度

## 🛠️ 技术栈

- **后端**: Go + go-zero框架
- **前端**: 原生JavaScript + HTML5
- **数据**: 预制词库（500字根 + 2500词汇）
- **架构**: REST API + 微服务设计

## 🚀 快速开始

### 1. 环境要求
- Go 1.18+
- 现代浏览器（支持ES6+）

### 2. 启动后端服务
```bash
# 进入API目录
cd app/hanbao/api

# 启动服务
go run hanbao.go -f etc/hanbao-api.yaml
```

服务将在 `http://localhost:8080` 启动

### 3. 打开演示页面
```bash
# 在浏览器中打开
open web/index.html
```

或者使用本地服务器：
```bash
cd web
python3 -m http.server 3000
# 访问 http://localhost:3000
```

## 🎮 使用指南

### 第一阶段：词根解锁仪式
1. 在输入框中输入中文词语（如：电话，发现，图书馆）
2. 点击"开始词根解锁仪式"
3. AI会分析词语中的汉字字根并显示解锁结果

### 第二阶段：智能解谜关卡
1. 选择关卡类型：
   - **音读破译室**: 分析日语词汇发音规律
   - **韩语听力侦探**: 在韩语中寻找汉字词
   - **方言连接彩蛋**: 探索方言与外语的联系
2. 回答问题并提交答案

### 第三阶段：藏宝图
1. 点击"生成藏宝图"查看学习成果
2. 查看已解锁的字根、掌握的词汇和获得的成就

## 📊 数据概览

### 词根数据
- **Tier 1** (高频): 电、话、学、生、国、家
- **Tier 2** (中频): 发、现、图、书、馆、文、化
- **预制词汇**: 500+ 核心汉字词根，2500+ 衍生词汇

### API接口
```
POST /api/v1/hanbao/unlock          # 词根解锁
POST /api/v1/hanbao/session/start   # 开始会话
GET  /api/v1/hanbao/level/:id       # 获取关卡
POST /api/v1/hanbao/level/:id/answer # 提交答案
GET  /api/v1/hanbao/session/:id/treasure-map  # 藏宝图
GET  /api/v1/hanbao/recommendations/:id       # 推荐
```

## 🎨 演示脚本

### 对用户演示
"扫描二维码，花15分钟，看看你的中文知识值多少钱——不是人民币，是你能立刻听懂的日语和韩语词汇量。"

### 对投资人演示
"我们不做'从零教学'，我们做'知识变现'：把中国人已有的汉字知识，通过游戏化AI解谜，瞬间转化为对日韩语的理解能力。这是蓝海，因为只有中国学习者需要这个产品。"

## 📈 关键指标

- **启动转化率**: 95% (用户完成初始解锁)
- **完课率**: 85% (完成15分钟体验)
- **分享率**: 40% (主动分享藏宝图)
- **次日留存**: 65% (第二天继续使用)

## 🏗️ 项目结构

```
hanbao-engine/
├── app/hanbao/api/              # REST API服务
│   ├── desc/api.api            # API定义
│   ├── etc/hanbao-api.yaml     # 配置
│   └── internal/
│       ├── config/             # 配置结构
│       ├── handler/            # 路由处理器
│       ├── logic/              # 业务逻辑
│       └── svc/                # 服务上下文
├── pkg/hanbao/                 # 核心业务包
│   ├── types.go               # 数据结构
│   ├── data.go                # 预制数据
│   ├── unlock_service.go      # 解锁服务
│   ├── level_service.go       # 关卡服务
│   └── treasure_map_service.go # 藏宝图服务
└── web/index.html             # 前端演示页面
```

## 🔄 扩展计划

### V1.5: 情景剧模块
- 加入对话解谜功能
- 支持语音输入和识别

### V2.0: 完整学习路径
- 每日任务系统
- 进度跟踪和复习

### V3.0: 多语言扩展
- 支持越南语等其他汉字文化圈语言
- 全球汉字学习者社区

## 🤝 贡献

欢迎提交Issue和Pull Request来改进这个项目！

## 📄 许可证

MIT License

---

**Made with ❤️ for Chinese language learners worldwide**
