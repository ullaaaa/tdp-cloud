package tencent

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"

	"tdp-cloud/helper/tencent"
	"tdp-cloud/module/dborm/vendor"
)

func apiProxy(c *gin.Context) {

	rq := &vendor.FetchParam{
		Id:     cast.ToUint(c.Param("id")),
		UserId: c.GetUint("UserId"),
	}

	if rq.Id == 0 {
		c.Set("Error", "参数错误")
		return
	}

	vendor, err := vendor.Fetch(rq)

	if err != nil || vendor.Id == 0 {
		c.Set("Error", "厂商不存在")
		return
	}

	// 构造参数

	params := &tencent.Params{
		SecretId:  vendor.SecretId,
		SecretKey: vendor.SecretKey,
	}

	if err = c.ShouldBindJSON(params); err != nil {
		c.Set("Error", err)
		return
	}

	// 发起请求

	if res, err := tencent.Request(params); err == nil {
		c.Set("Payload", res)
	} else {
		c.Set("Error", err)
	}

}

func vncProxy(c *gin.Context) {

	resp, err := http.Get("https://img.qcloud.com/qcloud/app/active_vnc/index.html")

	if err != nil {
		c.Set("Error", "获取资源失败")
		return
	}

	if res, err := ioutil.ReadAll(resp.Body); err == nil {
		c.Set("HTML", string(res))
	} else {
		c.Set("Error", err)
	}

}
