package user

import "github.com/google/uuid"
import "gorm.io/gorm"




type User struct{

	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"  gorm:"unique" ` 
	Password string `json:"-"`
}

func (u *User) BeforeCreate(tx *gorm.DB) error {

	u.ID = uuid.New()

	return nil
}