package services

import (
	"fmt"
	"new/database"
	"new/models"
	"new/utils"
)

type LikeFlow struct {
	userId     int64
	videoId    int64
	actionType int64
}

type GetLikeVideoListFlow struct {
	userId int64
	videos []*models.Videos
}

func Like(userId, videoId, actionType int64) error {
	return NewLikeFlow(userId, videoId, actionType).Do()
}
func NewLikeFlow(userId, videoId, actionType int64) *LikeFlow {
	return &LikeFlow{userId: userId, videoId: videoId, actionType: actionType}
}

func LikeVideoList(userId int64) ([]*models.Videos, error) {
	return NewLikeVideoListFlow(userId).Do()
}
func NewLikeVideoListFlow(userId int64) *GetLikeVideoListFlow {
	return &GetLikeVideoListFlow{userId: userId}
}

// Do 点赞操作
func (p *LikeFlow) Do() error {
	switch p.actionType {
	case 1:
		err := models.LikeOperation(p.userId, p.videoId)
		if err != nil {
			return fmt.Errorf("点赞操作 : %s", err)
		}
		// 缓存点赞状态
		database.UpdateLikeStatus(p.userId, p.videoId, true)
	case 2:
		err := models.UnLike(p.userId, p.videoId)
		if err != nil {
			return fmt.Errorf("取消点赞操作 : %s", err)
		}
		database.UpdateLikeStatus(p.userId, p.videoId, false)
	default:
		panic(fmt.Sprintf("未定义的操作"))
	}
	return nil
}

// Do 获取点赞过的视频列表
func (g *GetLikeVideoListFlow) Do() ([]*models.Videos, error) {
	// 获取点赞视频信息
	err := models.GetVideoListByUserId(g.userId, &g.videos)
	if err != nil {
		return nil, fmt.Errorf("GetVideoListByUserId : %s", err)
	}
	// 填充作者信息
	err = utils.VideoStatusSup(g.userId, g.videos)
	if err != nil {
		return nil, fmt.Errorf("VideoStatusSup : %s", err)
	}
	return g.videos, nil
}
