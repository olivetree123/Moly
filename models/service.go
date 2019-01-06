package models

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"io"
)

// Service 服务
type Service struct {
	gorm.Model
	Name string `gorm:"not null;unique;"`
	Host string `gorm:"not null;"`
	Port int    `gorm:"not null;"`
}

// NewServiceFromHTTPBody 从 HTTP 的 body 读取数据，生成 Service 对象
func NewServiceFromHTTPBody(body io.ReadCloser) *Service {
	var s Service
	r := json.NewDecoder(body)
	r.Decode(&s)
	return &s
}

// Validate 验证是否合法
func (s *Service) Validate() bool {
	if s.Name == "" || s.Host == "" || s.Port == 0 {
		return false
	}
	return true
}
