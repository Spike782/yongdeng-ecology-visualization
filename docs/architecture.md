

# 项目文件结构
```
yongdeng-ecology-visualization/
├── backend/ # 后端代码
│ ├── controllers/ # 控制器层（API 逻辑处理）
│ ├── models/ # 数据模型（数据库交互）
│ ├── routes/ # 路由（API 路由定义）
│ ├── services/ # 服务层（业务逻辑）
│ ├── config/ # 配置文件（数据库、环境变量等）
│ ├── utils/ # 工具类函数（如数据处理等）
│ ├── app.js # 后端入口文件
│ └── server.js # 后端启动文件
│
├── frontend/ # 前端代码
│
├── data/ # 数据文件夹
│ ├── raw/ # 原始数据（遥感影像、调查数据等）
│ ├── processed/ # 处理后的数据（例如，NDVI 数据、栅格数据）
│ └── analysis_results/ # 生态风险评估和土地利用分析结果
│
├── docs/ # 项目文档（包括设计文档、报告、API 文档等）
│ ├── api.md # API 文档
│ ├── architecture.md # 系统架构文档
│
├── .gitignore # Git 忽略文件
```



---
