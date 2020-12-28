package repository

import "crash-rest-api/entity"

// PostRepository interface
type PostRepository interface {
	Save(post *entity.Post) error
	UpdateByID(id string, post *entity.Post) error
	DeleteByID(id string) error
	FindByID(id string) (*entity.Post, error)
	FindAll() ([]entity.Post, error)
}
