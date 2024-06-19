package utils

import "new/models"

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

// FeedVideoResponse 视频流、视频发布（不含NextTime）、喜欢列表（不含NextTime）响应
type FeedVideoResponse struct {
	Response
	NextTime int64            `json:"next_time,omitempty"`
	Videos   []*models.Videos `json:"video_list,omitempty"`
}

// LoginResponse 用户登录或注册响应
type LoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

// InfoResponse 用户信息响应
type InfoResponse struct {
	Response
	User models.UserInfo `json:"user"`
}

// CommentResponse 用户评论响应
type CommentResponse struct {
	Response
	Comment models.Comment `json:"comment"`
}

// CommentsResponse 评论响应
type CommentsResponse struct {
	Response
	Comment *models.Comment `json:"comment"`
}

// CommentsListResponse 评论列表响应
type CommentsListResponse struct {
	Response
	CommentList []*models.Comment `json:"comment_list"`
}

// FollowList 关注、粉丝、好友列表响应
type FollowList struct {
	Response
	UserList []*models.UserInfo `json:"user_list"`
}

// ApifoxModel 聊天列表响应
type ApifoxModel struct {
	Response
	MessageList []*models.Message `json:"message_list"` // 用户列表
}
