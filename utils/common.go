package utils

import (
	"github.com/gin-gonic/gin"
)

type MyResponse struct {
	Status int         `json:"status"`
	Msg    string      `json:"msg" default:"-"`
	Error  interface{} `json:"error" default:"-"`
	Data   interface{} `json:"data"`
}

func (m *MyResponse) WriteJson(c *gin.Context) {
	c.JSON(m.Status, m)
}
