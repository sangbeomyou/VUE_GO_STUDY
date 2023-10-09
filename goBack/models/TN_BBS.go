package models

// TN_BBS is the model entity for the TN_BBS schema.
type TN_BBS struct {
	Idx      uint   `gorm:"primaryKey;column:idx"`
	UserID   string `gorm:"column:user_id"`
	UserName string `gorm:"column:user_name"`
	Title    string `gorm:"column:title"`
	Content  string `gorm:"column:content"`
	RegDate  string `gorm:"column:reg_date"`
	DeleteYn string `gorm:"column:delete_yn"`
}

// TableName sets the insert table name for this struct type
func (b *TN_BBS) TableName() string {
	return "TN_BBS"
}
