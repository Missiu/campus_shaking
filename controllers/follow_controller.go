package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
	"strconv"
)

func Follow(c *gin.Context) {
	// 获取并解析token
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
	// 被关注者id
	toUserIdStr := c.Query("to_user_id")
	videoId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "toUserIdStr Get ERROR "},
		})
		return
	}
	// 获取状态 1-关注 2-取消关注
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "actionType Get ERROR "},
		})
		return
	}
	err = services.Follow(tokenStruck.ID, videoId, actionType)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "Follow ERROR ",
		})
	} else {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 0,
		})
	}
}

func FollowerList(c *gin.Context) {
	// 获取并解析token
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
	list, err := services.FollowerList(tokenStruck.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.FollowList{
			Response: utils.Response{
				StatusCode: 1,
				StatusMsg:  "FollowerList ERROR",
			},
		})
		return
	}
	c.JSON(http.StatusOK, utils.FollowList{
		Response: utils.Response{StatusCode: 0},
		UserList: list,
	})
}

func FollowList(c *gin.Context) {
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
	list, err := services.FollowList(tokenStruck.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.FollowList{
			Response: utils.Response{
				StatusCode: 1,
				StatusMsg:  "FollowList ERROR",
			},
		})
		return
	}
	c.JSON(http.StatusOK, utils.FollowList{
		Response: utils.Response{StatusCode: 0},
		UserList: list,
	})
}
