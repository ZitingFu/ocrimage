package http_tools

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

/**
* 发送http请求 并返回响应结果
*
* @param url 请求的url
* @param method 请求方式
* @param param 请求参数
* @return string 响应结果
 */
func HttpRequest(url string, method string, queryString string, data string, header map[string]string) string {
	response := ""
	client := &http.Client{}
	if queryString != "" {
		url += "?" + queryString
	}

	dataReader := strings.NewReader(data)

	log.Printf("method %s,url %s", method, url)
	req, err := http.NewRequest(method, url, dataReader)

	if err != nil {
		log.Printf("创建请求失败:%v,%v,%v", url, method, data)
		return response
	}
	//设置请求头
	if header != nil {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	//log.Printf("request params:%s", data)
	resp, err := client.Do(req)
	log.Printf("response info:%v", resp)
	if err != nil {
		log.Printf("发送请求失败:%v,%v,%v", url, method, data)
		return response
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("读取响应失败:%v,%v,%v", url, method, data)
		return response
	}
	log.Printf("响应内容:%s", string(body))
	return string(body)
}
