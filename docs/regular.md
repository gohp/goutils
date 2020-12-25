# regular

- [x] 大陆手机号码判断
- [x] 银行卡号码判断
- [x] IP地址判断
- [x] 邮箱地址判断
- [x] 常规用户名判断
- [x] 身份证号码判断

## Usage

```go
package main

import (
    "fmt"
    "github.com/gohp/goutils/regular"
)

func main() {
	// ...

    // 用户名判断 仅包含a-z, A-Z, 0-9 的4到16位字符
	regular.IsUsername("A1023abbc23jol76")

    // 邮箱判断
    regular.IsEmail("1234@qq.com")

    // 手机号判断
    regular.IsMobile("+8613500001111")
    regular.IsMobile("8613500001111")
    regular.IsMobile("13500001111")

    // IP判断
    regular.IsIpv4Addr("255.255.255.255")

    // 银行卡判断
    regular.IsBankNo("62220000000000000000")

    // 身份证判断
    regular.IsIdCardNo("0123456789")
}
```