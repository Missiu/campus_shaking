package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
)

// UserInfo 获取用户信息
func UserInfo(c *gin.Context) {
	// 获取并解析token，得到userid
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}
	//验证token
	tokenStruck, ok := utils.ParseToken(tokenStr)
	if !ok {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "ParseToken ERROR "},
		})
		return
	}
	info, err := services.UserInfo(tokenStruck.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "userInfo Get ERROR "},
		})
		return
	}
	c.JSON(http.StatusOK, utils.InfoResponse{
		Response: utils.Response{StatusCode: 0},
		User:     info,
	})
}
