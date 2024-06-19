package models

import (
	"fmt"
	"new/database"
)

func FollowOperation(userId, toUserId int64) error {
	// 开始数据库事务
	tx := database.GetDB().Begin()

	// 更新自己的关注数据
	err := tx.Exec("UPDATE user_info SET follow_count=follow_count+1 WHERE id=?", userId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("FollowOperation: %s", err)
	}
	// 粉丝数量增加
	err = tx.Exec("UPDATE user_info SET follower_count=follower_count+1 WHERE id=?", toUserId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("FollowOperation: %s", err)
	}

	// 插入关注信息的外表
	err = tx.Exec("INSERT INTO user_relations (user_info_id,follow_id) VALUES (?,?)", userId, toUserId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("FollowOperation: %s", err)
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

// UnFollow 取消关注
func UnFollow(userId, toUserId int64) error {
	// 开始数据库事务
	tx := database.GetDB().Begin()

	// 关注数量减1
	err := tx.Exec("UPDATE user_info SET follow_count=follow_count-1 WHERE id=? AND follow_count>0", userId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UnFollow: %s", err)
	}

	// 粉丝数量减少1
	err = tx.Exec("UPDATE user_info SET follower_count=follower_count-1 WHERE id=? AND follower_count>0", toUserId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UnFollow: %s", err)
	}

	// 删除关注的的外表信息
	err = tx.Exec("DELETE FROM user_relations WHERE user_info_id=? AND follow_id=?", userId, toUserId).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UnFollow: %s", err)
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

// GetFollowList 获取关注列表
func GetFollowList(userId int64, list *[]*UserInfo) error {
	// 获取关注信息
	err := database.GetDB().Raw("select u.* from user_relations r inner join user_info u on r.follow_id=u.id where r.user_info_id=?", userId).
		Scan(list).Error
	if err != nil {
		return fmt.Errorf("GetFollowList: %s", err)
	}
	return nil
}

// GetFollowerList 获取粉丝列表
func GetFollowerList(userId int64, list *[]*UserInfo) error {
	// 获取粉丝信息
	err := database.GetDB().Raw("select u.* from user_relations r inner join user_info u on r.user_info_id=u.id where r.follow_id=?", userId).
		Scan(list).Error
	if err != nil {
		return fmt.Errorf("GetFollowerList: %s", err)
	}
	return nil
}

// GetFriendList 获取好友列表,相互关注才可以聊天
func GetFriendList(userId int64, list *[]*UserInfo) error {
	// 相互关注的对方的id信息
	var ID int64
	// 获取关注信息
	err := database.GetDB().Raw("select follow_id from user_relations where user_info_id=? and follow_id in (select user_info_id from user_relations where follow_id=?)", userId, userId).
		Scan(&ID).Error
	if err != nil {
		return fmt.Errorf("GetFriendList %s", err)
	}
	if ID == 0 {
		return fmt.Errorf("暂且无与你相互关注的用户")
	}
	err = database.GetDB().Raw("select * from user_info where id=?", ID).Scan(list).Error
	if err != nil {
		return fmt.Errorf("GetFriendList %s", err)
	}
	return nil
}
