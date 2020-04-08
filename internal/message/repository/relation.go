package repository

import "inn/internal/message/model"

type IRelationRepository interface {
	Insert(*model.MessageRelation) error
	FindAllByOwnerUidAndOtherUidAndMidIsGreaterThanOrderByMidAsc(ownerUid, otherUid, mid uint64) ([]*model.MessageRelation, error)
	FindAllByOwnerUidAndOtherUidAndMidIsLessThanOrderByMidDescLimit(ownerUid, otherUid, mid uint64, count int64) ([]*model.MessageRelation, error)
}
