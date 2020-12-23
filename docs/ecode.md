# ecode

error code 封装业务错误码

提取自bilibili


## Usage

```go
import (
	"github.com/wzyonggege/goutils/ecode"
	"github.com/wzyonggege/goutils/http"
	"log"
)

var (
	NetworkErr = ecode.New(10001, "network error")
)

func main()  {
	_, err := http.HttpGet("baidu.com")
	if err != nil {
		log.Fatal(NetworkErr)
	}
}
```
