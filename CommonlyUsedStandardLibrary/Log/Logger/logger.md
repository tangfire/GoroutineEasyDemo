# 配置logger

默认情况下的logger只会提供日志的时间信息，但是很多情况下我们希望得到更多信息，比如记录该日志的文件名和行号等。log标准库中为我们提供了定制这些设置的方法。

log标准库中的Flags函数会返回标准logger的输出配置，而SetFlags函数用来设置标准logger的输出配置。

```go
    func Flags() int
    func SetFlags(flag int)
```

# flag选项

log标准库提供了如下的flag选项，它们是一系列定义好的常量。

```go
const (
    // 控制输出日志信息的细节，不能控制输出的顺序和格式。
    // 输出的日志在每一项后会有一个冒号分隔：例如2009/01/23 01:23:23.123123 /a/b/c/d.go:23: message
    Ldate         = 1 << iota     // 日期：2009/01/23
    Ltime                         // 时间：01:23:23
    Lmicroseconds                 // 微秒级别的时间：01:23:23.123123（用于增强Ltime位）
    Llongfile                     // 文件全路径名+行号： /a/b/c/d.go:23
    Lshortfile                    // 文件名+行号：d.go:23（会覆盖掉Llongfile）
    LUTC                          // 使用UTC时间
    LstdFlags     = Ldate | Ltime // 标准logger的初始值
) 
```

下面我们在记录日志之前先设置一下标准logger的输出选项如下：

```go
func main() {
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
    log.Println("这是一条很普通的日志。")
} 
```

编译执行后得到的输出结果如下：

```go
2019/10/11 14:05:17.494943 .../log_demo/main.go:11: 这是一条很普通的日志。
```

# 配置日志前缀

log标准库中还提供了关于日志信息前缀的两个方法：

```go
    func Prefix() string
    func SetPrefix(prefix string) 
```

其中Prefix函数用来查看标准logger的输出前缀，SetPrefix函数用来设置输出前缀。

```go
func main() {
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
    log.Println("这是一条很普通的日志。")
    log.SetPrefix("[pprof]")
    log.Println("这是一条很普通的日志。")
}
```

上面的代码输出如下：

```go
[pprof]2019/10/11 14:05:57.940542 .../log_demo/main.go:13: 这是一条很普通的日志。 
```

这样我们就能够在代码中为我们的日志信息添加指定的前缀，方便之后对日志信息进行检索和处理。

# 配置日志输出位置

```go
func SetOutput(w io.Writer)
```

SetOutput函数用来设置标准logger的输出目的地，默认是标准错误输出。

例如，下面的代码会把日志输出到同目录下的xx.log文件中。

```go
func main() {
    logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("open log file failed, err:", err)
        return
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
    log.Println("这是一条很普通的日志。")
    log.SetPrefix("[小王子]")
    log.Println("这是一条很普通的日志。")
} 

```

如果你要使用标准的logger，我们通常会把上面的配置操作写到init函数中。

```go
func init() {
    logFile, err := os.OpenFile("./xx.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
    if err != nil {
        fmt.Println("open log file failed, err:", err)
        return
    }
    log.SetOutput(logFile)
    log.SetFlags(log.Llongfile | log.Lmicroseconds | log.Ldate)
}
```

# 创建logger

log标准库中还提供了一个创建新logger对象的构造函数–New，支持我们创建自己的logger示例。New函数的签名如下：

```go
func New(out io.Writer, prefix string, flag int) *Logger
```


New创建一个Logger对象。其中，参数out设置日志信息写入的目的地。参数prefix会添加到生成的每一条日志前面。参数flag定义日志的属性（时间、文件等等）。

举个例子：

```go
func main() {
    logger := log.New(os.Stdout, "<New>", log.Lshortfile|log.Ldate|log.Ltime)
    logger.Println("这是自定义的logger记录的日志。")
}
```

将上面的代码编译执行之后，得到结果如下：

```go
<New>2019/10/11 14:06:51 main.go:34: 这是自定义的logger记录的日志。
```

## 总结:


Go内置的log库功能有限，例如无法满足记录不同级别日志的情况，我们在实际的项目中根据自己的需要选择使用第三方的日志库，如logrus、zap等。