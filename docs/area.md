# Area2Zone

- [x] 根据区号查询行政区
- [x] 根据行政编码查询行政区
- [x] 查询上级行政区
- [x] 查询子行政区
- [x] 获取经纬度
- [ ] 根据名称查询行政区

## Usage

```go
package main

import (
    "fmt"
    "github.com/gohp/goutils/area"
)

func main() {
	// load area file
	z := area.Load("./area.json")
    
    // 根据区号查询
	ret := z.GetZoneByCityCode("0755")
	if ret != nil {
		fmt.Printf("city code:0755 get zone: %s\n", ret.GetName())
	}
    
    // 根据行政编码查询
	ret2 := z.GetZoneByAdCode("440306")
	if ret2 != nil {
		fmt.Printf("adcode: 440306 get zone: %s\n", ret2.GetName())
	}

    // 查询上级行政区
	ret3 := z.GetFather(ret2)
	if ret3 != nil {
		fmt.Printf("zone %s get father: %s\n", ret2.GetName(), ret3.GetName())
	}

    // 查询子行政区
	ret4 := z.GetChildren(ret)
	if ret4 != nil {
		fmt.Printf("%s get children %d\n", ret.GetName(), len(ret4))
		for _, k := range ret4 {
			if k != nil {
                
                // 获取经纬度
				lat, lon := k.GetLocation()
				fmt.Printf("get %s's children: %s location [lat:%s lon:%s]\n", ret.GetName(), k.GetName(), lat, lon )
			}
		}

	}
}

```