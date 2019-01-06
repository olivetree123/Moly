package http

import (
	"fmt"
	"net/http"
	"strings"
	"time"
)

// Reponse http 返回值结构
type Reponse struct {
	Protocol   string
	Headers    map[string]string
	StatusCode int
	Status     string
	Body       string
}

// Init 初始化
func (resp *Reponse) Init() {
	resp.Protocol = "HTTP/1.1"
	resp.SetCode(200)
	resp.Headers = make(map[string]string)
}

// SetCode 设置返回码
func (resp *Reponse) SetCode(code int) {
	resp.StatusCode = code
	if code == 200 {
		resp.Status = "OK"
	} else if code == 404 {
		resp.Status = "Not Found"
	}

}

// SetHeader 设置返回的header
func (resp *Reponse) SetHeader(key, value string) {
	key = strings.ToLower(key)
	value = strings.ToLower(value)
	resp.Headers[key] = value
}

// SetBody 设置 HTTP body
func (resp *Reponse) SetBody(content string, contentType string) {
	resp.Body = content
	resp.SetHeader("Content-Type", contentType)
}

// ToString 转换成字符串
func (resp *Reponse) ToString() string {
	// context := "HTTP/1.1 200 OK\nDate: Sat, 31 Dec 2005 23:59:59 GMT\nContent-Type: text/html;charset=ISO-8859-1\nContent-Length: 3\n\n123"
	format := "%s %d %s\n%s\n%s\n\n%s"
	headerText := ""
	for key, value := range resp.Headers {
		headerText += strings.Join([]string{key, value}, ":")
	}
	timeStr := time.Now().UTC().Format(http.TimeFormat)
	content := fmt.Sprintf(format, resp.Protocol, resp.StatusCode, resp.Status, timeStr, headerText, resp.Body)
	// fmt.Println("content = ", content)
	return content
}
