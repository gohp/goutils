package regular

import (
	"encoding/json"
	"fmt"
	"github.com/wzyonggege/goutils/http"
	"regexp"
)

var (
	emailRegexp = regexp.MustCompile(
		"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
	)

	mobileRegexp = regexp.MustCompile(
		`^(?:\+?86)?1(?:3\d{3}|5[^4\D]\d{2}|8\d{3}|7(?:[01356789]\d{2}|4(?:0\d|1[0-2]|9\d))|9[189]\d{2}|6[567]\d{2}|4[579]\d{2})\d{6}$`,
	)

	ipv4Regexp = regexp.MustCompile(
		`^((25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(25[0-5]|2[0-4]\d|[01]?\d\d?)$`,
	)
)

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
func IsBankNo(bankCard string) (b bool) {
	type ValidateBankCard struct {
		CardType  string        `json:"cardType"`
		Bank      string        `json:"bank"`
		Key       string        `json:"key"`
		Messages  []interface{} `json:"messages"`
		Validated bool          `json:"validated"`
		Stat      string        `json:"stat"`
	}
	url := "https://ccdcapi.alipay.com/validateAndCacheCardInfo.json?_input_charset=utf-8&cardNo=%s&cardBinCheck=true"
	url = fmt.Sprintf(url, bankCard)
	resp, err := http.HttpGet(url)
	if err != nil {
		return
	}
	result := ValidateBankCard{}
	err = json.Unmarshal(resp, &result)
	if err != nil {
		return
	}
	if result.Validated == true && result.Bank != "" {
		return true
	}
	return
}
