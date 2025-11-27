package albums

import "gorm.io/gorm"

// Repository defines the interface for data access methods.
type Repository interface {
	FindAll() ([]Album, error)
	Create(album Album) (Album, error)
	FindById(id uint) (Album, error)
	Update(album Album) (Album, error)
	Delete(id uint) error
}

// repository is the concrete implementation of the Repository interface.
type repository struct {
	DB *gorm.DB
}

// NewRepository is the constructor for the concrete repository.
func NewRepository(db *gorm.DB) Repository {
	return &repository{DB: db}
}

func (r *repository) FindAll() ([]Album, error) {
	var albums []Album
	if err := r.DB.Find(&albums).Error; err != nil {
		return nil, err
	}
	return albums, nil
}

func (r *repository) Create(album Album) (Album, error) {
	if err := r.DB.Create(&album).Error; err != nil {
		return Album{}, err
	}
	return album, nil
}

func (r *repository) FindById(id uint) (Album, error) {
	var album Album
	if err := r.DB.First(&album, "id = ?", id).Error; err != nil {
		return Album{}, err
	}
	return album, nil
}

func (r *repository) Update(album Album) (Album, error) {
	if err := r.DB.Save(&album).Error; err != nil {
		return Album{}, err
	}
	return album, nil
}

func (r *repository) Delete(id uint) error {
	if err := r.DB.Delete(&Album{}, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
