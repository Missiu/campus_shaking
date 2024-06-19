package services

import (
	"fmt"
	"new/models"
	"time"
)

type SendMessageFollow struct {
	userID   int64
	toUserID int64
	content  string

	message *models.Message
}

func SendMessage(userID, toUserID int64, content string) error {
	return NewSendMessageFlow(userID, toUserID, content).Do()
}
func NewSendMessageFlow(userID, toUserID int64, content string) *SendMessageFollow {
	return &SendMessageFollow{userID: userID, toUserID: toUserID, content: content}
}

type ChatHistoryFollow struct {
	userID   int64
	toUserID int64
	content  string

	messages []*models.Message
}

func ChatHistory(userID, toUserID int64) ([]*models.Message, error) {
	return NewChatHistoryFlow(userID, toUserID).Do()
}
func NewChatHistoryFlow(userID, toUserID int64) *ChatHistoryFollow {
	return &ChatHistoryFollow{userID: userID, toUserID: toUserID}
}

// Do 发送信息
func (p *SendMessageFollow) Do() error {
	p.message = &models.Message{
		Content:    p.content,
		CreateTime: time.Now().Unix(),
		FromUserID: p.userID,
		ToUserID:   p.toUserID,
	}
	err := models.SendMessage(p.message)
	if err != nil {
		return fmt.Errorf("SendMessage : %s", err)
	}
	return nil
}

// Do 获取消息列表
func (g *ChatHistoryFollow) Do() ([]*models.Message, error) {
	err := models.ChatHistory(g.userID, g.toUserID, &g.messages)
	if err != nil {
		return nil, fmt.Errorf("ChatHistory : %s", err)
	}
	return g.messages, nil
}
