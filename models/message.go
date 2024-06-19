package models

import (
	"fmt"
	"new/database"
)

type Message struct {
	ID         int64  `json:"id"`                             // 消息id
	Content    string `json:"content"`                        // 消息内容
	CreateTime int64  `json:"create_time" gorm:"uniqueIndex"` // 消息发送时间 将时间戳字段设置为唯一索引
	FromUserID int64  `json:"from_user_id"`                   // 消息发送者id
	ToUserID   int64  `json:"to_user_id"`                     // 消息接收者id
}

// TableName 设置表名
func (U Message) TableName() string {
	return "messages"
}

func SendMessage(message *Message) error {
	if message == nil {
		return fmt.Errorf("SendMessage %s", nil)
	}
	err := database.GetDB().Create(&message).Error
	if err != nil {
		return fmt.Errorf("SendMessage : %s", err)
	}
	return nil
}

func ChatHistory(userid, toUserId int64, messages *[]*Message) error {
	err := database.GetDB().Raw("SELECT * FROM messages WHERE (from_user_id=? and to_user_id=?) or (from_user_id=? and to_user_id=?)", userid, toUserId, toUserId, userid).
		Scan(&messages).Error
	if err != nil {
		return fmt.Errorf("ChatHistory : %s", err)
	}
	return nil
}
