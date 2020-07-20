package regular

import (
	"fmt"
	"github.com/wzyonggege/goutils/httplib"
	"regexp"
	"strings"
)

var (
	// 用户名判断 仅包含a-z, A-Z, 0-9 的4到16位字符
	usernameRegexp = regexp.MustCompile("^[a-zA-Z0-9]{4,16}$")
	// 邮箱判断
	emailRegexp = regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)
	// 大陆手机号判断
	mobileRegexp = regexp.MustCompile(
		`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`,
	)
	// IPv4地址
	ipv4Regexp = regexp.MustCompile(
		`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`,
	)
)

type ValidateBankCard struct {
	CardType  string        `json:"cardType"`
	Bank      string        `json:"bank"`
	Key       string        `json:"key"`
	Messages  []interface{} `json:"messages"`
	Validated bool          `json:"validated"`
	Stat      string        `json:"stat"`
}

// IsUsername 判断是否常规用户名：仅包含a-z, A-Z, 0-9 的4到16位字符
func IsUsername(username string) (b bool) {
	return usernameRegexp.MatchString(username)
}

// IsEmail 邮箱判断
func IsEmail(email string) bool {
	return emailRegexp.MatchString(email)
}

// IsMobile 大陆手机号判断
func IsMobile(phoneNo string) bool {
	return mobileRegexp.MatchString(phoneNo)
}

// IsIpv4Addr IPv4地址判断
func IsIpv4Addr(addr string) bool {
	return ipv4Regexp.MatchString(addr)
}

// IsBankNo 银行卡号判断... from alipay
func IsBankNo(bankCard string) bool {
	url := "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardNo=%s&cardBinCheck=true"
	url = fmt.Sprintf(url, bankCard)
	result := ValidateBankCard{}

	err := httplib.Get(url).ToJson(&result)
	if err != nil {
		return false
	}
	if result.Validated == true && result.Bank != "" {
		return true
	}
	return false
}

// IsIdCardNo 身份证号判断
func IsIdCardNo(idCardNo string) bool {
	/*
		身份证校验码的计算方法：
		1、将前面的身份证号码17位数分别乘以不同的系数，系数见：coefficient
		2、将这17位数字和系数相乘的结果相加，用加出来和除以11，得到余数Remainder
		3、余数Remainder作为位置值，在数组code中找到对应的值，就是身份证号码的第18位数值
	*/
	var (
		coefficient = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		code        = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
	)
	if len(idCardNo) != 18 {
		return false
	}

	idByte := []byte(strings.ToUpper(idCardNo))

	sum := 0
	for i := 0; i < 17; i++ {
		sum += int(idByte[i]-byte('0')) * coefficient[i]
	}
	return code[sum%11] == idByte[17]
}
