package auth

import "gorm.io/gorm"

type AuthRepository interface {
	Create(user User) (User, error)
	FindByUsername(username string) (User, error)
}

type authRepository struct {
	DB *gorm.DB
}

func NewRepository(db *gorm.DB) AuthRepository {
	return &authRepository{DB: db}
}

func (r *authRepository) Create(user User) (User, error) {
	if err := r.DB.Create(&user).Error; err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *authRepository) FindByUsername(username string) (User, error) {
	var user User
	if err := r.DB.First(&user, "username = ?", username).Error; err != nil {
		return User{}, err
	}
	return user, nil
}
