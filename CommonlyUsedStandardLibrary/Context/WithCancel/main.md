这段代码展示了 Go 语言中如何利用 `context` 包优雅地控制 goroutine 的生命周期，防止 goroutine 泄露。以下是逐段解析：

---

### **代码结构解析**
#### 1. **`gen` 函数**
```go
func gen(ctx context.Context) <-chan int {
    dst := make(chan int)
    n := 1
    go func() {
        for {
            select {
            case <-ctx.Done(): // 监听 Context 的取消信号
                return // 收到信号后退出 goroutine
            case dst <- n: // 向通道发送递增的整数
                n++
            }
        }
    }()
    return dst
}
```
- **功能**：生成一个整数序列的通道，持续发送 `1, 2, 3...`。
- **实现逻辑**：
    - 启动一个匿名 goroutine，通过 `select` 监听两个事件：
        1. **`ctx.Done()`**：当 Context 被取消时（例如调用 `cancel()`），关闭通道并退出 goroutine。
        2. **`dst <- n`**：向通道发送当前整数 `n`，然后递增。
    - **关键设计**：通过 `ctx.Done()` 监听取消信号，确保在父函数（如 `main`）不再需要数据时，goroutine 能及时终止，避免资源泄露。

#### 2. **`main` 函数**
```go
func main() {
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel() // 确保最终调用 cancel()

    for n := range gen(ctx) {
        fmt.Println(n)
        if n == 5 {
            break // 当 n=5 时终止循环
        }
    }
}
```
- **步骤解析**：
    1. **创建可取消的 Context**：`context.WithCancel` 返回一个可取消的上下文 `ctx` 和取消函数 `cancel`。
    2. **`defer cancel()`**：在 `main` 函数结束时自动调用 `cancel()`，触发 `ctx.Done()` 信号，通知所有依赖此 Context 的 goroutine 终止。
    3. **循环读取通道**：通过 `for range` 从 `gen(ctx)` 生成的通道中读取整数，直到 `n=5` 时跳出循环。
    4. **终止条件**：当 `n=5` 时，`break` 结束循环，`main` 函数退出，`defer cancel()` 执行，最终关闭 `gen` 中的 goroutine。

---

### **核心机制**
#### 1. **Context 的取消机制**
- **作用**：通过 `context.WithCancel` 创建父子 Context 的关联关系。调用 `cancel()` 会关闭 `ctx.Done()` 通道，所有监听此通道的代码（如 `gen` 中的 `select`）会立即响应。
- **优势**：避免 goroutine 泄露，确保资源及时释放。

#### 2. **通道与 select 的协作**
- **通道 `dst`**：用于在 goroutine 和 `main` 函数之间传递数据。
- **select 多路复用**：同时监听数据发送 (`dst <- n`) 和取消信号 (`<-ctx.Done()`)，确保两者互不阻塞。

---

### **执行流程**
1. `main` 调用 `gen(ctx)` 启动生成器 goroutine。
2. goroutine 持续发送 `1, 2, 3...` 到通道 `dst`。
3. `main` 循环读取这些值并打印，直到 `n=5` 时跳出循环。
4. `main` 函数结束前，`defer cancel()` 触发 `ctx.Done()` 信号。
5. `gen` 中的 goroutine 收到信号，通过 `return` 终止，通道关闭。

---

### **实际应用场景**
此模式适用于以下场景：
1. **生成器模式**：需要持续生成数据，但需在特定条件下终止。
2. **超时控制**：若需在指定时间内完成任务，可结合 `context.WithTimeout`。
3. **资源清理**：确保后台任务不会因主程序提前退出而泄露。

---

### **最佳实践**
1. **始终使用 `defer cancel()`**：确保无论函数如何退出（正常或异常），Context 都会被取消。
2. **避免传递 nil Context**：优先使用 `context.Background()` 或 `context.TODO()` 作为根 Context。
3. **仅传递必要数据**：若需跨 goroutine 传值，使用 `context.WithValue`，但避免滥用。

通过这段代码，可以清晰看到 Go 语言如何通过 Context 实现优雅的并发控制。