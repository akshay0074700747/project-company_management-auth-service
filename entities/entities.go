package entities

type Authentication struct {
	UserID   string `gorm:"primaryKey"`
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Authorization struct {
	UserID  string `gorm:"foreignKey:UserID;references:authentications(user_id)"`
	IsAdmin bool   `gorm:"not null"`
}

