package orm

import (
	"inn/internal/message/model"
	"inn/internal/message/repository"
)

type RelationRepository struct{}

func NewRelationRepository() repository.IRelationRepository {
	return &RelationRepository{}
}

func (rr *RelationRepository) Insert(relation *model.MessageRelation) error {
	return db.Create(relation).Error
}

func (rr *RelationRepository) FindAllByOwnerUidAndOtherUidAndMidIsGreaterThanOrderByMidAsc(ownerUid, otherUid, mid uint64) (res []*model.MessageRelation, err error) {
	err = db.Where("owner_uid = ? AND other_uid = ? AND mid > ?", ownerUid, otherUid, mid).Order("mid asc").Find(&res).Error
	return
}

func (rr *RelationRepository) FindAllByOwnerUidAndOtherUidAndMidIsLessThanOrderByMidDescLimit(ownerUid, otherUid, mid uint64, count int64) (res []*model.MessageRelation, err error) {
	if mid >= 0 {
		err = db.Where("owner_uid = ? AND other_uid = ? AND mid < ?", ownerUid, otherUid, mid).Order("mid desc").Limit(count).Find(&res).Error
	} else {
		err = db.Where("owner_uid = ? AND other_uid", ownerUid, otherUid).Order("mid desc").Find(&res).Limit(count).Error
	}
	return
}
