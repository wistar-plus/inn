package redis

import (
	"inn/internal/message/repository"
	"strconv"
)

type UnreadRepository struct{}

func NewUnreadRepository() repository.IUnreadRepository {
	return &UnreadRepository{}
}

func (ur *UnreadRepository) IncrementUnreadBy(uid, otherUid uint64, num int64) (int64, error) {
	return db.HIncrBy(strconv.FormatUint(uid, 10)+"_C", strconv.FormatUint(otherUid, 10), num).Result()
}
func (ur *UnreadRepository) IncrementTotalUnreadBy(uid uint64, num int64) (int64, error) {
	return db.IncrBy(strconv.FormatUint(uid, 10)+"_T", num).Result()
}
func (ur *UnreadRepository) GetUnread(uid, otherUid uint64) (int64, error) {
	return db.HGet(strconv.FormatUint(uid, 10)+"_C", strconv.FormatUint(otherUid, 10)).Int64()
}
func (ur *UnreadRepository) GetTotalUnread(uid uint64) (int64, error) {
	return db.Get(strconv.FormatUint(uid, 10) + "_T").Int64()
}

func (ur *UnreadRepository) DelUnread(uid, otherUid uint64) error {
	return db.HDel(strconv.FormatUint(uid, 10)+"_C", strconv.FormatUint(otherUid, 10)).Err()
}
func (ur *UnreadRepository) DelTotalUnread(uid uint64) error {
	return db.Del(strconv.FormatUint(uid, 10) + "_T").Err()
}
