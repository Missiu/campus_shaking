package services

import (
	"fmt"
	"new/database"
	"new/models"
	"time"
)

type CommentFlow struct {
	userId      int64
	videoId     int64
	commentId   int64
	actionType  int64
	commentText string
	comment     *models.Comment
}

type CommentListFlow struct {
	userId   int64
	videoId  int64
	comments []*models.Comment
}

func Comment(userid, videoId, actionType, commentId int64, commentText string) (*models.Comment, error) {
	return NewCommentFlow(userid, videoId, actionType, commentId, commentText).Do()
}
func NewCommentFlow(userid, videoId, actionType, commentId int64, commentText string) *CommentFlow {
	return &CommentFlow{userId: userid, videoId: videoId, actionType: actionType, commentId: commentId, commentText: commentText}
}

func CommentList(userid, videoId int64) ([]*models.Comment, error) {
	return NewCommentListFlow(userid, videoId).Do()
}
func NewCommentListFlow(userid, videoId int64) *CommentListFlow {
	return &CommentListFlow{userId: userid, videoId: videoId}
}

// Do 评论操作
func (p *CommentFlow) Do() (*models.Comment, error) {
	switch p.actionType {
	case 1:
		// 数据打包
		p.comment = &models.Comment{
			UserId:    p.userId,
			VideosId:  p.videoId,
			Content:   p.commentText,
			CreatedAt: time.Now(),
		}
		err := models.PushComment(p.comment)
		if err != nil {
			return nil, fmt.Errorf("PushComment 失败: %s", err)
		}
		database.UpdateFollowStatus(p.userId, p.videoId, true)
	case 2:
		err := models.DeleteComment(p.commentId, p.videoId)
		if err != nil {
			return nil, fmt.Errorf("DeleteComment 失败: %s", err)
		}
	default:
		panic(fmt.Sprintf("未定义的操作"))
	}
	return p.comment, nil
}

// Do 评论列表获取操作
func (g *CommentListFlow) Do() ([]*models.Comment, error) {
	err := models.GetCommentList(g.videoId, &g.comments)
	if err != nil {
		return nil, fmt.Errorf("GetCommentList 失败: %s", err)
	}

	// 评论信息填充
	for i := 0; i < len(g.comments); i++ {
		info, err := UserInfo(g.comments[i].UserId)
		if err != nil {
			return nil, fmt.Errorf("UserInfo 失败: %s", err)
		}
		g.comments[i].User = info
		g.comments[i].CreateDate = g.comments[i].CreatedAt.Format("01-02")
	}
	return g.comments, nil
}
