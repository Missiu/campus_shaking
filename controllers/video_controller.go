package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"new/services"
	"new/utils"
	"path/filepath"
)

// GetVideoFeed 获取源视频
func GetVideoFeed(c *gin.Context) {
	token, ok := c.GetQuery("token")
	//判断是否为登录状态
	//未登录
	if !ok {
		// 获取视频列表
		videoList, err := services.GetVideoFeed(0)
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{
				StatusCode: 1,
				StatusMsg:  "C videoList 获取失败",
			})
			panic(fmt.Errorf("C videoList 获取失败:%s", err))
		}
		c.JSON(http.StatusOK, utils.FeedVideoResponse{
			Response: utils.Response{StatusCode: 0},
			NextTime: videoList.NextTime,
			Videos:   videoList.Videos,
		})
		return
	}
	// 已登录
	if ok {
		// 解析token
		parsedData, ok := utils.ParseToken(token)
		if !ok {
			c.JSON(http.StatusOK, utils.Response{
				StatusCode: 1,
				StatusMsg:  "token 解析失败",
			})
			panic(fmt.Errorf("C token 解析失败"))
		}

		videoList, err := services.GetVideoFeed(parsedData.ID)
		if err != nil {
			c.JSON(http.StatusOK, utils.Response{
				StatusCode: 1,
				StatusMsg:  "C videoList 获取失败",
			})
			panic(fmt.Errorf("C videoList 获取失败:%s", err))
		}
		c.JSON(http.StatusOK, utils.FeedVideoResponse{
			Response: utils.Response{StatusCode: 0},
			NextTime: videoList.NextTime,
			Videos:   videoList.Videos,
		})
	}

}

// PushVideo 上传视频
func PushVideo(c *gin.Context) {
	// 获取并解析token，得到userid
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}
	// 解析token
	parsedData, ok := utils.ParseToken(tokenStr)
	if !ok {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "C token 解析失败",
		})
		return
	}
	// 获取标题
	title := c.PostForm("title")
	// 获取视频文件（未进行格式判断）
	file, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "C 获取文件失败",
		})
		return
	}

	// 按时间戳+userId生成唯一文件名
	name := utils.NewFileName(parsedData.ID)
	// 文件唯一命名
	newVideoName := name + ".mp4"
	newimgName := name + ".jpg"
	// 文件储存路径
	videoFilePath := filepath.Join("public", newVideoName)
	// 封面储存路径
	coverFilePath := filepath.Join("public", newimgName)
	// 保存视频文件
	savevideoErr := c.SaveUploadedFile(file, videoFilePath)
	if savevideoErr != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "C 文件保存错误",
		})
		return
	}
	outvideoFilePath := "public/zip" + newVideoName
	err = utils.VideoZIP(videoFilePath, outvideoFilePath)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "C VideoZIP",
		})
		return
	}

	// 视频截屏
	err = utils.ScreenShot(outvideoFilePath, coverFilePath)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "C 视频截屏错误",
		})
		return
	}
	// 获取保存路径
	videoUrl := utils.GetVideoURL("zip" + newVideoName)
	CoverUrl := utils.GetImgURL(newimgName)
	// 保存在数据库
	err = services.SavedVide(parsedData.ID, title, videoUrl, CoverUrl)
	if err != nil {
		c.JSON(http.StatusOK, utils.Response{
			StatusCode: 1,
			StatusMsg:  "数据库文件保存错误",
		})
		return
	}
	c.JSON(http.StatusOK, utils.Response{StatusCode: 0, StatusMsg: "上传成功"})
}

// GetVideoList 获取发布的视频列表
func GetVideoList(c *gin.Context) {
	// 获取token
	tokenStr := c.Query("token")
	if tokenStr == "" {
		tokenStr = c.PostForm("token")
	}
	// 解析token
	tokenStruck, ok := utils.ParseToken(tokenStr)
	if !ok {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{
				StatusCode: 1,
				StatusMsg:  "ParseToken ERROR",
			},
		})
		return
	}
	// 根据id调用视频列表
	videoList, err := services.GetAuthorVideoList(tokenStruck.ID)
	if err != nil {
		c.JSON(http.StatusOK, utils.InfoResponse{
			Response: utils.Response{
				StatusCode: 1,
				StatusMsg:  "GetAuthorVideoList ERROR",
			},
		})
		return
	}
	c.JSON(http.StatusOK, utils.FeedVideoResponse{
		Response: utils.Response{StatusCode: 0},
		Videos:   videoList,
	})
}
