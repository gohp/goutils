package geo

import (
	. "github.com/smartystreets/goconvey/convey"
	"math"
	"testing"
)

func TestCalDistance(t *testing.T) {
	Convey("test cal distance", t, func() {
		BeijingPoint := Point{lat: 39.9042, lon: 116.4074}
		ShenzhenPoint := Point{lat: 22.5431, lon: 114.0579}
		// 深圳到北京的经纬度 直线距离约为1942 米
		dis := CalDistance(BeijingPoint.Lon(), BeijingPoint.Lat(), ShenzhenPoint.Lon(), ShenzhenPoint.Lat())
		t.Log(dis)
		// 误差2km
		So(math.Abs(dis-1942), ShouldBeLessThan, 2)
	})
}
