package repository

import "inn/internal/message/model"

type IContentRepository interface {
	Insert(*model.MessageContent) (uint64, error)
	FindById(uint64) (*model.MessageContent, error)
}
