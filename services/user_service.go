package services

import (
	"fmt"
	"new/models"
)

type UserInfoFlow struct {
	userId   int64
	userinfo models.UserInfo
}

func UserInfo(userId int64) (models.UserInfo, error) {
	return NewUserInfoFlow(userId).Do()
}
func NewUserInfoFlow(userId int64) *UserInfoFlow {
	return &UserInfoFlow{userId: userId}
}

// Do 用户信息填充
func (g *UserInfoFlow) Do() (models.UserInfo, error) {
	err := models.GetUserInfo(g.userId, &g.userinfo)
	if err != nil {
		return models.UserInfo{}, fmt.Errorf("GetUserInfo : %s", err)
	}
	g.userinfo.Avatar = "https://pic2.zhimg.com/v2-76c91802ec47c452fb1155019e86725d_l.jpg?source=172ae18b"
	g.userinfo.BackgroundImage = "https://www.magedu.com/wp-content/uploads/2020/09/5ea652640838e8fc12000675-1.jpg"
	g.userinfo.Signature = "我亦无他，惟手熟尔"

	return g.userinfo, nil
}
