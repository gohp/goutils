package curip

import (
	. "github.com/smartystreets/goconvey/convey"
	"github.com/gohp/goutils/regular"
	"testing"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/19
 */

func Test_GetExternalIP(t *testing.T) {
	Convey("Test GetExternalIP", t, func() {
		ip, err := GetExternalIP()
		So(err, ShouldBeNil)
		t.Logf("%s", ip)
		So(regular.IsIpv4Addr(ip), ShouldBeTrue)
	})
}

func Test_LocalIP(t *testing.T) {
	Convey("Test LocalIP", t, func() {
		ip, err := LocalIP()
		So(err, ShouldBeNil)
		So(regular.IsIpv4Addr(ip), ShouldBeTrue)
		t.Log(ip)
	})
}

func Test_LocalDnsName(t *testing.T) {
	Convey("Test LocalDnsName", t, func() {
		dns, err := LocalDnsName()
		So(err, ShouldBeNil)
		t.Log(dns)
	})
}

func Test_IntranetIP(t *testing.T) {
	Convey("Test IntranetIP", t, func() {
		ips, err := IntranetIP()
		So(err, ShouldBeNil)
		for _, ip := range ips {
			So(regular.IsIpv4Addr(ip), ShouldBeTrue)
		}
	})
}
