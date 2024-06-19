package services

import (
	"fmt"
	"new/models"
	"new/utils"
	"time"
)

type VideoListFlow struct {
	nextTime int64
	userid   int64
	videos   []*models.Videos

	feedVideo *utils.FeedVideoResponse
}

type SavedVideFlow struct {
	videoPath string
	coverPath string
	title     string
	authorId  int64
	video     *models.Videos
}
type GetAuthorVideoListFlow struct {
	userId int64
	videos []*models.Videos
}

func GetVideoFeed(userId int64) (*utils.FeedVideoResponse, error) {
	return NewGetVideoFeedFlow(userId).Do()
}
func NewGetVideoFeedFlow(userId int64) *VideoListFlow {
	return &VideoListFlow{userid: userId}
}

func SavedVide(authorId int64, title, videoPath, coverPath string) error {
	return NewSavedVideFlow(authorId, title, videoPath, coverPath).Do()
}
func NewSavedVideFlow(authorId int64, title, videoPath, coverPath string) *SavedVideFlow {
	return &SavedVideFlow{authorId: authorId, title: title, videoPath: videoPath, coverPath: coverPath}
}

func GetAuthorVideoList(userId int64) ([]*models.Videos, error) {
	return NewGetAuthorVideoListFlow(userId).Do()
}
func NewGetAuthorVideoListFlow(userId int64) *GetAuthorVideoListFlow {
	return &GetAuthorVideoListFlow{userId: userId}
}

// Do 获取源视频
func (g *VideoListFlow) Do() (*utils.FeedVideoResponse, error) {
	// 数据操作
	g.nextTime = time.Now().Unix()
	// 按时间倒叙返回视频列表 后续可改进算法
	err := models.GetVideoFeed(time.Now(), &g.videos)
	if err != nil {
		return nil, fmt.Errorf("GetVideoFeed : %s", err)
	}
	// 视频信息填充
	err = utils.VideoStatusSup(g.userid, g.videos)
	if err != nil {
		return nil, fmt.Errorf("VideoStatusSup : %s", err)
	}
	// 信息打包
	g.feedVideo = &utils.FeedVideoResponse{
		NextTime: g.nextTime,
		Videos:   g.videos,
	}
	return g.feedVideo, nil
}

// Do 上传保存视频
func (p *SavedVideFlow) Do() error {
	// 打包数据
	p.video = &models.Videos{
		AuthorId: p.authorId,
		PlayUrl:  p.videoPath,
		CoverUrl: p.coverPath,
		Title:    p.title,
	}
	err := models.AddVideo(p.video)
	if err != nil {
		return fmt.Errorf("AddVideo : %s", err)
	}
	return nil
}

// Do 作者的视频列表获取
func (g *GetAuthorVideoListFlow) Do() ([]*models.Videos, error) {
	err := models.GetVideoListByAuthorId(g.userId, &g.videos)
	if err != nil {
		return nil, fmt.Errorf("GetVideoListByAuthorId : %s", err)
	}

	// 视频信息插入
	err = utils.VideoStatusSup(g.userId, g.videos)
	if err != nil {
		return nil, fmt.Errorf("GetVideoListByAuthorId : %s", err)
	}
	return g.videos, nil
}
