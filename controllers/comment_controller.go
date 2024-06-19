package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
	"strconv"
)

func Comment(c *gin.Context) {
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
	// 获取评论状态 1-评论 2-删除评论
	actionTypeStr := c.Query("action_type")
	actionType, err := strconv.ParseInt(actionTypeStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "actionType Get ERROR "},
		})
		return
	}

	if actionType == 1 {
		// 获取评论内容
		commentText := c.Query("comment_text")
		if err != nil {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "commentText Get ERROR "},
			})
			return
		}
		comment, err := services.Comment(tokenStruck.ID, videoId, actionType, 0, commentText)
		if err != nil {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "comment Get ERROR "},
			})
			return
		}
		c.JSON(http.StatusOK, utils.CommentsResponse{
			Response: utils.Response{StatusCode: 0},
			Comment:  comment,
		})
	}
	if actionType == 2 {
		// 获取视频id
		commentStr := c.Query("comment_id")
		commentId, err := strconv.ParseInt(commentStr, 10, 64)
		if err != nil {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "commentId Get ERROR "},
			})
			return
		}
		comment, err := services.Comment(tokenStruck.ID, videoId, actionType, commentId, "")
		if err != nil {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "comment Get ERROR "},
			})
			return
		}
		c.JSON(http.StatusOK, utils.CommentsResponse{
			Response: utils.Response{StatusCode: 0},
			Comment:  comment,
		})
	}
}

func CommentList(c *gin.Context) {
	// 获取视频id
	videoIdStr := c.Query("video_id")
	videoId, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{StatusCode: 1, StatusMsg: "videoId Get ERROR "},
		})
		return
	}
	// 获取并解析token，得到userid
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}
	// 未登录状态
	if tokenStr != "" {
		//验证token
		tokenStruck, ok := utils.ParseToken(tokenStr)
		if !ok {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "ParseToken ERROR "},
			})
			return
		}
		commentList, err := services.CommentList(tokenStruck.ID, videoId)
		if err != nil {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "commentList Get ERROR "},
			})
			return
		}
		c.JSON(http.StatusOK, utils.CommentsListResponse{
			Response:    utils.Response{StatusCode: 0},
			CommentList: commentList,
		})
	} else {
		commentList, err := services.CommentList(0, videoId)
		if err != nil {
			c.JSON(http.StatusOK, utils.InfoResponse{
				Response: utils.Response{StatusCode: 1, StatusMsg: "commentList Get ERROR "},
			})
			return
		}
		c.JSON(http.StatusOK, utils.CommentsListResponse{
			Response:    utils.Response{StatusCode: 0},
			CommentList: commentList,
		})
	}
}
