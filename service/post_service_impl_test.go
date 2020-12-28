package service

import (
	"crash-rest-api/entity"
	"crash-rest-api/rest_model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockPostRepository struct {
	mock.Mock
}

func (mock *MockPostRepository) Save(*entity.Post) error {
	args := mock.Called()

	return args.Error(1)
}

func (mock *MockPostRepository) UpdateByID(string, *entity.Post) error {
	args := mock.Called()

	return args.Error(1)
}

func (mock *MockPostRepository) DeleteByID(string) error {
	args := mock.Called()

	return args.Error(1)
}

func (mock *MockPostRepository) FindByID(string) (*entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (mock *MockPostRepository) FindAll() ([]entity.Post, error) {
	args := mock.Called()
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestFindAllPosts_whenRepositoryExist_thenShouldSuccess(t *testing.T) {
	mockPostRepository := new(MockPostRepository)

	// GIVEN
	post := entity.Post{
		Title: "title",
		Text:  "text",
	}
	mockPostRepository.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostServiceImpl(mockPostRepository)

	// WHEN
	result, _ := testService.FindAll()

	// THEN
	postExpected := rest_model.PostDTO{
		Title: "title",
		Text:  "text",
	}

	mockPostRepository.AssertExpectations(t)
	assert.NotNil(t, result)
	assert.Equal(t, postExpected, result[0])
}

func TestValidate_whenPostIsEmpty_thenShouldError(t *testing.T) {
	testService := NewPostServiceImpl(nil)
	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, "The post is empty", err.Message)
}

func TestValidate_whenPostTitleIsEmpty_thenShouldError(t *testing.T) {
	testService := NewPostServiceImpl(nil)

	post := rest_model.PostDTO{
		ID:    "1",
		Title: "",
		Text:  "",
	}
	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post title is empty", err.Message)
}

func TestValidate_whenPostTextIsEmpty_thenShouldError(t *testing.T) {
	testService := NewPostServiceImpl(nil)

	post := rest_model.PostDTO{
		ID:    "1",
		Title: "title",
		Text:  "",
	}
	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, "The post text is empty", err.Message)
}
