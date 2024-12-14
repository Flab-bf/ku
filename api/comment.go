package api

import (
	"context"
	"first/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

func Writer(ctx context.Context, c *app.RequestContext) {
	req := service.Comment{}
	if err := c.Bind(&req); err != nil {
		c.JSON(400, utils.H{"error": "参数错误"})
		return
	}
	err := service.WriteComment(req)
	if err != nil {
		c.JSON(400, utils.H{"error": "留言发表失败"})
		return
	}
	c.JSON(200, utils.H{"message": "留言发表成功"})
}

func Delete(ctx context.Context, c *app.RequestContext) {
	var req service.Comment
	if err := c.Bind(&req); err != nil {
		c.JSON(400, utils.H{"error": "参数错误"})
		return
	}
	err := service.DeleteComments(req.Id)
	if err != nil {
		c.JSON(400, utils.H{"error": "留言删除失败"})
		return
	}
	c.JSON(200, utils.H{"message": "删除成功"})
}

func GetAll(ctx context.Context, c *app.RequestContext) {
	Comments, err := service.ReadAll()
	if err != nil {
		c.JSON(500, utils.H{"error": "获取错误"})
		return
	}
	c.JSON(200, utils.H{"comments": Comments})
}
