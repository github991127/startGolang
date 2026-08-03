"""桌面入口：PyWebView 窗口加载本地 Flask 服务。"""
from __future__ import annotations

import os
import socket
import sys
import threading

import webview

from api import app
from web import find_free_port


def resource_path(relative_path: str) -> str:
    """兼容 PyInstaller 的资源路径。

    打包后资源被释放到临时目录 sys._MEIPASS；
    开发/源码运行时则返回脚本所在目录的相对路径。
    """
    if getattr(sys, "_MEIPASS", None):
        return os.path.join(sys._MEIPASS, relative_path)
    return os.path.join(os.path.dirname(os.path.abspath(__file__)), relative_path)


def start_server(port: int):
    """在独立线程中启动 Flask 服务供 PyWebView 加载。"""
    app.run(host="127.0.0.1", port=port, debug=False, use_reloader=False)


def main():
    """主流程：找端口 -> 起服务 -> 等就绪 -> 创建窗口并启动 GUI 消息循环。"""
    port = find_free_port()
    url = f"http://127.0.0.1:{port}"

    server_thread = threading.Thread(target=start_server, args=(port,), daemon=True)
    server_thread.start()

    # 等待服务就绪，避免窗口打开后页面无法访问
    import socket
    for _ in range(50):
        try:
            with socket.create_connection(("127.0.0.1", port), timeout=0.2):
                break
        except OSError:
            pass

    window = webview.create_window(
        "红黑树可视化演示",
        url,
        width=1400,
        height=900,
        min_size=(900, 600),
    )
    webview.start(debug=False)


if __name__ == "__main__":
    main()
