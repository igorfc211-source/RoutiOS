package user

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Create(user *User) error {
	return r.db.Create(user).Error
}

func (r *Repository) FindByEmail(email string) (*User, error) {
	var user User

	err := r.db.Where("email = ?", email).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *Repository) Delete(id uuid.UUID) error {
	return r.db.Delete(&User{}, id).Error
}

func (r *Repository) FindAll() ([]User, error) {

	var users []User

	err := r.db.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (r *Repository) FindByID(id uuid.UUID) (*User, error) {

	var user User

	err := r.db.First(&user, id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}