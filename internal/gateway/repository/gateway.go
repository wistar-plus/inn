package repository

import "inn/internal/gateway/model"

type IUserConnRepository interface {
	Add(uid uint64, conn *model.Conn)
	Get(uid uint64) (*model.Conn, bool)
	Del(uid uint64)
}

type IUserTopicRepository interface {
	Add(uid uint64, topic string) (err error)
	Del(uid uint64) (err error)
}
