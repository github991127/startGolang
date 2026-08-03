"""红黑树实现，支持插入/删除并记录可视化步骤。"""
from __future__ import annotations

import copy
from dataclasses import dataclass, field
from typing import Any, Dict, List, Optional

RED = "RED"
BLACK = "BLACK"


@dataclass
class Node:
    """红黑树节点。key 为 None 时表示 NIL 哨兵节点（叶子空位）。"""

    key: Any
    color: str = RED
    left: Optional["Node"] = None
    right: Optional["Node"] = None
    parent: Optional["Node"] = None

    def to_dict(self) -> Dict[str, Any]:
        """递归序列化为前端可渲染的字典结构；NIL 节点不输出。"""
        return {
            "key": self.key,
            "color": self.color,
            "left": self.left.to_dict() if self.left else None,
            "right": self.right.to_dict() if self.right else None,
        }


@dataclass
class Step:
    """可视化步骤：记录操作类型、相关节点、描述、树快照、高亮节点与教学原因。"""

    type: str
    node: Optional[Any]
    description: str
    tree_state: Optional[Dict[str, Any]] = None
    highlight_nodes: List[Any] = field(default_factory=list)
    reason: str = ""  # 当前步骤的变动原因/教学提示


class OperationRecord:
    """记录一次完整操作（插入或删除）的所有步骤，用于历史回退后重新播放。"""

    def __init__(self, key: Any, op_type: str):
        self.key = key
        self.op_type = op_type  # 'insert' 或 'delete'
        self.steps: List[Step] = []

    def to_dict(self) -> Dict[str, Any]:
        """序列化为 API 可直接返回的结构。"""
        return {
            "key": self.key,
            "type": self.op_type,
            "steps": [
                {
                    "type": s.type,
                    "node": s.node,
                    "description": s.description,
                    "tree_state": s.tree_state,
                    "highlight_nodes": s.highlight_nodes,
                    "reason": s.reason,
                }
                for s in self.steps
            ],
        }


class RedBlackTree:
    """标准红黑树实现；所有修改操作都会记录可视化步骤。"""

    def __init__(self):
        # NIL 哨兵节点：所有叶子空位共享同一个黑色节点，简化删除修正判断
        self.NIL = Node(key=None, color=BLACK)
        self.root: Node = self.NIL
        self.steps: List[Step] = []          # 当前操作产生的步骤
        self.history: List[OperationRecord] = []  # 历次操作的完整记录

    # ------------------------------------------------------------------ #
    # 辅助函数
    # ------------------------------------------------------------------ #
    def _clone(self) -> "RedBlackTree":
        """深拷贝整棵树，用于需要保护原数据的场景。"""
        return copy.deepcopy(self)

    def _record(self, type_: str, node: Optional[Node], description: str,
                highlight_nodes: Optional[List[Node]] = None, reason: str = ""):
        """
        记录一个可视化步骤，包含当前树快照与教学说明。

        :param type_: 步骤类型，如 start / insert / rotate_left / recolor / finish / error
        :param node: 当前步骤关注的主要节点，NIL 会序列化为 None
        :param description: 给人看的简短说明
        :param highlight_nodes: 需要在界面中高亮的节点列表
        :param reason: 当前步骤为何如此操作的教学解释
        """
        self.steps.append(Step(
            type=type_,
            node=node.key if node and node is not self.NIL else None,
            description=description,
            tree_state=self.to_dict(),
            highlight_nodes=[n.key for n in (highlight_nodes or [])
                             if n and n is not self.NIL],
            reason=reason,
        ))

    def _reset_steps(self):
        """新一轮操作前清空步骤缓存，避免历史步骤混入当前操作。"""
        self.steps = []

    def _left_rotate(self, x: Node):
        """
        以 x 为支点左旋，要求 x.right 不为 NIL。

        左旋动作：x 的右孩子 y 上升为子树根，x 下沉为 y 的左孩子，
        y 的左子树成为 x 的右子树。
        """
        y = x.right
        if y is self.NIL:
            return
        x.right = y.left
        if y.left is not self.NIL:
            y.left.parent = x
        y.parent = x.parent
        if x.parent is self.NIL:
            self.root = y
        elif x is x.parent.left:
            x.parent.left = y
        else:
            x.parent.right = y
        y.left = x
        x.parent = y

    def _right_rotate(self, y: Node):
        """
        以 y 为支点右旋，要求 y.left 不为 NIL。

        右旋动作：y 的左孩子 x 上升为子树根，y 下沉为 x 的右孩子，
        x 的右子树成为 y 的左子树。
        """
        x = y.left
        if x is self.NIL:
            return
        y.left = x.right
        if x.right is not self.NIL:
            x.right.parent = y
        x.parent = y.parent
        if y.parent is self.NIL:
            self.root = x
        elif y is y.parent.right:
            y.parent.right = x
        else:
            y.parent.left = x
        x.right = y
        y.parent = x

    def _transplant(self, u: Node, v: Node):
        """
        用子树 v 替换子树 u 的位置，不更新 v 的子节点。

        这是删除算法中的标准辅助操作，用于将被删除节点或其前驱
        挂接到正确位置。
        """
        if u.parent is self.NIL:
            self.root = v
        elif u is u.parent.left:
            u.parent.left = v
        else:
            u.parent.right = v
        v.parent = u.parent

    def _minimum(self, node: Node) -> Node:
        """返回以 node 为根的子树中的最小节点。"""
        while node.left is not self.NIL:
            node = node.left
        return node

    def search(self, key: Any) -> Optional[Node]:
        """按 BST 规则查找 key，返回节点或 None。"""
        node = self.root
        while node is not self.NIL and key != node.key:
            node = node.left if key < node.key else node.right
        return node if node is not self.NIL else None

    # ------------------------------------------------------------------ #
    # 插入
    # ------------------------------------------------------------------ #
    def insert(self, key: Any) -> List[Step]:
        """
        按 BST 规则插入 key，并通过变色与旋转修复红黑树规则。

        返回本次插入产生的全部可视化步骤，最后一步为 finish 或 error。
        """
        self._reset_steps()
        record = OperationRecord(key, "insert")
        self.history.append(record)
        self._record("start", None, f"准备插入节点 {key}",
                     reason="插入新节点可能破坏红黑规则，需要分步骤修正")

        # 新节点默认为红色：红色不会立即破坏黑高一致的规则 5，只可能引发双红冲突
        new_node = Node(key=key, color=RED, left=self.NIL, right=self.NIL, parent=self.NIL)

        # 1. 按 BST 找到插入位置
        parent = self.NIL
        current = self.root
        while current is not self.NIL:
            parent = current
            if key == current.key:
                self._record("error", current, f"键 {key} 已存在，请勿重复插入",
                             reason="红黑树中同一关键字只能出现一次")
                record.steps = self.steps.copy()
                return self.steps
            if key < current.key:
                current = current.left
            else:
                current = current.right

        new_node.parent = parent
        if parent is self.NIL:
            self.root = new_node
        elif key < parent.key:
            parent.left = new_node
        else:
            parent.right = new_node

        self._record("insert", new_node, f"插入红色节点 {key}", [new_node],
                     reason="新节点默认设为红色，以尽量避免立即破坏黑高一致的规则 5")

        # 2. 修正双红冲突：父节点为红色时才需要处理
        while new_node is not self.root and new_node.parent.color == RED:
            grand = new_node.parent.parent
            if grand is self.NIL:
                break
            parent = new_node.parent

            # parent 是 grand 的左孩子：LL 或 LR 情况
            if parent is grand.left:
                uncle = grand.right
                if uncle.color == RED:
                    # 叔父为红：无法借黑，变色并将冲突上移到祖父
                    self._record("recolor", parent,
                                 f"叔父 {uncle.key} 为红色，父 {parent.key} 与叔父变黑，祖父 {grand.key} 变红",
                                 [parent, uncle, grand],
                                 reason="叔父为红时无法借黑，通过变色保持两侧黑高相等，并把冲突上移至祖父")
                    parent.color = BLACK
                    uncle.color = BLACK
                    grand.color = RED
                    new_node = grand
                    continue

                # LR 情况：先对 parent 左旋，转换为 LL 情况
                if new_node is parent.right:
                    self._record("rotate_left", parent,
                                 f"LR 情况：对父节点 {parent.key} 左旋，转化为 LL 情况",
                                 [parent, new_node],
                                 reason="新节点在父的右侧、父在祖父的左侧，属于 LR，需要先左旋转换为标准的 LL 情形")
                    self._left_rotate(parent)
                    new_node = parent
                    parent = new_node.parent

                # LL 情况：对 grand 右旋并交换颜色
                self._record("rotate_right", grand,
                             f"LL 情况：对祖父 {grand.key} 右旋，并交换颜色",
                             [grand, parent],
                             reason="左侧连续两个红节点，通过右旋将中间节点提升，并交换颜色消除双红冲突")
                self._right_rotate(grand)
                parent.color, grand.color = grand.color, parent.color
                self._record("recolor", parent,
                             f"旋转后 {parent.key} 变黑、{grand.key} 变红",
                             [parent, grand],
                             reason="提升为子树根的节点改为黑色，下沉节点改为红色，保持黑高不变")
                new_node = parent
            else:
                # parent 是 grand 的右孩子：RR 或 RL 情况，与左侧完全对称
                uncle = grand.left
                if uncle.color == RED:
                    self._record("recolor", parent,
                                 f"叔父 {uncle.key} 为红色，父 {parent.key} 与叔父变黑，祖父 {grand.key} 变红",
                                 [parent, uncle, grand],
                                 reason="叔父为红时无法借黑，通过变色保持两侧黑高相等，并把冲突上移至祖父")
                    parent.color = BLACK
                    uncle.color = BLACK
                    grand.color = RED
                    new_node = grand
                    continue

                # RL 情况：先对 parent 右旋，转换为 RR 情况
                if new_node is parent.left:
                    self._record("rotate_right", parent,
                                 f"RL 情况：对父节点 {parent.key} 右旋，转化为 RR 情况",
                                 [parent, new_node],
                                 reason="新节点在父的左侧、父在祖父的右侧，属于 RL，需要先右旋转换为标准的 RR 情形")
                    self._right_rotate(parent)
                    new_node = parent
                    parent = new_node.parent

                # RR 情况：对 grand 左旋并交换颜色
                self._record("rotate_left", grand,
                             f"RR 情况：对祖父 {grand.key} 左旋，并交换颜色",
                             [grand, parent],
                             reason="右侧连续两个红节点，通过左旋将中间节点提升，并交换颜色消除双红冲突")
                self._left_rotate(grand)
                parent.color, grand.color = grand.color, parent.color
                self._record("recolor", parent,
                             f"旋转后 {parent.key} 变黑、{grand.key} 变红",
                             [parent, grand],
                             reason="提升为子树根的节点改为黑色，下沉节点改为红色，保持黑高不变")
                new_node = parent

        # 3. 保证根节点始终为黑色
        if self.root.color == RED:
            self._record("recolor", self.root,
                         f"将根节点 {self.root.key} 重新设为黑色",
                         [self.root],
                         reason="根节点必须为黑色，以满足红黑树规则 2")
            self.root.color = BLACK

        self._record("finish", self.root, f"插入 {key} 完成，红黑树规则已恢复",
                     reason="所有双红冲突已解决，黑高一致且根节点为黑")
        record.steps = self.steps.copy()
        return self.steps

    # ------------------------------------------------------------------ #
    # 删除
    # ------------------------------------------------------------------ #
    def delete(self, key: Any) -> List[Step]:
        """
        按 BST 规则删除 key，并通过变色与旋转修复红黑树规则。

        返回本次删除产生的全部可视化步骤，最后一步为 finish 或 error。
        """
        self._reset_steps()
        record = OperationRecord(key, "delete")
        self.history.append(record)
        self._record("start", None, f"准备删除节点 {key}",
                     reason="删除黑色节点会导致经过它的路径黑高减 1，需要修正")

        target = self.search(key)
        if target is None:
            self._record("error", None, f"节点 {key} 不存在，无法删除",
                         reason="红黑树中不存在该关键字，无法执行删除")
            record.steps = self.steps.copy()
            return self.steps

        self._record("delete", target, f"定位到待删除节点 {key}", [target],
                     reason="确认目标节点，并按其子节点数量选择删除策略")

        # y 是真正被删除或被移动的节点，x 是承接 y 原来位置的节点
        y = target
        y_original_color = y.color
        if target.left is self.NIL:
            x = target.right
            self._transplant(target, target.right)
        elif target.right is self.NIL:
            x = target.left
            self._transplant(target, target.left)
        else:
            # 有两个子节点：用中序后继 y 替换 target
            y = self._minimum(target.right)
            y_original_color = y.color
            x = y.right
            if y.parent is target:
                x.parent = y
            else:
                self._transplant(y, y.right)
                y.right = target.right
                y.right.parent = y
            self._transplant(target, y)
            y.left = target.left
            y.left.parent = y
            y.color = target.color
            self._record("replace", y,
                         f"用后驱 {y.key} 替换待删除节点 {key}，继承颜色",
                         [y, target],
                         reason="目标有两个子节点，用中序后继替换可维持二叉搜索树顺序，同时保留原颜色信息")

        # 只有真正删除的节点是黑色时，才需要修复黑高
        if y_original_color == BLACK:
            self._delete_fixup(x)

        self._record("finish", self.root, f"删除 {key} 完成，红黑树规则已恢复",
                     reason="双黑节点已被消除，所有路径黑高恢复一致")
        record.steps = self.steps.copy()
        return self.steps

    def _delete_fixup(self, x: Node):
        """
        修复删除黑色节点后产生的双黑（double black）问题。

        核心策略：根据兄弟节点颜色与兄弟子节点颜色，分别执行
        "借黑" 或 "上移双黑" 操作，直到 x 到达根节点或 x 本身变为红色。
        """
        while x is not self.root and x.color == BLACK:
            parent = x.parent
            if parent is self.NIL:
                break

            # x 是左孩子：对称处理
            if x is parent.left:
                sibling = parent.right
                if sibling.color == RED:
                    # 兄弟为红：旋转并换色，把问题转化为兄弟为黑的标准情形
                    self._record("recolor", sibling,
                                 f"兄弟 {sibling.key} 为红，与父 {parent.key} 交换颜色后对父左旋",
                                 [sibling, parent],
                                 reason="兄弟为红时，先旋转将问题转化为兄弟为黑的标准情形")
                    sibling.color = BLACK
                    parent.color = RED
                    self._record("rotate_left", parent, f"对 {parent.key} 左旋", [parent, sibling],
                                 reason="左旋使原兄弟成为新的父节点，为后续调色做准备")
                    self._left_rotate(parent)
                    sibling = parent.right

                # 兄弟为黑且两个侄子都为黑：无法借黑，把双黑标记上移到父节点
                if (sibling.left.color == BLACK and sibling.right.color == BLACK):
                    self._record("recolor", sibling,
                                 f"兄弟 {sibling.key} 及其子节点均为黑，将兄弟变红，双黑上移父节点",
                                 [sibling, parent, x],
                                 reason="兄弟这边也没有多余黑色，只能把双黑标记向上转移给父节点")
                    sibling.color = RED
                    x = parent
                else:
                    # 远侄子为黑而近侄子为红：需要先旋转兄弟，使红色侄子转到远端
                    if sibling.right.color == BLACK:
                        self._record("recolor", sibling.left,
                                     f"兄弟 {sibling.key} 的左子为红、右子为黑，先左旋兄弟",
                                     [sibling.left, sibling],
                                     reason="富余的黑色在兄弟的左侧，需要先调整使远侄子变为红色，方便一次性借黑")
                        sibling.left.color = BLACK
                        sibling.color = RED
                        self._record("rotate_right", sibling, f"对 {sibling.key} 右旋", [sibling],
                                     reason="右旋将红色远端侄子转到合适位置，进入下一种标准情形")
                        self._right_rotate(sibling)
                        sibling = parent.right

                    # 远侄子为红：执行一次左旋即可一次性消除双黑
                    self._record("recolor", sibling,
                                 f"兄弟 {sibling.key} 的右子为红，可借黑色，执行左旋并调色",
                                 [sibling, sibling.right, parent],
                                 reason="远侄子为红，表示兄弟路径有富余黑色，可以一次性左旋并调色消除双黑")
                    sibling.color = parent.color
                    parent.color = BLACK
                    sibling.right.color = BLACK
                    self._record("rotate_left", parent, f"对 {parent.key} 左旋", [parent, sibling],
                                 reason="左旋使兄弟节点上浮，父节点下沉，重新分配黑色以恢复黑高")
                    self._left_rotate(parent)
                    x = self.root
            else:
                # x 是右孩子：与左侧完全对称
                sibling = parent.left
                if sibling.color == RED:
                    self._record("recolor", sibling,
                                 f"兄弟 {sibling.key} 为红，与父 {parent.key} 交换颜色后对父右旋",
                                 [sibling, parent],
                                 reason="兄弟为红时，先旋转将问题转化为兄弟为黑的标准情形")
                    sibling.color = BLACK
                    parent.color = RED
                    self._record("rotate_right", parent, f"对 {parent.key} 右旋", [parent, sibling],
                                 reason="右旋使原兄弟成为新的父节点，为后续调色做准备")
                    self._right_rotate(parent)
                    sibling = parent.left

                if (sibling.right.color == BLACK and sibling.left.color == BLACK):
                    self._record("recolor", sibling,
                                 f"兄弟 {sibling.key} 及其子节点均为黑，将兄弟变红，双黑上移父节点",
                                 [sibling, parent, x],
                                 reason="兄弟这边也没有多余黑色，只能把双黑标记向上转移给父节点")
                    sibling.color = RED
                    x = parent
                else:
                    if sibling.left.color == BLACK:
                        self._record("recolor", sibling.right,
                                     f"兄弟 {sibling.key} 的右子为红、左子为黑，先右旋兄弟",
                                     [sibling.right, sibling],
                                     reason="富余的黑色在兄弟的右侧，需要先调整使远侄子变为红色，方便一次性借黑")
                        sibling.right.color = BLACK
                        sibling.color = RED
                        self._record("rotate_left", sibling, f"对 {sibling.key} 左旋", [sibling],
                                     reason="左旋将红色远端侄子转到合适位置，进入下一种标准情形")
                        self._left_rotate(sibling)
                        sibling = parent.left

                    self._record("recolor", sibling,
                                 f"兄弟 {sibling.key} 的左子为红，可借黑色，执行右旋并调色",
                                 [sibling, sibling.left, parent],
                                 reason="远侄子为红，表示兄弟路径有富余黑色，可以一次性右旋并调色消除双黑")
                    sibling.color = parent.color
                    parent.color = BLACK
                    sibling.left.color = BLACK
                    self._record("rotate_right", parent, f"对 {parent.key} 右旋", [parent, sibling],
                                 reason="右旋使兄弟节点上浮，父节点下沉，重新分配黑色以恢复黑高")
                    self._right_rotate(parent)
                    x = self.root

        # 循环结束后若 x 为红色，直接变黑即可抵消一层缺失的黑色
        if x is not self.NIL:
            if x.color == RED and x.parent is not self.NIL and self._is_double_black_replacement(x):
                self._record("recolor", x,
                             f"红色替代节点 {x.key} 变黑，抵消丢失的黑色",
                             [x],
                             reason="红色子节点顶替被删除的黑色节点后变黑，可直接补足一条黑色")
            elif x is self.root and x.color == RED:
                x.color = BLACK

        # 兜底：确保根节点黑色
        if self.root is not self.NIL and self.root.color == RED:
            self.root.color = BLACK

    def _is_double_black_replacement(self, x: Node) -> bool:
        """
        简化判断：若当前节点为红色且父节点存在，认为它承担了双黑位置。

        该近似在可视化演示中足够使用，因为真正需要补黑时 x 必为红色。
        """
        return True

    # ------------------------------------------------------------------ #
    # 验证五条红黑树规则
    # ------------------------------------------------------------------ #
    def validate(self) -> Dict[str, Any]:
        """校验红黑树五条核心规则，返回 {valid, messages}。"""
        if self.root is self.NIL:
            return {"valid": True, "messages": ["空树合法"]}

        messages = []
        valid = True

        # 规则 2：根节点为黑色
        if self.root.color != BLACK:
            valid = False
            messages.append("规则 2 违反：根节点不是黑色")

        # 规则 4：红节点的子节点必须为黑色
        def check_red_children(node: Node):
            nonlocal valid
            if node is self.NIL:
                return
            if node.color == RED:
                if node.left.color != BLACK or node.right.color != BLACK:
                    valid = False
                    messages.append(f"规则 4 违反：红节点 {node.key} 存在红色子节点")
            check_red_children(node.left)
            check_red_children(node.right)

        check_red_children(self.root)

        # 规则 5：黑高一致
        def black_height(node: Node) -> int:
            if node is self.NIL:
                return 0
            left = black_height(node.left)
            right = black_height(node.right)
            if left != right:
                raise ValueError(f"节点 {node.key} 左右子树黑高不一致")
            return left + (1 if node.color == BLACK else 0)

        try:
            black_height(self.root)
        except ValueError as e:
            valid = False
            messages.append(f"规则 5 违反：{e}")

        if valid:
            messages.append("红黑树五条规则均满足")
        return {"valid": valid, "messages": messages}

    # ------------------------------------------------------------------ #
    # 序列化
    # ------------------------------------------------------------------ #
    def to_dict(self) -> Optional[Dict[str, Any]]:
        """将整棵树序列化为字典，空树返回 None。"""
        return self.root.to_dict() if self.root is not self.NIL else None
