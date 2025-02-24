package terminal

import (
	"github.com/gin-gonic/gin"

	"tdp-cloud/helper/webssh"
)

func ssh(c *gin.Context) {

	// 获取 SSH 参数

	var option *webssh.SSHClientOption

	if err := c.ShouldBindQuery(&option); err != nil {
		c.Set("Error", err)
		return
	}

	// 创建 SSH 连接

	err := webssh.Connect(&webssh.ConnectParam{
		Request: c.Request,
		Writer:  c.Writer,
		Option:  option,
	})

	if err != nil {
		c.Set("Error", err)
		return
	}

}
