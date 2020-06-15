package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/15
 */

// http server
func main() {
	type RespData struct {
		Params string `json:"params"`
		Success bool `json:"success"`
		UA string `json:"ua"`
		Charset string `json:"charset"`
		TestHeader string `json:"test_header"`
		ContentType string `json:"content_type"`
	}

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			fmt.Printf("receive get request\n")
			resp := &RespData{
				Success: true,
				Params: r.URL.String(),
				UA: r.UserAgent(),
				TestHeader: r.Header.Get("Test-Header"),
			}
			data, _ := json.Marshal(resp)
			w.Write(data)
		case http.MethodPost:
			fmt.Printf("receive post request\n")
			resp := &RespData{
				Success: true,
				Params: r.URL.String(),
				UA: r.UserAgent(),
				TestHeader: r.Header.Get("Test-Header"),
				ContentType: r.Header.Get("Content-Type"),
				Charset: r.Header.Get("Charset"),
			}
			data, _ := json.Marshal(resp)
			w.Write(data)
		}
	})

	if err := http.ListenAndServe(":7777", nil); err != nil {
		panic(err)
	}
}
