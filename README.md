# 红黑树可视化教学应用

一个基于 Python + Flask + 原生 HTML/JS 的交互式红黑树教学演示工具。支持插入、删除操作的逐步可视化，可回退到任意历史节点并重播完整修正过程。

## 功能特性

- **插入演示**：单节点插入、一键插入示例序列，自动播放插入修正动画
- **删除演示**：节点删除及双黑修正过程可视化
- **操作历史**：右侧历史列表可点击任意一次操作，重新播放完整修正过程
- **步骤控制**：首页、上一步、下一步、末页、播放/暂停、重播
- **规则校验**：实时显示红黑树五条规则校验结果
- **变动原因**：每一步操作都附带教学级说明，解释为什么这样做
- **中文提示**：所有错误与状态提示均为中文

## 技术栈

- **后端**：Python 3.10+、Flask（无 flask-cors，手动设置 CORS 头）
- **数据结构**：自实现红黑树（插入、删除、旋转、变色、规则校验）
- **桌面窗口**：PyWebView
- **打包**：PyInstaller `--onedir`
- **前端**：原生 HTML5 + CSS3 + JavaScript，SVG 绘制树结构
- **操作系统**：Windows 10+

## 目录结构

```
Red-BlackTree/
├── app.py                 # PyWebView 桌面入口
├── web.py                 # 网页入口：自动寻找空闲端口并打开浏览器
├── api.py                 # Flask REST API
├── build.bat              # PyInstaller onedir 打包脚本
├── requirements.txt       # Python 依赖
├── README.md              # 本文件
├── Red-BlackTree.md       # 红黑树算法原理文档
├── static/                # 前端静态资源
│   ├── index.html
│   ├── css/style.css
│   └── js/app.js
├── rb_tree/               # 红黑树核心库
│   ├── __init__.py
│   └── tree.py
└── tests/                 # 单元测试
    ├── __init__.py
    └── test_tree.py
```

## 安装依赖

```bash
pip install -r requirements.txt
```

## 运行方式

### 网页版

```bash
python web.py
```

服务启动后会自动寻找空闲端口并打开默认浏览器。

### 桌面版

```bash
python app.py
```

使用 PyWebView 打开 1400×900 的独立窗口。

### 运行测试

```bash
python -m pytest tests/ -v
```

### 打包

```bash
build.bat
```

输出目录为 `dist/RedBlackTree/`，运行 `dist/RedBlackTree/RedBlackTree.exe`。打包使用 `--onedir` 模式，静态资源通过 `--add-data` 一并打包。

## API 接口

所有接口统一返回如下格式：

```json
{
  "success": true,
  "data": { ... },
  "error": "..."
}
```

预检请求 `OPTIONS` 会自动返回 204，前端无需额外处理 CORS。

| 接口 | 方法 | 说明 |
|------|------|------|
| `/api/init` | POST | 创建新会话，返回 `session_id` |
| `/api/tree/<session_id>` | GET | 获取当前树结构及规则校验结果 |
| `/api/reset` | POST | 清空指定会话的树 |
| `/api/insert` | POST | 插入节点，返回树快照、完整步骤与历史 |
| `/api/delete` | POST | 删除节点，返回树快照、完整步骤与历史 |

### 请求示例

```bash
curl -X POST http://127.0.0.1:5000/api/init -H "Content-Type: application/json" -d '{}'
```

## 设计原则

- **原数据不可修改**：核心算法通过深拷贝保护内部状态
- **全局状态集中**：后端 `session_store` 统一管理所有会话
- **不要硬编码**：关键参数（如节点上限）提取为常量
- **错误提示友好**：空值、非整数、重复节点、节点不存在、数量超限均有中文提示

## 开发注意事项

- CORS 处理方式：未引入 `flask-cors`，而是在 `api.py` 中通过 `after_request` 手动设置响应头。
- 静态资源：生产运行与 PyInstaller 打包均依赖 `api.py` 中的 `STATIC_DIR`；新增资源请一并放入 `build.bat` 的 `--add-data` 参数。
- 端口冲突：`web.py` 与 `app.py` 都会自动从 5000 开始寻找空闲端口。
- 节点数量限制：`api.py` 中 `MAX_NODES = 100`，防止过大树导致前端渲染异常。

## 红黑树参考资料

详见 [Red-BlackTree.md](Red-BlackTree.md)。
