package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
)

// UserRegister 获取路由中 url 的数据信息，调用 server 层生成 token
func UserRegister(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// 调用
	token, err := services.UserRegister(username, password)
	if err != nil {
		c.JSON(http.StatusOK, utils.LoginResponse{
			Response: utils.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, utils.LoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   token.UserId,
		Token:    token.Token,
	})
}

// UserLogin 用户登录
func UserLogin(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	// 调用
	token, err := services.UserLogin(username, password)
	if err != nil {
		c.JSON(http.StatusOK, utils.LoginResponse{
			Response: utils.Response{
				StatusCode: 1,
				StatusMsg:  err.Error(),
			},
		})
		return
	}
	c.JSON(http.StatusOK, utils.LoginResponse{
		Response: utils.Response{StatusCode: 0},
		UserId:   token.UserId,
		Token:    token.Token,
	})
}
