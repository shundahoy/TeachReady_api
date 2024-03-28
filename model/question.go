package model

type Question struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Title   string `json:"title"`
	Example string `json:"example"`
	Tags    []*Tag `gorm:"many2many:question_tags; constraint:OnDelete:CASCADE"`
}

// type QuestionResPonse struct {
// 	ID      uint   `json:"id" gorm:"primaryKey"`
// 	Title   string `json:"title"`
// 	Example string `json:"example"`
// }
