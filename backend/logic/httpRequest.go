package logic

import (
	"io/ioutil"
	"net/http"
	"strings"
	"unsafe"
)

type HttpRequest struct {
	header map[string]string
}

func NewRequest() *HttpRequest {
	return &HttpRequest{}
}

func (this *HttpRequest) SetHeader(key, val string) {
	this.header[key] = val
}

func (this *HttpRequest) sendPost(url, postData string) string {
	client := &http.Client{}
	req, err := http.NewRequest("post", url, strings.NewReader(postData))
	if err != nil {
		//
	}
	for key, val := range this.header {
		// req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set(key, val)
	}

	resp, err := client.Do(req)
	if err != nil {
		return ""
	}

	defer resp.Body.Close()
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	//byte数组直接转成string，优化内存
	str := (*string)(unsafe.Pointer(&respBytes))
	return *str
}
