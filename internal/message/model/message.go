package model

import "time"

type MessageContent struct {
	Mid         uint64    `gorm:"primary_key;auto_increment" json:"mid"`
	SenderId    uint64    `gorm:"not null" json:"senderId"`
	RecipientId uint64    `gorm:"not null" json:"recipientId"`
	Content     string    `gorm:"size:1000;not null" json:"content"`
	MsgType     int       `gorm:"not null" json:"msgType"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type MessageRelation struct {
	OwnerUid  uint64    `gorm:"primary_key;not null" json:"ownerUid"`
	OtherUid  uint64    `gorm:"not null" json:"otherUid"`
	Mid       uint64    `gorm:"primary_key;not null" json:"mid"`
	Type      int       `gorm:"not null" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type MessageContact struct {
	OwnerUid  uint64    `primary_key;gorm:"not null" json:"ownerUid"`
	OtherUid  uint64    `primary_key;gorm:"not null" json:"otherUid"`
	Mid       uint64    `gorm:"not null" json:"mid"`
	Type      int       `gorm:"not null" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}
