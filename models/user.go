package models

import (
	"fmt"
	"new/database"
)

type UserLogin struct {
	ID         int64     `gorm:"primary_key"`
	Username   string    `gorm:"primary_key"`
	UserInfoID int64     // 使用驼峰式命名法
	Password   string    `gorm:"size:200;notnull"`
	UserInfo   *UserInfo `gorm:"foreignKey:UserInfoID"` // 使用驼峰式命名法
}

type UserInfo struct {
	Id              int64       `json:"id,omitempty" gorm:"primary_key"`      // 用户id
	FollowCount     int64       `json:"follow_count,omitempty"`               // 关注总数
	FollowerCount   int64       `json:"follower_count,omitempty"`             // 粉丝总数
	IsFollow        bool        `json:"is_follow,omitempty"`                  // true-已关注，false-未关注
	Name            string      `json:"name,omitempty"`                       // 用户名称
	Avatar          string      `json:"avatar,omitempty"`                     // 头像
	BackgroundImage string      `json:"background_image,omitempty"`           // 用户个人页顶部大图
	Signature       string      `json:"signature,omitempty"`                  // 个人简介
	TotalFavorited  int64       `json:"total_favorited,omitempty"`            // 获赞数量
	WorkCount       int64       `json:"work_count,omitempty"`                 // 作品数
	FavoriteCount   int64       `json:"favorite_count,omitempty"`             // 喜欢总数
	Follow          []*UserInfo `json:"-" gorm:"many2many:user_relations"`    // 用户之间相互关注，多对多通过外表user_relations联系
	FavorVideos     []*Videos   `json:"-" gorm:"many2many:user_favor_videos"` // 用户点赞视频
	Comments        []*Comment  `json:"-" gorm:"foreignKey:UserId"`           // 用户评论，设置用户id为外键
	Videos          []*Videos   `json:"-" gorm:"foreignKey:AuthorId"`         // 投稿视频，设置作者id为外键
}

// TableName 设置表名
func (U UserLogin) TableName() string {
	return "user_login"
}

func (u UserInfo) TableName() string {
	return "user_info"
}

// Register 注册 创建数据
func Register(user *UserLogin) error {
	if user == nil {
		return fmt.Errorf("user : %s", nil)
	}
	return database.GetDB().Create(&user).Error
}

// AddUserInfo 添加信用户信息
func AddUserInfo(info *UserInfo) error {
	if info == nil {
		return fmt.Errorf("info : %s", nil)
	}
	return database.GetDB().Create(&info).Error
}

// Login 登录
func Login(username string, password string, login *UserLogin) error {
	err := database.GetDB().Raw("select * from user_login where username=? and password=?", username, password).Scan(login).Error
	if err != nil {
		return fmt.Errorf("login : %s", err)
	}
	if login.ID == 0 {
		return fmt.Errorf("login : %s", err)
	}
	return nil
}

// IsUserExist 判断用户名是否存在
func IsUserExist(username string) bool {
	var count int
	err := database.GetDB().Raw("select count(*) from user_login WHERE username = ?", username).Scan(&count).Error
	if err != nil {
		panic(fmt.Errorf("login : %s", err))
	}
	return count > 0
}

// GetUserInfo 通过id获取用户信息
func GetUserInfo(userId int64, userinfo *UserInfo) error {
	err := database.GetDB().Raw("SELECT * FROM user_info WHERE id=?", userId).Scan(userinfo).Error
	if err != nil {
		return fmt.Errorf("GetUserInfo : %s", err)
	}
	if userinfo.Id == 0 {
		return fmt.Errorf("GetUserInfo : %s", err)
	}
	return nil
}

// UpdateWorkCount 更新作品信息,视频被喜欢的个数
func UpdateWorkCount(id int64) error {
	var count int64
	err := database.GetDB().Raw("SELECT COUNT(author_id)  FROM videos where author_id=?", id).Scan(&count).Error
	if err != nil {
		return fmt.Errorf("UpdateWorkCount %s", err)
	}

	var fc int64
	err = database.GetDB().Raw("SELECT SUM(favorite_count) FROM videos where author_id=?", id).Scan(&fc).Error
	if err != nil {
		return fmt.Errorf("UpdateWorkCount %s", err)
	}
	//fmt.Println(fmt.Sprintf("count: %d", count))
	// 开始数据库事务
	tx := database.GetDB().Begin()

	// 更新用户获赞数量
	err = tx.Exec("UPDATE user_info SET total_favorited=? WHERE id=?", fc, id).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UpdateWorkCount 作品数量更新失败: %s", err)
	}
	// 更新用户作品数量
	err = tx.Exec("UPDATE user_info SET work_count=? WHERE id=?", count, id).Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UpdateWorkCount 作品数量更新失败: %s", err)
	}
	// 提交事务
	err = tx.Commit().Error
	if err != nil {
		// 发生错误，回滚事务并返回错误
		tx.Rollback()
		return fmt.Errorf("UpdateWorkCount 提交事务失败: %s", err)
	}

	return nil
}
