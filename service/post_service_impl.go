package service

import (
	"crash-rest-api/entity"
	"crash-rest-api/repository"
	"crash-rest-api/rest_model"

	"gopkg.in/mgo.v2/bson"
)

type PostServiceImpl struct{}

var (
	postRepository repository.PostRepository
)

func NewPostServiceImpl(repository repository.PostRepository) PostService {
	postRepository = repository
	return &PostServiceImpl{}
}

func (*PostServiceImpl) Validate(post *rest_model.PostDTO) *rest_model.ServiceError {
	if post == nil {
		return &rest_model.ServiceError{
			Message: "The post is empty",
		}
	}

	if post.Title == "" {
		return &rest_model.ServiceError{
			Message: "The post title is empty",
		}
	}

	if post.Text == "" {
		return &rest_model.ServiceError{
			Message: "The post text is empty",
		}
	}

	return nil
}

func (service *PostServiceImpl) Create(postDTO *rest_model.PostDTO) (*rest_model.PostDTO, *rest_model.ServiceError) {
	validateErr := service.Validate(postDTO)
	if validateErr != nil {
		return nil, validateErr
	}

	post := castingToPost(postDTO)
	err := postRepository.Save(post)
	if err != nil {
		serviceError := &rest_model.ServiceError{
			Message: err.Error(),
		}
		return nil, serviceError
	}

	return castingToPostDTO(post), nil
}

func (*PostServiceImpl) FindAll() ([]rest_model.PostDTO, *rest_model.ServiceError) {
	posts, err := postRepository.FindAll()
	if err != nil {
		serviceError := &rest_model.ServiceError{
			Message: err.Error(),
		}
		return nil, serviceError
	}

	postsDTO := make([]rest_model.PostDTO, 0, len(posts))
	for _, post := range posts {
		postsDTO = append(postsDTO, *castingToPostDTO(&post))
	}

	return postsDTO, nil
}

func castingToPostDTO(post *entity.Post) *rest_model.PostDTO {
	return &rest_model.PostDTO{
		ID:        post.ID.Hex(),
		Title:     post.Title,
		Text:      post.Text,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func castingToPost(postDTO *rest_model.PostDTO) *entity.Post {
	var id bson.ObjectId
	if postDTO.ID != "" {
		id = bson.ObjectIdHex(postDTO.ID)
	}

	return &entity.Post{
		ID:        id,
		Title:     postDTO.Title,
		Text:      postDTO.Text,
		CreatedAt: postDTO.CreatedAt,
		UpdatedAt: postDTO.UpdatedAt,
	}
}
