package models

import (
	"fmt"
	"new/database"
	"time"
)

type Videos struct {
	ID            int64       `json:"id,omitempty"`                          //视频id
	AuthorId      int64       `json:"author_id"`                             // 与作者id对应
	Author        UserInfo    `json:"author" gorm:"-"`                       // 视频的作者信息
	PlayUrl       string      `json:"play_url,omitempty"`                    // 视频播放地址
	CoverUrl      string      `json:"cover_url,omitempty"`                   // 视频封面地址
	FavoriteCount int64       `json:"favorite_count"`                        // 视频的点赞总数
	CommentCount  int64       `json:"comment_count"`                         // 视频的评论总数
	IsFavorite    bool        `json:"is_favorite"`                           // true-已点赞，false-未点赞
	Title         string      `json:"title"`                                 // 视频标题
	Users         []*UserInfo `json:"-" gorm:"many2many:user_favor_videos;"` // 用户点赞视频
	Comments      []*Comment  `json:"-"`                                     // 视频评论
	CreatedAt     time.Time   `json:"-"`                                     // 创建时间
	UpdatedAt     time.Time   `json:"-"`                                     // 更新时间
}

func (v Videos) TableName() string {
	return "videos"
}

// AddVideo 添加视频
func AddVideo(video *Videos) error {
	if video == nil {
		return fmt.Errorf("AddVideo : %s", nil)
	}
	err := database.GetDB().Create(video).Error
	if err != nil {
		return fmt.Errorf("AddVideo 添加视频失败: %s", err)
	}

	return nil
}

// GetVideoFeed 获取视频Feed
func GetVideoFeed(latestTime time.Time, videoList *[]*Videos) error {
	if videoList == nil {
		return fmt.Errorf("GetVideoFeed : %s", nil)
	}
	err := database.GetDB().Raw("select * from videos where created_at<? order by created_at desc LIMIT 30", latestTime).Scan(videoList).Error
	if err != nil {
		return fmt.Errorf("GetVideoFeed 获取失败: %v", err)
	}
	return nil
}

// GetVideoListByAuthorId 根据作者id获取视频列表
func GetVideoListByAuthorId(userId int64, videoList *[]*Videos) error {
	if videoList == nil {
		return fmt.Errorf("GetVideoListByAuthorId : %s", nil)
	}
	err := database.GetDB().Raw("select * from videos where author_id=?", userId).Scan(videoList).Error
	if err != nil {
		return err
	}

	return nil
}

// GetVideoListByUserId 通过用户id获取视频信息
func GetVideoListByUserId(userId int64, videoList *[]*Videos) error {
	if videoList == nil {
		return fmt.Errorf("GetVideoListByUserId : %s", nil)
	}
	err := database.GetDB().Raw("select v.* from user_favor_videos uf right outer join videos v on uf.videos_id = v.id where uf.user_info_id = ?", userId).Scan(videoList).Error
	if err != nil {
		return fmt.Errorf("GetVideoListByUserId : %s", err)
	}
	return nil
}
