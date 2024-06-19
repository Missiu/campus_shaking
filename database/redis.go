package database

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

// 创建一个 context 上下文对象，用于在 Redis 操作中传递上下文信息，如取消操作等。
var ctx = context.Background()

// 定义了一个 Redis 客户端对象，用于与 Redis 服务器进行交互。
var client *redis.Client

func InitRedis() error {
	client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", viper.GetString("redis.ip"), viper.GetInt64("redis.port")), // Redis 服务器地址和端口
		Password: viper.GetString("redis.password"),                                               // Redis 密码
		DB:       viper.GetInt("redis.database"),                                                  // Redis 数据库索引
	})
	//pong, err := client.Ping(ctx).Result()
	//if err != nil {
	//	return fmt.Errorf("链接错误 %s", err)
	//}
	//fmt.Println(pong)
	return nil
}

func GetRDB() *redis.Client {
	if client == nil {
		err := InitViper()
		if err != nil {
			err = InitViperOtherPath()
			if err != nil {
				panic(fmt.Errorf("初始化Viper失败: %s", err))
			}
		}
		err = InitRedis()
		if err != nil {
			panic(fmt.Errorf("初始化redis缓存失败: %s", err))
		}
	}
	return client
}

// UpdateLikeStatus 更新点赞状态state:true为点赞，false为取消点赞
func UpdateLikeStatus(userId int64, videoId int64, state bool) {
	// like:userId
	key := fmt.Sprintf("%s:%d", "like", userId)
	if state {
		//videoId添加到键名为key的集合中
		GetRDB().SAdd(ctx, key, videoId)
		return
	} else {
		//将videoId从键名为key的集合中移除
		GetRDB().SRem(ctx, key, videoId)
	}
}

// GetLikeStatus 得到点赞状态
func GetLikeStatus(userId int64, videoId int64) bool {
	key := fmt.Sprintf("%s:%d", "like", userId)
	result, _ := GetRDB().SIsMember(ctx, key, videoId).Result()
	return result
}

// UpdateFollowStatus 更新关注状态
func UpdateFollowStatus(userId int64, followId int64, state bool) {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	if state {
		GetRDB().SAdd(ctx, key, followId)
		return
	}
	GetRDB().SRem(ctx, key, followId)
}

// GetFollowStatus 得到关注状态
func GetFollowStatus(userId int64, followId int64) bool {
	key := fmt.Sprintf("%s:%d", "relation", userId)
	result, _ := GetRDB().SIsMember(ctx, key, followId).Result()
	return result
}
