# Goutils

Common, utility packages for Go


## Package

| usage | package | remark |
| :--- | :--- | :--- |
| 类型转换 | [convert](convert/convert.go) | 常用类型转换 |
| 文件操作 | [file](file/file.go) | 文件路径，判断，读写等 |
| 随机生成 | [rand](rand/rand.go) | 生成随机字符串 |
| 切片操作 | [slice](slice/slice.go) | slice 一些操作 |
| safemap | [safemap](safemap/safemap.go) | 线程安全的map |
| 常规判断 | [regular](regular/regular.go) | 银行卡，手机，邮箱判断 |
| HTTP操作 | [http](http/http.go) | http get 快捷操作 |
| hash操作 | [hash](hash/hash.go) | md5, sha1, sha256等哈希算法 |
| 阻塞操作 | [choke](choke/choke.go) | 实现阻塞程序 |
| 错误代码 | [ecode](ecode/ecode.go) | 提取自B站的错误封装 |
| 获取IP | [curip](curip/curip.go) | 获取内网，外网IP |
| 时间操作 | [gotime](gotime/gotime.go) | 时间输出等 |
| jwt | [jwt](jwt/jwt.go) | jwt编码解码 |
| 流量限制 | [ratelimit](ratelimit/README.md) | 服务流量限制 |
| 距离计算 | [geo](geo/geo.go) | 计算两个经纬度之前的距离 |
| 敏感词过滤 | [sensitive](sensitive/sensitive.go) | 过滤敏感词汇 |
