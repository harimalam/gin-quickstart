package auth

import (
	"errors"
	"gin-quickstart/internal/config"
)

var ErrInvalidCredentials = errors.New("invalid credentials")

type AuthService interface {
	SignUp(req RegisterRequest) (User, error)
	Login(req LoginRequest) (string, error)
}

type authService struct {
	Repo AuthRepository
	Cfg  config.Config
}

func NewService(repo AuthRepository, cfg config.Config) AuthService {
	return &authService{
		Repo: repo,
		Cfg:  cfg,
	}
}

func (s *authService) SignUp(req RegisterRequest) (User, error) {
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return User{}, err
	}
	user := User{
		Username:     req.Username,
		PasswordHash: hashedPassword,
		Role:         req.Role,
	}
	createdUser, err := s.Repo.Create(user)
	if err != nil {
		return User{}, err
	}
	return createdUser, nil
}

func (s *authService) Login(req LoginRequest) (string, error) {
	username := req.Username
	user, err := s.Repo.FindByUsername(username)
	if err != nil {
		return "", err
	}
	if !CheckPasswordHash(req.Password, user.PasswordHash) {
		return "", ErrInvalidCredentials
	}
	token, err := GenerateToken(user, []byte(s.Cfg.JWTSecret))
	if err != nil {
		return "", err
	}
	return token, nil
}
