"""Flask REST API for Red-Black Tree visualizer."""
from __future__ import annotations

import os
import uuid
from typing import Any, Dict

from flask import Flask, jsonify, request, send_from_directory

from rb_tree import RedBlackTree

app = Flask(__name__)

# 简化的 CORS 支持（不依赖 flask-cors）
# 前端直接从浏览器访问 Flask 服务，需要允许跨域 OPTIONS/POST/GET。
@app.after_request
def add_cors_headers(response):
    """为每个响应手动添加通用 CORS 头。"""
    response.headers["Access-Control-Allow-Origin"] = "*"
    response.headers["Access-Control-Allow-Methods"] = "GET, POST, OPTIONS"
    response.headers["Access-Control-Allow-Headers"] = "Content-Type"
    return response


# 全局状态集中管理，session_id 隔离多用户
session_store: Dict[str, RedBlackTree] = {}

# 节点数量上限，防止前端绘制过大的树导致性能下降
MAX_NODES = 100

# 静态资源目录：在 PyInstaller 环境中由 __file__ 所在目录决定
STATIC_DIR = os.path.join(os.path.dirname(os.path.abspath(__file__)), "static")


@app.route("/", methods=["GET", "OPTIONS"])
def index():
    """返回前端首页。"""
    if request.method == "OPTIONS":
        return jsonify({}), 204
    return send_from_directory(STATIC_DIR, "index.html")


@app.route("/<path:filename>", methods=["GET", "OPTIONS"])
def static_files(filename):
    """返回 static/ 目录下的静态文件。"""
    if request.method == "OPTIONS":
        return jsonify({}), 204
    return send_from_directory(STATIC_DIR, filename)


# --------------------------------------------------------------------------- #
# 统一响应
# --------------------------------------------------------------------------- #
def _success(data: Any = None) -> tuple:
    """组装成功响应，data 可为任意 JSON 可序列化对象。"""
    return jsonify({"success": True, "data": data}), 200


def _error(message: str, code: int = 400) -> tuple:
    """组装错误响应，默认 400 Bad Request，404 用于会话不存在。"""
    return jsonify({"success": False, "error": message}), code


# --------------------------------------------------------------------------- #
# API
# --------------------------------------------------------------------------- #
@app.route("/api/init", methods=["POST"])
def api_init():
    """
    创建新会话。

    返回：
    {
      "session_id": "uuid",
      "tree": null
    }
    """
    session_id = str(uuid.uuid4())
    session_store[session_id] = RedBlackTree()
    return _success({"session_id": session_id, "tree": None})


@app.route("/api/tree/<session_id>", methods=["GET"])
def api_get_tree(session_id: str):
    """
    获取指定会话的当前树快照及红黑树规则校验结果。

    返回：
    {
      "tree": { ... },  // 可能为 null
      "valid": { "valid": true, "messages": [...] }
    }
    """
    tree = session_store.get(session_id)
    if tree is None:
        return _error("会话不存在或已过期，请重新初始化", 404)
    return _success({"tree": tree.to_dict(), "valid": tree.validate()})


@app.route("/api/reset", methods=["POST"])
def api_reset():
    """
    清空指定会话中的树。

    请求体：{ "session_id": "..." }
    返回：{ "tree": null }
    """
    payload = request.get_json(silent=True) or {}
    session_id = payload.get("session_id")
    if not session_id:
        return _error("缺少 session_id")
    session_store[session_id] = RedBlackTree()
    return _success({"tree": None})


@app.route("/api/insert", methods=["POST"])
def api_insert():
    """
    在指定会话中插入一个节点。

    请求体：{ "session_id": "...", "key": 15 }
    返回：
    {
      "tree": { ... },
      "steps": [...],
      "valid": { ... },
      "history": [...]
    }
    """
    payload = request.get_json(silent=True) or {}
    session_id = payload.get("session_id")
    key = payload.get("key")

    if not session_id:
        return _error("缺少 session_id")
    tree = session_store.get(session_id)
    if tree is None:
        return _error("会话不存在或已过期，请重新初始化", 404)

    try:
        key = _parse_key(key)
    except ValueError as exc:
        return _error(str(exc))

    if tree.search(key):
        return _error(f"节点 {key} 已存在，请勿重复插入")

    existing = _count_nodes(tree.root)
    if existing >= MAX_NODES:
        return _error(f"节点数量已达上限 {MAX_NODES}，请删除部分节点后再插入")

    steps = tree.insert(key)
    if steps and steps[-1].type == "error":
        return _error(steps[-1].description)

    public_steps = [_step_to_dict(s) for s in steps]
    return _success({
        "tree": tree.to_dict(),
        "steps": public_steps,
        "valid": tree.validate(),
        "history": [h.to_dict() for h in tree.history],
    })


@app.route("/api/delete", methods=["POST"])
def api_delete():
    """
    在指定会话中删除一个节点。

    请求体：{ "session_id": "...", "key": 5 }
    返回结构与 /api/insert 相同。
    """
    payload = request.get_json(silent=True) or {}
    session_id = payload.get("session_id")
    key = payload.get("key")

    if not session_id:
        return _error("缺少 session_id")
    tree = session_store.get(session_id)
    if tree is None:
        return _error("会话不存在或已过期，请重新初始化", 404)

    try:
        key = _parse_key(key)
    except ValueError as exc:
        return _error(str(exc))

    if tree.search(key) is None:
        return _error(f"节点 {key} 不存在，无法删除")

    steps = tree.delete(key)
    if steps and steps[-1].type == "error":
        return _error(steps[-1].description)

    public_steps = [_step_to_dict(s) for s in steps]
    return _success({
        "tree": tree.to_dict(),
        "steps": public_steps,
        "valid": tree.validate(),
        "history": [h.to_dict() for h in tree.history],
    })


# --------------------------------------------------------------------------- #
# 工具函数
# --------------------------------------------------------------------------- #
def _parse_key(value: Any) -> int:
    """解析并校验输入键，返回整数或抛出 ValueError。"""
    if value is None or value == "":
        raise ValueError("请输入节点值")
    try:
        number = int(value)
    except (TypeError, ValueError):
        raise ValueError("节点值必须是整数")
    if number < -999999 or number > 999999:
        raise ValueError("节点值超出允许范围（-999999 ~ 999999）")
    return number


def _count_nodes(node) -> int:
    """递归统计以 node 为根的子树中的非 NIL 节点数量。"""
    if node is None or getattr(node, "key", None) is None:
        return 0
    return 1 + _count_nodes(node.left) + _count_nodes(node.right)


def _step_to_dict(step):
    """将后端 Step 对象转换为前端可消费的纯字典。"""
    return {
        "type": step.type,
        "node": step.node,
        "description": step.description,
        "tree_state": step.tree_state,
        "highlight_nodes": step.highlight_nodes,
        "reason": step.reason,
    }


if __name__ == "__main__":
    # 直接运行本文件时使用固定端口，便于开发调试
    app.run(host="127.0.0.1", port=5000, debug=True)
