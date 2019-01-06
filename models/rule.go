package models

import (
	"fmt"
	"io"
	"strings"
	// "Moly/entity"
	"encoding/json"
	"github.com/jinzhu/gorm"
)

// Rule 转发规则
type Rule struct {
	gorm.Model
	SrcPath  string `gorm:"not null;"`
	DestHost string `gorm:"not null;"`
	DestPort string `gorm:"not null;"`
	DestPath string `gorm:"not null;"`
	Type     string `gorm:"not null;"`
}

// NewRuleFromHTTPBody 从 HTTP 的 body 读取数据，生成 Rule 对象
func NewRuleFromHTTPBody(body io.ReadCloser) *Rule {
	var rule Rule
	r := json.NewDecoder(body)
	r.Decode(&rule)
	return &rule
}

// Validate 验证是否合法
func (rule *Rule) Validate() bool {
	if rule.SrcPath == "" || rule.DestHost == "" || rule.DestPort == "" || rule.DestPath == "" {
		fmt.Println(rule.SrcPath, rule.DestHost, rule.DestPort, rule.DestPath)
		return false
	}
	return true
}

// Match 检查是否匹配
func (rule *Rule) Match(url string) (bool, string) {
	if strings.ToUpper(rule.Type) == "PATHPREFIX" {
		// 前缀匹配
		if len(url) >= len(rule.SrcPath) && url[:len(rule.SrcPath)] == rule.SrcPath {
			return true, rule.DestPath + url[len(rule.SrcPath):]
		}
	} else if strings.ToUpper(rule.Type) == "PATH" {
		// 完全匹配
		if rule.SrcPath == url {
			return true, url
		}
	} else if strings.ToUpper(rule.Type) == "PATHPREFIXSTRIP" {
		// 前缀匹配，并去掉前缀
		if url[:len(rule.SrcPath)] == rule.SrcPath {
			return true, url[len(rule.SrcPath):]
		}
	}
	return false, ""
}

// ListRule 获取 Rule 列表
func ListRule() []Rule {
	var rules []Rule
	DB.Find(&rules)
	return rules
}
