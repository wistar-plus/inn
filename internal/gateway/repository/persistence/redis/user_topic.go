package redis

import (
	"inn/internal/gateway/repository"

	"github.com/prometheus/common/log"
)

var key = "user.topic"

type UserTopicRepository struct{}

func NewUserTopicRepository() repository.IUserTopicRepository {
	return &UserTopicRepository{}
}

func (utr *UserTopicRepository) Add(uid uint64, topic string) (err error) {
	if _, err = db.Do("HSET", key, uid, topic).Result(); err != nil {
		log.Error("db.Do(HSET %s %s %s) error: %v", key, uid, topic, err)
	}
	return
}

func (utr *UserTopicRepository) Del(uid uint64) (err error) {
	if _, err = db.Do("HDEL", key, uid).Result(); err != nil {
		log.Error("db.Do(HDEL %s %s) error: %v", key, uid, err)
	}
	return
}
