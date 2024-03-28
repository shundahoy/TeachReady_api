package model

type Answer struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	QuestionId uint     `json:"question_id" gorm:"not null"`
	Question   Question `json:"question" gorm:"foreignKey:QuestionId; constraint:OnDelete:CASCADE"`
	UserId     uint     `json:"user_id" gorm:"not null"`
	User       User     `json:"user" gorm:"foreignKey:UserId; constraint:OnDelete:CASCADE"`
	Answer     string   `json:"answer"`
}
