# Geo

## 1. 计算两个经纬度的距离

> http://www.movable-type.co.uk/scripts/latlong.html

### Usage

```go

import (
	"github.com/gohp/goutils/geo"
	"log"
)

func main() {
    BeijingPoint := geo.Point{lat: 39.9042, lon: 116.4074}
    ShenzhenPoint := geo.Point{lat: 22.5431, lon: 114.0579}

    // 深圳到北京的经纬度 直线距离约为1942 米
    dis := geo.CalDistance(BeijingPoint.Lon(), BeijingPoint.Lat(), ShenzhenPoint.Lon(), ShenzhenPoint.Lat())
    // ...
}
```
