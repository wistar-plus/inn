package redis

import (
	"inn/internal/message/repository"
)

var key = "user.topic"

type UserTopicRepository struct{}

func NewConnAddrRepository() repository.IUserTopicRepository {
	return &UserTopicRepository{}
}

func (utr *UserTopicRepository) Get(uid uint64) string {
	res, _ := db.Do("HGET", key, uid).Result()
	topic, ok := res.(string)
	if !ok {
		return ""
	}
	return topic
}
