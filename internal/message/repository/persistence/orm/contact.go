package orm

import (
	"inn/internal/message/model"
	"inn/internal/message/repository"
)

type ContactRepository struct{}

func NewContactRepository() repository.IContactRepository {
	return &ContactRepository{}
}

func (cr *ContactRepository) Save(contact *model.MessageContact) error {
	return db.Save(contact).Error
}

func (cr *ContactRepository) FindOne(uid, otherUid uint64) *model.MessageContact {
	contact := &model.MessageContact{}
	if err := db.Where("owner_uid = ? and other_uid = ?", uid, otherUid).First(&contact).Error; err != nil {
		return nil
	}
	return contact
}

func (cr *ContactRepository) FindMessageContactsByOwnerUidOrderByMidDesc(ownerUid uint64) (res []*model.MessageContact, err error) {
	err = db.Where("owner_uid = ?", ownerUid).Order("mid desc").Find(&res).Error
	return
}
