package model

type User struct {
	ID        uint       `json:"id" gorm:"primaryKey"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"password"`
	Admin     bool       `json:"admin"`
	Questions []Question `gorm:"many2many:user_questions; constraint:OnDelete:CASCADE"`
}

type UserResponse struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Email string `json:"email" gorm:"unique"`
	Admin bool   `json:"admin"`
}
