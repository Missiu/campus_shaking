package models

import (
	"fmt"
	"new/database"
)

func LikeOperation(userId, videoId int64) error {
	// 开始数据库事务
	tx := database.GetDB().Begin()

	// 更新视频点赞数据
	err := tx.Exec("UPDATE videos SET favorite_count=favorite_count+1 WHERE id=?", videoId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("LikeOperation: %s", err)
	}

	// 更新视频点赞数据
	err = tx.Exec("UPDATE user_info SET favorite_count=favorite_count+1 WHERE id=?", userId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("LikeOperation: %s", err)
	}

	// 插入喜欢信息的外表
	err = tx.Exec("INSERT INTO user_favor_videos (user_info_id,videos_id) VALUES (?,?)", userId, videoId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("LikeOperation: %s", err)
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

func UnLike(userId, videoId int64) error {
	// 开始数据库事务
	tx := database.GetDB().Begin()

	// 喜欢的数目减1
	err := tx.Exec("UPDATE videos SET favorite_count=favorite_count-1 WHERE id=? AND favorite_count>0", videoId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UnLike: %s", err)
	}

	err = tx.Exec("UPDATE user_info SET favorite_count=favorite_count-1 WHERE id=?  AND favorite_count>0", userId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("LikeOperation: %s", err)
	}

	// 删除外表数据
	err = tx.Exec("DELETE FROM user_favor_videos WHERE user_info_id=? AND videos_id=?", userId, videoId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UnLike: %s", err)
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
