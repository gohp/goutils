# Goutils

Common, utility packages for Go


## Package

| usage | package | remark |
| :--- | :--- | :--- |
| 类型转换 | [convert](convert/convert.go) | 常用类型转换 |
| 颜色 | [color](color/README.md) | 多种颜色输出 |
| 文件操作 | [file](file/file.go) | 文件路径，判断，读写等 |
| 随机生成 | [rand](rand/rand.go) | 生成随机字符串, 随机数字 |
| 切片操作 | [slice](slice/slice.go) | slice 一些操作 |
| safemap | [safemap](safemap/safemap.go) | 线程安全的map |
| 集合 | [set](set/README.md) | 实现集合及其操作 |
| 排序 | [sort](sort/README.md) | 常见排序算法 |
| 常规判断 | [regular](regular/regular.go) | 银行卡，手机，邮箱，IP地址, 身份证判断 |
| HTTP操作 | [httplib](httplib/httplib.go) | http get 快捷操作 |
| hash操作 | [hash](hash/hash.go) | md5, sha1, sha256等哈希算法 |
| 阻塞操作 | [choke](choke/choke.go) | 实现阻塞程序 |
| 错误代码 | [ecode](ecode/ecode.go) | 提取自B站的错误封装 |
| 获取IP | [curip](curip/README.md) | 获取内网，外网IP |
| 时间操作 | [gotime](gotime/README.md) | 时间输出, json格式化等 |
| jwt | [jwt](jwt/jwt.go) | jwt编码解码 |
| 流量限制 | [ratelimit](ratelimit/README.md) | 服务流量限制 |
| 距离计算 | [geo](geo/geo.go) | 计算两个经纬度之前的距离 |
| 敏感词过滤 | [sensitive](sensitive/sensitive.go) | 过滤敏感词汇 |
| 行政区域查询 | [area](area/README.md) | 根据行政编码，区号，名称等查询行政区域信息 |
| websocket | [websocket](ws/README.md) | 基于gorilla/websocket的服务端，客户端 |
| 缓存 | [cache](cache/README.md) | 两级缓存 |
| 字符串操作 | [str](str/README.md) | 两级缓存 |

## Usage

install 

```shell script
go get -u github.com/gohp/goutils
```
