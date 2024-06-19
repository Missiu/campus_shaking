package routes

import (
	"github.com/gin-gonic/gin"
	"new/controllers"
)

func Router() *gin.Engine {
	// gin的初始化
	r := gin.Default()
	// 静态文件服务路由
	r.Static("/static", "./public")

	r.GET("/douyin/feed/", controllers.GetVideoFeed)                   // 视频流
	r.POST("/douyin/user/register/", controllers.UserRegister)         // 注册
	r.POST("/douyin/user/login/", controllers.UserLogin)               // 登录
	r.GET("/douyin/user/", controllers.UserInfo)                       // 用户信息
	r.POST("/douyin/publish/action/", controllers.PushVideo)           // 投稿视频
	r.GET("/douyin/publish/list/", controllers.GetVideoList)           // 发布列表
	r.POST("/douyin/favorite/action/", controllers.Like)               // 点赞操作
	r.GET("/douyin/favorite/list/", controllers.LikeVideoList)         // 获取点赞列表
	r.POST("/douyin/comment/action/", controllers.Comment)             // 评论操作
	r.GET("/douyin/comment/list/", controllers.CommentList)            // 获取评论列表
	r.POST("/douyin/relation/action/", controllers.Follow)             // 关注操作
	r.GET("/douyin/relation/follow/list/", controllers.FollowList)     // 获取关注列表
	r.GET("/douyin/relation/follower/list/", controllers.FollowerList) // 获取粉丝列表
	r.GET("/douyin/relation/friend/list/", controllers.FriendList)     // 获取好友列表
	r.POST("/douyin/message/action/", controllers.SendMessage)         // 发送消息
	r.GET("/douyin/message/chat/", controllers.ChatHistory)            // 聊天记录
	return r
}
