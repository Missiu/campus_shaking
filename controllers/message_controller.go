package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
	"strconv"
	"time"
)

// var lastTimestamp int64 // 全局变量，保存上次轮询请求的时间戳
func FriendList(c *gin.Context) {
	// 获取并解析token，得到userid
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}

	//验证token
	tokenStruck, ok := utils.ParseToken(tokenStr)
	if !ok {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1, StatusMsg: "ParseToken ERROR ",
		})
		return
	}

	list, err := services.FriendList(tokenStruck.ID)
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
		Response: utils.Response{
			StatusCode: 0,
		},
		UserList: list,
	})
}

func SendMessage(c *gin.Context) {
	// 获取并解析token，得到userid
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}
	//验证token
	tokenStruck, ok := utils.ParseToken(tokenStr)
	if !ok {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1, StatusMsg: "ParseToken ERROR ",
		})
		return
	}
	// 获取聊天对象id
	toUserIdStr := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1, StatusMsg: "ParseInt ERROR ",
		})
		return
	}
	// 获取状态
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1, StatusMsg: "actionType Get ERROR ",
		})
		return
	}

	if actionType == 1 {
		// 获取评论内容
		content := c.Query("content")
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{
				StatusCode: 1, StatusMsg: "content Get ERROR ",
			})
			return
		}
		err = services.SendMessage(tokenStruck.ID, toUserId, content)
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{
				StatusCode: 1, StatusMsg: "SendMessage  ERROR ",
			})
			return
		}
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 0,
		})
	}
}

func ChatHistory(c *gin.Context) {
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
	toUserIdStr := c.Query("to_user_id")
	toUserId, err := strconv.ParseInt(toUserIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "toUserId Get ERROR "},
		})
		return
	}
	// 上次最新消息的时间
	msgTimeStr := c.Query("pre_msg_time")
	msgTime, err := strconv.ParseInt(msgTimeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "msgTime Get ERROR "},
		})
		return
	}
	chatHistory, err := services.ChatHistory(tokenStruck.ID, toUserId)
	if err != nil {
		c.JSON(http.StatusOK, utils.ApifoxModel{
			Response: utils.Response{StatusCode: 1, StatusMsg: "ChatHistory ERROR "},
		})
		return
	}
	timeNow := time.Now().Unix()
	var rUserid int64
	var rTouserid int64
	// 没有新消息
	if msgTime < timeNow && msgTime != 0 {
		c.JSON(http.StatusOK, utils.ApifoxModel{
			Response:    utils.Response{StatusCode: 0},
			MessageList: nil,
		})
	} else {
		// 聊天对象没变
		if rUserid == tokenStruck.ID && rTouserid == toUserId {
			c.JSON(http.StatusOK, utils.ApifoxModel{
				Response:    utils.Response{StatusCode: 0},
				MessageList: nil,
			})
			return
		} else {
			rUserid = tokenStruck.ID
			rTouserid = toUserId
			c.JSON(http.StatusOK, utils.ApifoxModel{
				Response:    utils.Response{StatusCode: 0},
				MessageList: chatHistory,
			})
		}
		return
	}
}
