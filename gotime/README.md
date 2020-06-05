# gotime

- [x] 为内置的time.Time 输出json时格式化
- [x] 更加友好的time formatter
    

# Usage

1. time.Time 时间格式化

```go
func main() {
    // for time.Time 
    t1 := time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local)

    // output: 2020-02-25
    fmt.Println(gotime.FormatDatetime(t1, "YYYY-MM-DD"))
    
    // output: 2020-02-25 10:12:09
    fmt.Println(gotime.FormatDatetime(t1, "YYYY-MM-DD HH:mm:ss"))
}

```


2. gotime.Time 时间格式化

```go
func main() {
    // for gotime.Time 
    t1 := gotime.Time(time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local))

    // output: 2020-02-25
    fmt.Println(t1.FormatX( "YYYY-MM-DD"))
    
    // output: 2020-02-25 10:12:09
    fmt.Println(t1.FormatX( "YYYY-MM-DD HH:mm:ss"))
}

```

3. json marshal

```go
func main() {
    type aa struct {
        StartTime Time `json:"start_time"`
    }
    
    a := &aa{}
    a.StartTime = Time(time.Date(2020, 02, 25, 10, 12, 9, 0, time.Local))
    data, _ := j.Marshal(a)

    // output: `{"start_time":"2020-02-25 10:12:09"}`
    fmt.Println(string(data))
}
```