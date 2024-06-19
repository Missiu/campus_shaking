package models

import (
	"fmt"
	"new/database"
	"time"
)

type Comment struct {
	ID         int64     `json:"id"`
	UserId     int64     `json:"-"` // 用户评论的id
	VideosId   int64     `json:"-"`
	User       UserInfo  `json:"user" gorm:"-"`
	Content    string    `json:"content"`
	CreatedAt  time.Time `json:"-"`
	CreateDate string    `json:"create_date" gorm:"-"`
}

func (c Comment) TableName() string {
	return "comments"
}

func PushComment(comment *Comment) error {
	// 开始数据库事务
	tx := database.GetDB().Begin()

	// 更新视频评论数据
	err := tx.Create(comment).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("PushComment: %s", err)
	}
	// 评论数量加一
	err = tx.Exec("UPDATE videos SET comment_count = comment_count+1 WHERE id=?", comment.VideosId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("PushComment: %s", err)
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("提交事务: %s", err)
	}

	return nil
}

func DeleteComment(commentId, videoId int64) error {
	// 开始数据库事务
	tx := database.GetDB().Begin()

	err := tx.Exec("DELETE FROM comments WHERE id = ?", commentId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("DeleteComment: %s", err)
	}

	err = tx.Exec("UPDATE videos SET comment_count = comment_count-1 WHERE id=? AND comment_count>0", videoId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("DeleteComment: %s", err)
	}

	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("提交事务: %s", err)
	}

	return nil
}

// GetCommentList 获取评论列表
func GetCommentList(videoId int64, comments *[]*Comment) error {
	if comments == nil {
		return fmt.Errorf("GetCommentList %s", nil)
	}
	err := database.GetDB().Raw("SELECT * FROM comments WHERE videos_id=?", videoId).Scan(comments).Error
	if err != nil {
		return fmt.Errorf("GetCommentList: %s", err)
	}
	return nil
}
