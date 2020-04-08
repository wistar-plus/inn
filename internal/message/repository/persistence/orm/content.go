package orm

import (
	"inn/internal/message/model"
	"inn/internal/message/repository"
)

type ContentRepository struct{}

func NewContentRepository() repository.IContentRepository {
	return &ContentRepository{}
}

func (cr *ContentRepository) Insert(content *model.MessageContent) (uint64, error) {
	err := db.Create(content).Error

	return content.Mid, err
}

func (cr *ContentRepository) FindById(mid uint64) (res *model.MessageContent, err error) {
	res = new(model.MessageContent)
	err = db.Where("mid = ?", mid).First(res).Error
	return
}
