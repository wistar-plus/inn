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
	OwnerUid  uint64    `gorm:"primary_key;auto_increment:false;not null" json:"ownerUid"`
	OtherUid  uint64    `gorm:"not null" json:"otherUid"`
	Mid       uint64    `gorm:"primary_key;auto_increment:false;not null" json:"mid"`
	Type      int       `gorm:"not null" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type MessageContact struct {
	OwnerUid  uint64    `gorm:"primary_key;auto_increment:false;not null" json:"ownerUid"`
	OtherUid  uint64    `gorm:"primary_key;auto_increment:false;not null" json:"otherUid"`
	Mid       uint64    `gorm:"not null" json:"mid"`
	Type      int       `gorm:"not null" json:"type"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
}

type MessageContactVO struct {
	OwnerUid     uint64
	OwnerAvatar  string
	OwnerName    string
	TotalUnread  int64
	ContactInfos []*ContactInfo
}

type ContactInfo struct {
	OtherUid    uint64
	OtherName   string
	OtherAvatar string
	Mid         uint64
	Type        int
	Content     string
	ConvUnread  int64
	CreateTime  int64
}

type MessageVO struct {
	Mid            uint64
	Content        string
	OwnerUid       uint64
	Type           int
	OtherUid       uint64
	CreateTime     int64
	OwnerUidAvatar string
	OtherUidAvatar string
	OwnerName      string
	OtherName      string
}
