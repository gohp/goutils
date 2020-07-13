package regular

import (
	"fmt"
	"github.com/wzyonggege/goutils/httplib"
	"regexp"
	"strings"
)

var (
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

// IsEmail ...
func IsEmail(v string) bool {
	return emailRegexp.MatchString(v)
}

// IsMobile ...
func IsMobile(v string) bool {
	return mobileRegexp.MatchString(v)
}

// IsIpv4Addr...
func IsIpv4Addr(v string) bool {
	return ipv4Regexp.MatchString(v)
}

// IsBankNo ... from alipay
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

func IsIdCardNo(idCardNo string) bool {
	var (
		coefficient []int = []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
		code        []byte  = []byte{'1', '0', 'X', '9', '8', '7', '6', '5', '4', '3', '2'}
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