package main

import (
	"crash-rest-api/config"
	"crash-rest-api/rest"
	"crash-rest-api/router"
	"fmt"
	"net/http"
)

var (
	httpRouter *router.ChiRouter = router.NewChiRouter()
)

func main() {
	db, err := config.GetMongoDB()
	if err != nil {
		panic(err)
	}

	postController := rest.NewPostControllerImpl(db)

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Up and Running")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)

	httpRouter.SERVE(":8080")
}
