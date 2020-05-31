package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("http get error : url=%v , statusCode=%v", url, response.StatusCode)
	}
	return ioutil.ReadAll(response.Body)
}
