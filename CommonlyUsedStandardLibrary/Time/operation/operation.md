# 时间操作

## Add

我们在日常的编码过程中可能会遇到要求时间+时间间隔的需求，Go语言的时间对象有提供Add方法如下：

```go
func (t Time) Add(d Duration) Time
```

举个例子，求一个小时之后的时间：

```go
func main() {
    now := time.Now()
    later := now.Add(time.Hour) // 当前时间加1小时后的时间
    fmt.Println(later)
} 
```

## Sub

求两个时间之间的差值：


```go
func (t Time) Sub(u Time) Duration 
```


返回一个时间段t-u。如果结果超出了Duration可以表示的最大值/最小值，将返回最大值/最小值。要获取时间点t-d（d为Duration），可以使用t.Add(-d)。

# Equal


```go
func (t Time) Equal(u Time) bool

```

判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较。本方法和用t==u不同，这种方法还会比较地点和时区信息。

# Before


```go
func (t Time) Before(u Time) bool

```

如果t代表的时间点在u之前，返回真；否则返回假。

# After

```go
func (t Time) After(u Time) bool 
```

如果t代表的时间点在u之后，返回真；否则返回假。