package api

import (
	"context"
	"first/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Register(ctx context.Context, c *app.RequestContext) {
	var req service.Registeruser
	err := c.Bind(&req)
	if err != nil {
		c.JSON(400, utils.H{
			"error": "参数错误",
		})
		return
	}
	err = service.Register(req)
	if err != nil {
		c.JSON(400, utils.H{
			"error": "注册失败",
		})
		return
	}
	c.JSON(200, utils.H{"error": "注册成功"})
}

func Login(ctx context.Context, c *app.RequestContext) {
	req := service.Loginuser{}
	if err := c.Bind(&req); err != nil {
		c.JSON(400, utils.H{"error": "参数错误"})
		return
	}
	err := service.Login(req)
	if err != nil {
		c.JSON(401, utils.H{"error": "用户名或密码错误"})
		return
	}
	c.JSON(200, utils.H{"message": "登录成功"})
}
