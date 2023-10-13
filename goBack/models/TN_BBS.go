package models

// TN_BBS is the model entity for the TN_BBS schema.
type TN_BBS struct {
	Idx      uint   `gorm:"primaryKey;column:idx;autoIncrement"`
	UserID   string `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
	RegDate  string `gorm:"column:reg_date"`
	UpdDate  string `gorm:"column:upd_date"`
	DeleteYn string `gorm:"column:delete_yn"`
}

type TN_BBS_API struct {
	Idx      string `json:"idx"`
	UserID   string `json:"user_id,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	RegDate  string `json:"reg_date,omitempty"` // omitempty: 필드가 비어있을 경우 JSON에서 생략
	UpdDate  string `json:"upd_date,omitempty"`
	DeleteYn string `json:"delete_yn,omitempty"`
}

// TableName sets the insert table name for this struct type
func (b *TN_BBS) TableName() string {
	return "TN_BBS"
}
