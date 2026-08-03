"use strict";

/**
 * 红黑树可视化前端模块
 *
 * 负责：会话初始化、API 调用、步骤/历史渲染、SVG 绘制、播放控制。
 */

const API_BASE = "";
const EXAMPLE_KEYS = [10, 5, 15, 3, 7, 13, 18];

let sessionId = null;
let steps = [];            // 当前操作的步骤数组
let currentStepIndex = -1;
let isPlaying = false;
let playTimer = null;

// 后端返回的完整操作历史：每个操作包含其所有步骤
let backendHistory = [];

const el = (id) => document.getElementById(id);
const svg = el("tree-svg");

// 前端内部状态：用于渲染、校验结果缓存
const state = {
  treesByStep: [],
  validity: null,
};

// ---------------------------------------------------------------------------
// API 封装
// ---------------------------------------------------------------------------

/**
 * 统一调用后端 API。
 * @param {string} path - 接口路径，如 /api/insert
 * @param {object|null} body - POST 请求体，GET 时传 null
 */
async function api(path, body = null) {
  const opts = { method: body ? "POST" : "GET", headers: {} };
  if (body) {
    opts.headers["Content-Type"] = "application/json";
    opts.body = JSON.stringify(body);
  }
  const res = await fetch(`${API_BASE}${path}`, opts);
  const json = await res.json();
  if (!json.success) throw new Error(json.error || "请求失败");
  return json.data;
}

// ---------------------------------------------------------------------------
// UI 小工具
// ---------------------------------------------------------------------------

/** 在底部弹出提示，type 仅用于样式分类（当前未区分）。 */
function showToast(message, type = "info") {
  const toast = el("toast");
  toast.textContent = message;
  toast.classList.add("show");
  setTimeout(() => toast.classList.remove("show"), 2500);
}

/** 更新右侧面板的规则校验显示。 */
function setValidity(valid, messages = []) {
  const box = el("validity");
  const text = messages.join("；") || (valid ? "红黑树规则校验通过" : "校验未通过");
  box.textContent = text;
  box.className = `validity ${valid ? "ok" : "bad"}`;
}

// ---------------------------------------------------------------------------
// 会话与操作
// ---------------------------------------------------------------------------

/** 初始化新会话，获取 session_id。 */
async function initSession() {
  const data = await api("/api/init", {});
  sessionId = data.session_id;
  state.treesByStep = [null];
  state.validity = { valid: true, messages: ["空树合法"] };
  steps = [];
  currentStepIndex = -1;
  renderSteps();
  drawTree(null);
  setValidity(true);
  enableControls(true);
}

/** 确保 sessionId 存在，否则弹出提示。 */
function ensureSessionId() {
  if (!sessionId) {
    showToast("会话尚未初始化，请稍候或刷新页面", "error");
    return false;
  }
  return true;
}

/** 插入节点并自动播放其修正过程。 */
async function insertKey(key) {
  if (key === undefined || key === "") return;
  if (!ensureSessionId()) return;
  const data = await api("/api/insert", { session_id: sessionId, key });
  backendHistory = data.history || [];
  loadSteps(data.steps, data.valid);
  await autoPlaySteps();
  renderHistory();
  showToast(`已插入节点 ${key}`);
}

/** 删除节点并自动播放其修正过程。 */
async function deleteKey(key) {
  if (key === undefined || key === "") return;
  if (!ensureSessionId()) return;
  const data = await api("/api/delete", { session_id: sessionId, key });
  backendHistory = data.history || [];
  loadSteps(data.steps, data.valid);
  await autoPlaySteps();
  renderHistory();
  showToast(`已删除节点 ${key}`);
}

// ---------------------------------------------------------------------------
// 操作历史
// ---------------------------------------------------------------------------

/** 渲染右侧操作历史标签。 */
function renderHistory() {
  const container = el("history-list");
  if (!container) return;
  container.innerHTML = "";
  backendHistory.forEach((item, idx) => {
    const btn = document.createElement("div");
    const label = item.type === "insert" ? `插入 ${item.key}` : `删除 ${item.key}`;
    btn.className = `history-item ${item.type}`;
    btn.textContent = `${idx + 1}. ${label}`;
    btn.title = `重新播放${label}的完整修正过程`;
    btn.addEventListener("click", () => replayOperationHistory(idx));
    container.appendChild(btn);
  });
}

/** 点击历史标签时，加载该操作的全部步骤并播放。 */
function replayOperationHistory(index) {
  if (index < 0 || index >= backendHistory.length) return;
  const item = backendHistory[index];
  stopPlay();
  loadSteps(item.steps || [], item.validity || state.validity);
  autoPlaySteps();
}

/** 清空当前会话的树。 */
async function resetTree() {
  if (!ensureSessionId()) return;
  await api("/api/reset", { session_id: sessionId });
  state.treesByStep = [null];
  backendHistory = [];
  steps = [];
  currentStepIndex = -1;
  renderSteps();
  renderHistory();
  drawTree(null);
  setValidity(true);
  showToast("已清空红黑树");
}

// ---------------------------------------------------------------------------
// 步骤加载与播放控制
// ---------------------------------------------------------------------------

/** 启用/禁用核心控制按钮，未初始化时禁用。 */
function enableControls(enabled) {
  [
    "btn-insert",
    "btn-insert-example",
    "btn-delete",
    "btn-reset",
    "btn-first",
    "btn-prev",
    "btn-next",
    "btn-last",
    "btn-play",
    "btn-replay",
  ].forEach((id) => {
    el(id).disabled = !enabled;
  });
}

/** 加载新步骤并准备播放。 */
function loadSteps(newSteps, validity) {
  steps = newSteps;
  state.validity = validity;
  currentStepIndex = -1;
  state.treesByStep = [null];
  for (const step of steps) {
    state.treesByStep.push(step.tree_state);
  }
  renderSteps();
}

/** 自动从第 0 步播放到最后一步；返回 Promise，播完后 resolve。 */
async function autoPlaySteps() {
  return new Promise((resolve) => {
    if (!steps.length) {
      resolve();
      return;
    }
    jumpToStep(0);
    const playFromStart = () => {
      playDirection(1);
      // 覆盖停止逻辑以触发 resolve
      const originalStop = stopPlay;
      stopPlay = () => {
        clearInterval(playTimer);
        isPlaying = false;
        el("btn-play").textContent = "播放";
        stopPlay = originalStop;
        resolve();
      };
    };
    playFromStart();
  });
}

/** 渲染步骤列表到右侧面板。 */
function renderSteps() {
  const container = el("step-list");
  container.innerHTML = "";
  steps.forEach((step, idx) => {
    const item = document.createElement("div");
    item.className = "step-item";
    item.textContent = `${idx + 1}. ${step.description}`;
    item.dataset.index = idx;
    item.addEventListener("click", () => jumpToStep(idx));
    container.appendChild(item);
  });
}

/** 跳转到指定步骤并更新界面。 */
function jumpToStep(index) {
  if (index < 0 || index >= steps.length) return;
  currentStepIndex = index;
  const step = steps[index];
  drawTree(step.tree_state, step.highlight_nodes);
  updateReason(step.reason);

  // 高亮步骤列表并滚动到可视区
  const items = document.querySelectorAll(".step-item");
  let activeItem = null;
  items.forEach((item, idx) => {
    const isActive = idx === index;
    item.classList.toggle("active", isActive);
    if (isActive) activeItem = item;
  });
  if (activeItem) activeItem.scrollIntoView({ block: "nearest" });

  updatePlayerButtons();

  if (state.validity) setValidity(state.validity.valid, state.validity.messages);
}

/** 根据当前步骤索引更新播放器按钮的 disabled 状态。 */
function updatePlayerButtons() {
  const hasSteps = steps.length > 0;
  el("btn-first").disabled = !hasSteps || currentStepIndex <= 0;
  el("btn-prev").disabled = !hasSteps || currentStepIndex <= 0;
  el("btn-next").disabled = !hasSteps || currentStepIndex >= steps.length - 1;
  el("btn-last").disabled = !hasSteps || currentStepIndex >= steps.length - 1;
}

/** 更新“当前变动原因”面板内容。 */
function updateReason(reason) {
  const box = el("reason-text");
  if (!box) return;
  box.textContent = reason || "当前步骤无额外说明";
}

/** 按指定方向（+1 或 -1）自动播放步骤。 */
function playDirection(direction) {
  clearInterval(playTimer);
  isPlaying = true;
  el("btn-play").textContent = "暂停";
  playTimer = setInterval(() => {
    const next = currentStepIndex + direction;
    if (next < 0 || next >= steps.length) {
      stopPlay();
      return;
    }
    jumpToStep(next);
  }, 900);
}

/** 停止自动播放。 */
function stopPlay() {
  clearInterval(playTimer);
  isPlaying = false;
  el("btn-play").textContent = "播放";
}

/** 切换播放/暂停状态；到达末尾时反向播放。 */
function togglePlay() {
  if (isPlaying) {
    stopPlay();
    return;
  }
  const direction = currentStepIndex >= steps.length - 1 ? -1 : 1;
  playDirection(direction);
}

// ---------------------------------------------------------------------------
// SVG 绘制
// ---------------------------------------------------------------------------

/**
 * 根据树数据绘制 SVG。
 * @param {object|null} treeData - 后端返回的树字典
 * @param {Array} highlightKeys - 需要高亮的节点键列表
 */
function drawTree(treeData, highlightKeys = []) {
  const highlightSet = new Set((highlightKeys || []).filter((k) => k !== null && k !== undefined));
  const width = svg.clientWidth || 1200;
  const height = 540;
  svg.setAttribute("viewBox", `0 0 ${width} ${height}`);
  svg.innerHTML = "";

  if (!treeData) {
    const text = document.createElementNS("http://www.w3.org/2000/svg", "text");
    text.setAttribute("x", width / 2);
    text.setAttribute("y", height / 2);
    text.setAttribute("text-anchor", "middle");
    text.setAttribute("fill", "#9ca3af");
    text.textContent = "当前为空树，请先插入节点";
    svg.appendChild(text);
    return;
  }

  // 计算节点位置：按深度分层，横向等分当前子树宽度
  const positions = new Map();
  const nodeRadius = 22;
  const levelHeight = 75;

  function assign(node, depth, left, right) {
    if (!node || node.key === null || node.key === undefined) return;
    const x = (left + right) / 2;
    const y = 50 + depth * levelHeight;
    positions.set(node.key, { x, y, node });
    assign(node.left, depth + 1, left, x);
    assign(node.right, depth + 1, x, right);
  }
  assign(treeData, 0, nodeRadius, width - nodeRadius);

  // 先画边（线和圆无重叠时更自然）
  positions.forEach(({ x, y, node }) => {
    [node.left, node.right].forEach((child) => {
      if (!child) return;
      const pos = positions.get(child.key);
      if (!pos) return;
      const line = document.createElementNS("http://www.w3.org/2000/svg", "line");
      line.setAttribute("x1", x);
      line.setAttribute("y1", y);
      line.setAttribute("x2", pos.x);
      line.setAttribute("y2", pos.y);
      line.setAttribute("class", "edge");
      svg.appendChild(line);
    });
  });

  // 再画节点
  positions.forEach(({ x, y, node }) => {
    const g = document.createElementNS("http://www.w3.org/2000/svg", "g");
    g.setAttribute("class", `node ${highlightSet.has(node.key) ? "node-highlight" : ""}`);
    g.setAttribute("transform", `translate(${x}, ${y})`);

    const circle = document.createElementNS("http://www.w3.org/2000/svg", "circle");
    circle.setAttribute("r", nodeRadius);
    circle.setAttribute("class", `node-circle ${node.color === "RED" ? "node-red" : "node-black"}`);
    g.appendChild(circle);

    const text = document.createElementNS("http://www.w3.org/2000/svg", "text");
    text.setAttribute("class", "node-text");
    text.textContent = node.key;
    g.appendChild(text);

    svg.appendChild(g);
  });
}

// ---------------------------------------------------------------------------
// 事件绑定
// ---------------------------------------------------------------------------

// 切换 Tab
document.querySelectorAll(".tab").forEach((tab) => {
  tab.addEventListener("click", () => {
    document.querySelectorAll(".tab").forEach((t) => t.classList.remove("active"));
    document.querySelectorAll(".panel").forEach((p) => p.classList.remove("active"));
    tab.classList.add("active");
    el(`panel-${tab.dataset.tab}`).classList.add("active");
  });
});

// 插入
document.getElementById("btn-insert").addEventListener("click", async () => {
  const value = el("insert-value").value;
  try { await insertKey(value); }
  catch (e) { showToast(e.message, "error"); }
});

// 插入示例序列
document.getElementById("btn-insert-example").addEventListener("click", async () => {
  for (const key of EXAMPLE_KEYS) {
    try {
      el("insert-value").value = key;
      await insertKey(key);
    } catch (e) { showToast(e.message, "error"); }
  }
});

// 删除
document.getElementById("btn-delete").addEventListener("click", async () => {
  const value = el("delete-value").value;
  try { await deleteKey(value); }
  catch (e) { showToast(e.message, "error"); }
});

// 清空
document.getElementById("btn-reset").addEventListener("click", resetTree);

// 播放器按钮
document.getElementById("btn-first").addEventListener("click", () => jumpToStep(0));
document.getElementById("btn-prev").addEventListener("click", () => jumpToStep(currentStepIndex - 1));
document.getElementById("btn-next").addEventListener("click", () => jumpToStep(currentStepIndex + 1));
document.getElementById("btn-last").addEventListener("click", () => jumpToStep(steps.length - 1));
document.getElementById("btn-play").addEventListener("click", togglePlay);
document.getElementById("btn-replay").addEventListener("click", () => {
  jumpToStep(0);
  playDirection(1);
});

// 窗口大小变化时重绘，避免 SVG 比例失真
window.addEventListener("resize", () => {
  const step = steps[currentStepIndex];
  if (step) drawTree(step.tree_state, step.highlight_nodes);
});

// 启动应用
initSession().catch((e) => showToast(e.message, "error"));
