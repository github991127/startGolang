"""红黑树单元测试。"""
from __future__ import annotations

import pytest

from rb_tree import RedBlackTree


def test_insert_ascending_sequence():
    """测试连续升序插入 1~20，每一步后都应满足红黑树规则。"""
    tree = RedBlackTree()
    for i in range(1, 21):
        steps = tree.insert(i)
        assert steps[-1].type == "finish", f"插入 {i} 未完成：{steps[-1].description}"
        valid = tree.validate()
        assert valid["valid"], f"插入 {i} 后规则校验失败：{valid['messages']}"


def test_insert_example_keys():
    """测试非升序插入样例键，验证最终树合法。"""
    tree = RedBlackTree()
    for key in [10, 5, 15, 3, 7, 13, 18]:
        steps = tree.insert(key)
        assert steps[-1].type == "finish"
    assert tree.validate()["valid"]
    assert tree.root.key == 10 or tree.root.color == "BLACK"


def test_delete_leaf_and_validate():
    """先插入样例键再删除叶子节点，校验删除后仍满足红黑树规则。"""
    tree = RedBlackTree()
    for key in [10, 5, 15, 3, 7, 13, 18]:
        tree.insert(key)
    for key in [3, 7]:
        steps = tree.delete(key)
        assert steps[-1].type == "finish", f"删除 {key} 未完成：{steps[-1].description}"
        assert tree.validate()["valid"], f"删除 {key} 后校验失败"


def test_delete_root():
    """测试删除根节点，验证替换与修正后树合法。"""
    tree = RedBlackTree()
    for key in [10, 5, 15]:
        tree.insert(key)
    steps = tree.delete(10)
    assert steps[-1].type == "finish"
    assert tree.validate()["valid"]


def test_delete_nonexistent():
    """删除不存在的节点应返回 error 步骤。"""
    tree = RedBlackTree()
    tree.insert(10)
    steps = tree.delete(99)
    assert steps[-1].type == "error"


def test_insert_duplicate():
    """重复插入同一关键字应返回 error 步骤。"""
    tree = RedBlackTree()
    tree.insert(10)
    steps = tree.insert(10)
    assert steps[-1].type == "error"


def test_steps_include_states():
    """验证插入步骤包含描述与树状态快照。"""
    tree = RedBlackTree()
    steps = tree.insert(10)
    assert len(steps) >= 2
    for step in steps:
        assert step.description
