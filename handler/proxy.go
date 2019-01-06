package handler

import (
	"fmt"
	"io/ioutil"
	// "github.com/julienschmidt/httprouter"
	"Moly/models"
	"net/http"
)

// ProxyHandler 处理代理请求
func ProxyHandler(w http.ResponseWriter, r *http.Request) {
	status, destURL := false, ""
	rules := models.ListRule()
	for _, rule := range rules {
		status, destURL = rule.Match(r.URL.Path)
		if status {
			destURL = "http://" + rule.DestHost + ":" + rule.DestPort + destURL
			break
		}
	}
	if !status {
		// 404
		w.WriteHeader(404)
		w.Write([]byte("NOT FOUND"))
		return
	}
	fmt.Println("srcURL = ", r.URL.Path, ", destURL = ", destURL, ", status = ", status)
	code, resp := ProxyRequest(r, destURL)
	w.WriteHeader(code)
	w.Write(resp)
}

// ProxyRequest 转发请求
func ProxyRequest(request *http.Request, url string) (int, []byte) {
	client := &http.Client{}
	req, err := http.NewRequest(request.Method, url, request.Body)
	if err != nil {
		panic(err)
	}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return resp.StatusCode, body
}
