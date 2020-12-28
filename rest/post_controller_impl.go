package rest

import (
	"crash-rest-api/repository"
	"crash-rest-api/rest_model"
	"crash-rest-api/service"
	"encoding/json"
	"net/http"

	"gopkg.in/mgo.v2"
)

type PostControllerImpl struct {
	postService service.PostService
}

func NewPostControllerImpl(database *mgo.Database) *PostControllerImpl {
	postRepository := repository.NewPostRepositoryMongo(database)
	postService := service.NewPostServiceImpl(postRepository)
	return &PostControllerImpl{
		postService: postService,
	}
}

func (controller *PostControllerImpl) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, serviceErr := controller.postService.FindAll()
	if serviceErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (controller *PostControllerImpl) AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post rest_model.PostDTO
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	createdPost, serviceErr := controller.postService.Create(&post)
	if serviceErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(serviceErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdPost)
}
