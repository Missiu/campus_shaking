package services

import (
	"fmt"
	"new/database"
	"new/models"
)

type FollowFlow struct {
	userId     int64
	toUserId   int64
	actionType int64
}

type FollowListFlow struct {
	userId int64

	userList []*models.UserInfo
}

type FollowerListFlow struct {
	userId int64

	userList []*models.UserInfo
}

func Follow(userId, toUserId, actionType int64) error {
	return NewFollowFlow(userId, toUserId, actionType).Do()
}
func NewFollowFlow(userId, toUserId, actionType int64) *FollowFlow {
	return &FollowFlow{userId: userId, toUserId: toUserId, actionType: actionType}
}

func FollowerList(userId int64) ([]*models.UserInfo, error) {
	return NewFollowerListFlow(userId).Do()
}
func NewFollowerListFlow(userId int64) *FollowerListFlow {
	return &FollowerListFlow{userId: userId}
}

func FollowList(userId int64) ([]*models.UserInfo, error) {
	return NewFollowListFlow(userId).Do()
}
func NewFollowListFlow(userId int64) *FollowListFlow {
	return &FollowListFlow{userId: userId}
}

// Do 关注操作
func (p FollowFlow) Do() error {
	switch p.actionType {
	case 1:
		if p.userId == p.toUserId {
			return fmt.Errorf("不能关注自己哦！")
		}
		err := models.FollowOperation(p.userId, p.toUserId)
		if err != nil {
			return fmt.Errorf("关注操作 : %s", err)
		}
		//缓存关注状态
		database.UpdateFollowStatus(p.userId, p.toUserId, true)
	case 2:
		err := models.UnFollow(p.userId, p.toUserId)
		if err != nil {
			return fmt.Errorf("取消关注操作 : %s", err)
		}
		database.UpdateFollowStatus(p.userId, p.toUserId, false)
	default:
		panic(fmt.Sprintf("未定义的操作"))
	}
	return nil
}

// Do 获取粉丝列表
func (g *FollowerListFlow) Do() ([]*models.UserInfo, error) {
	err := models.GetFollowerList(g.userId, &g.userList)
	if err != nil {
		panic(fmt.Errorf("err: %s", err))
	}
	return g.userList, nil
}

// Do 获取关注列表
func (g *FollowListFlow) Do() ([]*models.UserInfo, error) {
	err := models.GetFollowList(g.userId, &g.userList)
	if err != nil {
		return nil, fmt.Errorf("GetFollowList : %s", err)
	}
	// 信息添加
	for i := 0; i < len(g.userList); i++ {
		info, err := UserInfo(g.userList[i].Id)
		if err != nil {
			return nil, fmt.Errorf("UserInfo : %s", err)
		}
		g.userList[i].Avatar = info.Avatar
	}
	return g.userList, nil
}

type FriendListFlow struct {
	userId int64

	userList []*models.UserInfo
}

func FriendList(userId int64) ([]*models.UserInfo, error) {
	return NewFriendListFlow(userId).Do()
}
func NewFriendListFlow(userId int64) *FriendListFlow {
	return &FriendListFlow{userId: userId}
}

// Do 获取好友列表
func (g *FriendListFlow) Do() ([]*models.UserInfo, error) {
	err := models.GetFriendList(g.userId, &g.userList)
	if err != nil {
		return nil, fmt.Errorf("GetFollowList : %s", err)
	}
	// 信息添加
	for i := 0; i < len(g.userList); i++ {
		info, err := UserInfo(g.userList[i].Id)
		if err != nil {
			return nil, fmt.Errorf("UserInfo : %s", err)
		}
		g.userList[i].Avatar = info.Avatar
	}
	return g.userList, nil
}
