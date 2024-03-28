package model

type Tag struct {
	ID        uint        `json:"id" gorm:"primaryKey"`
	Name      string      `json:"name" gorm:"unique"`
	Questions []*Question `gorm:"many2many:question_tags; constraint:OnDelete:CASCADE"`
}

// type TagResPonse struct {
// 	ID   uint   `json:"id" gorm:"primaryKey"`
// 	Name string `json:"name" gorm:"unique"`
// }
