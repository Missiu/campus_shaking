package services

import (
	"errors"
	"fmt"
	"new/models"
	"new/utils"
)

type UserToken struct {
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserRegisterFlow struct {
	username string
	password string
	userid   int64
	token    string

	data *UserToken
}

type UserLoginFlow struct {
	username string
	password string
	userid   int64
	token    string

	data *UserToken
}

func UserRegister(username, password string) (*UserToken, error) {
	return NewUserRegisterFlow(username, password).Do()
}
func NewUserRegisterFlow(username string, password string) *UserRegisterFlow {
	return &UserRegisterFlow{username: username, password: password}
}

func UserLogin(username, password string) (*UserToken, error) {
	return NewUserLoginFlow(username, password).Do()
}
func NewUserLoginFlow(username string, password string) *UserLoginFlow {
	return &UserLoginFlow{username: username, password: password}
}

// Do 注册操作
func (p *UserRegisterFlow) Do() (*UserToken, error) {
	// 检查数据
	if len(p.password) < 8 {
		return nil, errors.New("密码太短，需大于8位")
	}
	if p.password == "" {
		return nil, errors.New("密码为空")
	}
	if len(p.username) > 100 {
		return nil, errors.New("用户名过长")
	}
	if p.username == "" {
		return nil, errors.New("用户名为空")
	}
	if models.IsUserExist(p.username) {
		return nil, errors.New("该用户名已经存在")
	}

	// 准备把数据打包
	userInfo := &models.UserInfo{Name: p.username}
	userLogin := &models.UserLogin{Username: p.username, Password: p.password, UserInfo: userInfo}
	// 用户信息传入models
	err := models.AddUserInfo(userInfo)
	if err != nil {
		return nil, fmt.Errorf("AddUserInfo 失败: %s", err)
	}
	// 登录信息传入models
	err = models.Register(userLogin)
	if err != nil {
		return nil, fmt.Errorf("register 失败: %s", err)
	}
	// token生成
	token, err := utils.GenerateToken(userLogin)
	if err != nil {
		return nil, fmt.Errorf("GenerateToken 失败: %s", err)
	}

	p.token = token
	p.userid = userInfo.Id
	p.data = &UserToken{
		UserId: p.userid,
		Token:  p.token,
	}
	return p.data, nil
}

// Do 用户登录
func (p *UserLoginFlow) Do() (*UserToken, error) {
	// 数据检查

	login := &models.UserLogin{}
	err := models.Login(p.username, p.password, login)
	if err != nil {
		return nil, fmt.Errorf("login 失败: %s", err)
	}
	token, err := utils.GenerateToken(login)
	if err != nil {
		return nil, fmt.Errorf("GenerateToken 失败: %s", err)
	}
	p.token = token
	p.userid = login.ID
	p.data = &UserToken{
		UserId: p.userid,
		Token:  p.token,
	}
	return p.data, nil
}
