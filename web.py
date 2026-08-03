"""网页入口：自动寻找空闲端口并打开浏览器。"""
from __future__ import annotations

import socket
import threading
import time
import webbrowser

from api import app


def find_free_port(start: int = 5000, end: int = 65535) -> int:
    """从 start 开始扫描，返回第一个未被占用的 TCP 端口。"""
    for port in range(start, end + 1):
        with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sock:
            sock.settimeout(0.1)
            if sock.connect_ex(("127.0.0.1", port)) != 0:
                return port
    raise RuntimeError("未找到可用端口")


def open_browser(url: str, delay: float = 1.0):
    """延迟 delay 秒后自动打开默认浏览器，避免服务尚未就绪。"""
    def _open():
        time.sleep(delay)
        webbrowser.open(url)
    threading.Thread(target=_open, daemon=True).start()


if __name__ == "__main__":
    port = find_free_port()
    url = f"http://127.0.0.1:{port}"
    print(f"启动网页版红黑树演示：{url}")
    open_browser(url)
    app.run(host="127.0.0.1", port=port, debug=False)
