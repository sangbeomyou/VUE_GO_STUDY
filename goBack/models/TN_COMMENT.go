package models

// TN_BBS is the model entity for the TN_BBS schema.
type TN_COMMENT struct {
	Idx       uint         `gorm:"primaryKey;column:idx;autoIncrement"`
	ParentIdx *uint        `gorm:"column:parent_idx"`
	BbsId     string       `gorm:"column:bbs_id"`
	UserID    string       `gorm:"column:user_id"`
	UserName  string       `gorm:"column:user_name"`
	Content   string       `gorm:"column:content"`
	RegDate   string       `gorm:"column:reg_date"`
	UpdDate   string       `gorm:"column:upd_date"`
	DeleteYn  string       `gorm:"column:delete_yn"`
	Replies   []TN_COMMENT `gorm:"-"` // 대댓글 목록을 위한 필드, DB에 저장되지 않음
}

type TN_COMMENT_API struct {
	Idx       string `json:"idx"`
	ParentIdx string `json:"parent_idx,omitempty"`
	BbsId     string `json:"bbs_id,omitempty"`
	UserID    string `json:"user_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	Content   string `json:"content"`
	RegDate   string `json:"reg_date,omitempty"` // omitempty: 필드가 비어있을 경우 JSON에서 생략
	UpdDate   string `json:"upd_date,omitempty"`
	DeleteYn  string `json:"delete_yn,omitempty"`
}

// TableName sets the insert table name for this struct type
func (b *TN_COMMENT) TableName() string {
	return "TN_COMMENT"
}
