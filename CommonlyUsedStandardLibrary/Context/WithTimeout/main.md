这段代码展示了 Go 语言中如何通过 `context.WithTimeout` 实现超时控制，并配合 `sync.WaitGroup` 管理 goroutine 的生命周期。以下是逐段解析与核心机制说明：

---

### **代码功能**
程序通过设置 50 毫秒的超时时间，演示了如何强制终止耗时过长的数据库连接操作。当超时触发时，子 goroutine 会立即退出，避免资源泄露。

---

### **逐行解析**
#### **1. 全局定义**
```go
var wg sync.WaitGroup
```
- **作用**：用于同步主 goroutine 和子 goroutine，确保主程序等待子任务完成后再退出。

#### **2. worker 函数**
```go
func worker(ctx context.Context) {
LOOP:
    for {
        fmt.Println("db connecting ...")
        time.Sleep(time.Millisecond * 10) // 模拟数据库连接耗时
        select {
        case <-ctx.Done(): // 监听超时或取消信号
            break LOOP
        default:
        }
    }
    fmt.Println("worker done!")
    wg.Done() // 通知 WaitGroup 任务完成
}
```
- **核心逻辑**：
    - 通过无限循环模拟数据库连接操作。
    - `time.Sleep(10ms)` 表示每次连接尝试的耗时。
    - `select` 监听 `ctx.Done()`：当超时（50ms）或手动调用 `cancel()` 时，触发退出循环。
    - 退出后调用 `wg.Done()` 通知主程序。

#### **3. main 函数**
```go
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    wg.Add(1)
    go worker(ctx)
    time.Sleep(5 * time.Second) // 主程序休眠 5 秒
    cancel() // 手动取消（此处实际不会生效，因已超时）
    wg.Wait()
    fmt.Println("over")
}
```
- **关键步骤**：
    1. **设置超时**：`context.WithTimeout` 创建 50ms 后自动取消的上下文。
    2. **启动子任务**：`go worker(ctx)` 开启 goroutine。
    3. **主程序休眠**：`time.Sleep(5秒)` 模拟主程序其他耗时操作（实际此处设计存在问题，见下文分析）。
    4. **手动取消**：`cancel()` 理论上用于主动终止，但此处因已超时，实际由超时机制触发退出。
    5. **等待同步**：`wg.Wait()` 确保子任务退出后主程序才结束。

---

### **核心机制**
#### **1. 超时自动取消**
- `context.WithTimeout` 会在 50ms 后自动关闭 `ctx.Done()` 通道，子 goroutine 通过监听此通道实现超时退出。
- **与手动取消的关系**：即使不调用 `cancel()`，超时仍会触发退出。但 `defer cancel()` 是推荐写法，确保资源释放。

#### **2. 代码设计问题**
- **主程序休眠过长**：`time.Sleep(5秒)` 导致主程序阻塞时间远超超时时间，但子任务已在 50ms 超时后退出，因此 `wg.Wait()` 会立即继续。
- **冗余的 cancel()**：由于超时机制已自动触发退出，此处 `cancel()` 调用实际是冗余的，但符合“始终清理 Context”的最佳实践。

#### **3. 执行流程**
1. 子 goroutine 启动后，每次循环尝试连接数据库（10ms/次）。
2. 50ms 后，`ctx.Done()` 被关闭，触发 `break LOOP`。
3. 子 goroutine 打印 "worker done!" 并调用 `wg.Done()`。
4. 主程序在 5 秒休眠结束后，执行 `cancel()` 和 `wg.Wait()`，最终输出 "over"。

---

### **应用场景**
此模式适用于：
1. **数据库/API 调用限时**：防止因网络问题导致无限等待。
2. **任务队列管理**：强制终止执行超时的任务。
3. **资源密集型操作**：如文件处理、图像渲染等需严格时间控制的场景。

---

### **改进建议**
1. **使用 `defer cancel()`**：在创建 Context 后立即添加 `defer cancel()`，确保即使提前退出也能释放资源。
2. **优化主程序逻辑**：移除不必要的 `time.Sleep(5秒)`，改用 `wg.Wait()` 直接等待。
3. **错误处理**：可通过 `ctx.Err()` 区分超时原因（`context.DeadlineExceeded` 或手动取消）。

改进后的 main 函数示例：
```go
func main() {
    ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
    defer cancel() // 确保资源释放
    wg.Add(1)
    go worker(ctx)
    wg.Wait() // 直接等待子任务完成
    fmt.Println("over")
}
```

---

通过这段代码，可以清晰理解 Go 语言如何通过 `context` 包实现精准的超时控制，避免因任务阻塞导致的资源泄露问题。