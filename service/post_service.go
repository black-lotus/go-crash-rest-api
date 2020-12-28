package service

import "crash-rest-api/rest_model"

type PostService interface {
	Validate(post *rest_model.PostDTO) *rest_model.ServiceError
	Create(post *rest_model.PostDTO) (*rest_model.PostDTO, *rest_model.ServiceError)
	FindAll() ([]rest_model.PostDTO, *rest_model.ServiceError)
}
