package main

import (
	"Moly/handler"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// func httpHandler(request *http.Request) string {
// 	fmt.Println(request.URL.Path)
// 	//fmt.Println(request.URL)
// 	fmt.Println("host = ", request.Host)
// 	port := strings.Split(request.Host, ":")[1]

// 	if port == "8000" {
// 		// 网关本身提供的服务
// 		return ""
// 	}
// 	url := "http://" + request.Host + request.URL.Path
// 	r := rule.Rule{
// 		Type:    "pathPrefix",
// 		Content: "/gaojian",
// 		Host:    "localhost",
// 		Port:    5000,
// 	}
// 	status, url2 := r.Match(request.URL.Path)
// 	resp := &moly_http.Reponse{}
// 	resp.Init()
// 	if !status {
// 		resp.SetCode(404)
// 		return resp.ToString()
// 	}
// 	// context := "HTTP/1.1 200 OK\nDate: Sat, 31 Dec 2005 23:59:59 GMT\nContent-Type: text/html;charset=ISO-8859-1\nContent-Length: 3\n\n123"
// 	url = "http://" + r.Host + ":" + strconv.Itoa(r.Port) + "/" + url2
// 	client := &http.Client{}
// 	req, err := http.NewRequest(request.Method, url, request.Body)
// 	if err != nil {
// 		panic(err)
// 	}
// 	res, err := client.Do(req)
// 	if err != nil {
// 		panic(err)
// 	}
//	var body []byte
//	_, err = res.Body.Read(body)
//	if err != nil {
//		panic(err)
//	}
// 	resp.SetCode(res.StatusCode)
// 	resp.SetBody(string(body), resp.Headers["Content-Type"])
// 	return resp.ToString()
// }

// func handler(conn net.Conn) string {
// 	for {
// 		buf := make([]byte, 1024)
// 		length, err := conn.Read(buf)
// 		if err != nil {
// 			panic(err)
// 		}
// 		content := string(buf[:length])
// 		r := bufio.NewReader(strings.NewReader(content))
// 		lines := strings.Split(content, "\n")
// 		fmt.Println(lines[0])
// 		rs := strings.Split(lines[0], " ")
// 		if len(rs) == 3 && rs[2][:6] == "HTTP/1" {
// 			// HTTP 协议转发
// 			request, err := http.ReadRequest(r)
// 			if err != nil {
// 				panic(err)
// 			}
// 			content := httpHandler(request)
// 			conn.Write([]byte(content))
// 			fmt.Println("conn close !")
// 			conn.Close()
// 		} else {
// 			// TCP 协议转发
// 			fmt.Println(rs)
// 		}
// 		break
// 	}
// 	print("return")
// 	return ""
// }

func main() {
	router := httprouter.New()
	// 注册服务
	router.POST("/moly/service", handler.PostService)
	// 获取服务列表
	router.GET("/moly/service/list", handler.GetServiceList)
	// 获取服务详情
	// router.GET("/moly/service/:id", handler.GetService)
	// 创建转发规则
	router.POST("/moly/rule", handler.PostRule)
	// 获取转发规则列表
	router.GET("/moly/rule/list", handler.ListRule)
	// 代理服务
	// router.GET("/", handler.ProxyHandler)
	// router.POST("/", handler.ProxyHandler)
	router.NotFound = http.HandlerFunc(handler.ProxyHandler)
	http.ListenAndServe(":8000", router)
}
