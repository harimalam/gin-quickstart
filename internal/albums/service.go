package albums

import "gorm.io/gorm"

// Service defines the methods for business logic.
type Service interface {
	FindAll() ([]Album, error)
	Create(album Album) (Album, error)
	FindById(id uint) (Album, error)
	Update(album Album) (Album, error)
	Delete(id uint) error
}

// service is the concrete implementation of Service.
type service struct {
	repo Repository
}

// NewService is the constructor.
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) FindAll() ([]Album, error) {
	return s.repo.FindAll()
}

func (s *service) Create(album Album) (Album, error) {
	return s.repo.Create(album)
}

func (s *service) FindById(id uint) (Album, error) {
	return s.repo.FindById(id)
}

func (s *service) Update(album Album) (Album, error) {
	exit, err := s.repo.FindById(album.ID)
	if err != nil {
		return Album{}, err
	}
	if exit.ID == 0 {
		return Album{}, gorm.ErrRecordNotFound
	}
	return s.repo.Update(album)
}

func (s *service) Delete(id uint) error {
	return s.repo.Delete(id)
}
