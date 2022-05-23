// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameFavorite = "favorite"

// Favorite mapped from table <favorite>
type Favorite struct {
	ID      int32 `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	UserID  int32 `gorm:"column:user_id;not null" json:"user_id"`
	VideoID int32 `gorm:"column:video_id;not null" json:"video_id"`
	Removed int32 `gorm:"column:removed;not null" json:"removed"`
	Deleted int32 `gorm:"column:deleted;not null" json:"deleted"`
}

// TableName Favorite's table name
func (*Favorite) TableName() string {
	return TableNameFavorite
}
