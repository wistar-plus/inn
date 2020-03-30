package syncmap

import (
	"inn/internal/gateway/model"
	"inn/internal/gateway/repository"
)

type UserConnRepository struct{}

func NewUserConnRepository() repository.IUserConnRepository {
	return &UserConnRepository{}
}

func (ucr *UserConnRepository) Add(uid uint64, conn *model.Conn) {
	m.Store(uid, conn)
}

func (r *UserConnRepository) Get(uid uint64) (*model.Conn, bool) {
	val, ok := m.Load(uid)
	if !ok {
		return nil, ok
	}
	conn, ok := val.(*model.Conn)
	if !ok {
		return nil, ok
	}
	return conn, ok
}

func (ucr *UserConnRepository) Del(uid uint64) {
	m.Delete(uid)
}
