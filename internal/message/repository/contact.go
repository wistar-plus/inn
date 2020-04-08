package repository

import "inn/internal/message/model"

type IContactRepository interface {
	Save(*model.MessageContact) error
	FindOne(uid, otherUid uint64) *model.MessageContact
	FindMessageContactsByOwnerUidOrderByMidDesc(ownerUid uint64) ([]*model.MessageContact, error)
}
