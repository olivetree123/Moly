package handler

import (
	"Moly/models"
	"Moly/utils"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// PostService 注册服务
func PostService(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	service := models.NewServiceFromHTTPBody(r.Body)
	if !service.Validate() {
		resp := utils.NewResponse(1000, nil)
		w.Write(resp.ToBytes())
		return
	}
	models.DB.Create(service)
	if service.ID == 0 {
		resp := utils.NewResponse(1001, nil)
		w.Write(resp.ToBytes())
		return
	}
	resp := utils.NewResponse(0, service)
	w.Write(resp.ToBytes())
}

// GetService 获取服务详情
func GetService(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// GetServiceList 获取服务列表
func GetServiceList(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var ss []models.Service
	models.DB.Find(&ss)
	rs, err := json.Marshal(ss)
	if err != nil {
		panic(err)
	}
	rs = bytes.ToLower(rs)
	w.Write([]byte(rs))
}
