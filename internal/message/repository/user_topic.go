package repository

type IUserTopicRepository interface {
	Get(uid uint64) string
}
