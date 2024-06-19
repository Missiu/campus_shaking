package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
	"strconv"
)

func Like(c *gin.Context) {
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
	// 获取视频id
	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "videoId Get ERROR "},
		})
		return
	}
	// 获取状态
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "actionType Get ERROR "},
		})
		return
	}
	err = services.Like(tokenStruck.ID, videoId, actionType)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "Like ERROR ",
		})
	}
	c.JSON(http.StatusOK, utils.Response{StatusCode: 0})

}

func LikeVideoList(c *gin.Context) {
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
	likeVideoList, err := services.LikeVideoList(tokenStruck.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "likeVideoList ERROR",
		})
		return
	}
	c.JSON(http.StatusOK, utils.FeedVideoResponse{
		Response: utils.Response{StatusCode: 0},
		Videos:   likeVideoList,
	})
}
