package handler

import (
	"Moly/models"
	"Moly/utils"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// PostRule 创建转发规则
func PostRule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	rule := models.NewRuleFromHTTPBody(r.Body)
	if !rule.Validate() {
		resp := utils.NewResponse(1000, nil)
		w.Write(resp.ToBytes())
		return
	}
	models.DB.Create(rule)
	if rule.ID == 0 {
		resp := utils.NewResponse(1001, nil)
		w.Write(resp.ToBytes())
		return
	}
	resp := utils.NewResponse(0, rule)
	w.Write(resp.ToBytes())
	return
}

// ListRule 获取规则列表
func ListRule(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// var rules []models.Rule
	// models.DB.Find(&rules)
	rules := models.ListRule()
	resp := utils.NewResponse(0, rules)
	w.Write(resp.ToBytes())
	return
}
