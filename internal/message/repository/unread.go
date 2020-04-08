package repository

type IUnreadRepository interface {
	IncrementUnreadBy(uid, otherUid uint64, num int64) (int64, error)
	IncrementTotalUnreadBy(uint64, int64) (int64, error)
	GetUnread(uint64, uint64) (int64, error)
	GetTotalUnread(uint64) (int64, error)
	DelUnread(uint64, uint64) error
	DelTotalUnread(uint64) error
}
