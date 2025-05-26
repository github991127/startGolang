<h1>
<p align="center">Go</p>
</h1>
<font color="GhostWhite">菖蒲</font>
<div STYLE="page-break-after: always;"></div>

# Go

8 小时转职 Golang 工程师（如果你想低成本学习 Go 语言）

**视频链接**：[8 小时转职 Golang 工程师](https://www.bilibili.com/video/BV1gf4y1r79E)

## 一、入门

### （一）配置

- **GOPATH 和 GOROOT**
  - `go env【cmd】`
  - `set GOPATH=C:\Users\wenjiabao\go`
  - `set GOROOT=C:\Program Files\Go`

### （二）go 特点

1. **部署简单**

   - 直译机器码
   - 不依赖三方库
   - 可直接运行

2. **底层并发**
3. **标准库优化**
   - runtime 调度
   - GC 垃圾回收

### （三）语法注意点

函数名和 `{` 在同一行；函数返回类型，可以实名，`return` 时不用传变量名；名字大写：`public`，小写：`private`。

### （四）执行过程

单包只能执行包内函数，多包需要跨包引用，执行步骤如下：

1. 去 `import` 引用的包体
2. `const`
3. `var`
4. 执行 `init()`
5. 执行 `main()`

递归进入包体直到栈底，依次栈调用执行 2 - 4，直到回退到 main，开始执行 2 - 5。

### （五）语法

- **指针** ：golang 只有值传递，没有引用传递；golang 指针拷贝（类似引用传递效果）的有：`slice`，`map`，`channel` 等等，其余为值拷贝。
- **defer** ：先 `return` 后 `defer`；`defer` 函数的执行顺序是先进后出的栈，即先注册的 `defer` 函数最后执行。
- **slice 切片** ：静态数组规定长度，值转递；`slice` 不规定长度，引用传递；`make()` 函数创建或初始化赋值创建；`slice` 是引用类型，为避免一损俱损，需要使用 `copy()` 函数进行复制改为值传递。
- **类的继承** ：

```go
type SuperMan struct {
    Man // 继承其实是利用组合实现的（匿名结构体属性）
    score int
}
```
