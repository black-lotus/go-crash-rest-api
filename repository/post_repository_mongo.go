package repository

import (
	"crash-rest-api/entity"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Post Repository Mongo
type repo struct{}

var (
	database *mgo.Database
)

const (
	collection string = "post"
)

// Create post collection
func NewPostRepositoryMongo(db *mgo.Database) PostRepository {
	database = db
	return &repo{}
}

func (*repo) Save(post *entity.Post) error {
	post.ID = bson.NewObjectId()
	post.CreatedAt = time.Now()
	post.UpdatedAt = time.Now()

	err := database.C(collection).Insert(post)
	return err
}

func (*repo) UpdateByID(id string, post *entity.Post) error {
	post.UpdatedAt = time.Now()
	err := database.C(collection).Update(bson.M{"_id": bson.ObjectIdHex(id)}, post)
	return err
}

func (*repo) DeleteByID(id string) error {
	err := database.C(collection).Remove(bson.M{"_id": bson.ObjectIdHex(id)})
	return err
}

func (*repo) FindByID(id string) (*entity.Post, error) {
	var post entity.Post

	err := database.C(collection).Find(bson.M{"_id": bson.ObjectIdHex(id)}).One(&post)

	if err != nil {
		return nil, err
	}

	return &post, nil
}

func (*repo) FindAll() ([]entity.Post, error) {
	var posts []entity.Post

	err := database.C(collection).Find(bson.M{}).All(posts)

	if err != nil {
		return nil, err
	}

	return posts, nil
}
