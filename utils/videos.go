package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"new/database"
	"new/models"
	"os/exec"
	"strconv"
	"time"
)

func NewFileName(userID int64) string {
	return strconv.FormatInt(userID+time.Now().Unix(), 10)
}

func GetVideoURL(fileName string) string {
	return fmt.Sprintf("http://%s:%d/static/%s", viper.GetString("server.ip"), viper.GetInt("server.port"), fileName)
}

func GetImgURL(fileName string) string {
	return fmt.Sprintf("http://%s:%d/static/%s", viper.GetString("server.ip"), viper.GetInt("server.port"), fileName)
}

// ScreenShot 截取一帧画面
func ScreenShot(videoFilePath, coverFilePath string) error {
	var ffmpegPath string
	if viper.GetString("ffmpeg.path") == "" {
		ffmpegPath = "./ffmpeg"
	} else {
		ffmpegPath = viper.GetString("ffmpeg.path") // 指定 FFmpeg 的路径
	}
	cmd := exec.Command(
		ffmpegPath,
		"-i", videoFilePath,
		"-ss", "00:00:01.000", // 从视频第一秒开始截取
		"-vframes", "1", // 只截取一帧
		coverFilePath,
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("视频截屏失败: %v", err)
	}
	return nil
}

// VideoZIP 视频画质压缩
func VideoZIP(videoFilePath, outVideoFilePath string) error {
	var ffmpegPath string
	if viper.GetString("ffmpeg.path") == "" {
		ffmpegPath = "./ffmpeg"
	} else {
		ffmpegPath = viper.GetString("ffmpeg.path") // 指定 FFmpeg 的路径
	}
	cmd := exec.Command(
		ffmpegPath,
		"-i", videoFilePath,
		"-c:v", "libx264", // 设置视频编码器为 libx264
		"-crf", "26", // 设置压缩质量，值越小质量越高（18-28）
		"-preset", "medium", // 设置编码速度和压缩效率的预设值
		"-c:a", "aac", // 设置音频编码器为 aac
		"-b:a", "128k", // 设置音频比特率为 128k
		"-y", // 自动覆盖输出文件
		outVideoFilePath,
	)
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("视频画质压缩: %v", err)
	}
	return nil
}

// VideoStatusSup 视频信息插入
func VideoStatusSup(userId int64, videos []*models.Videos) error {
	for i := 0; i < len(videos); i++ {
		// 作者信息插入,针对源视频获取，传入的为用户(非作者)id，主要进行点赞，关注等信息的插入
		var userInfo models.UserInfo
		err1 := models.GetUserInfo(videos[i].AuthorId, &userInfo)
		if err1 != nil {
			continue
		}
		// 判断是否登录
		if userId > 0 {
			// 插入喜欢信息
			(videos)[i].IsFavorite = database.GetLikeStatus(userId, videos[i].ID)
			// 插入关注信息
			userInfo.IsFollow = database.GetFollowStatus(userId, videos[i].AuthorId)
			// 插入作品数量
			err := models.UpdateWorkCount(videos[i].AuthorId)
			if err != nil {
				return err
			}
		}
		(videos)[i].Author = userInfo
	}
	return nil
}
