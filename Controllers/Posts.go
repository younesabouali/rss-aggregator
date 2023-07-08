package Controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/younesabouali/rss-aggregator/internal/database"
	"github.com/younesabouali/rss-aggregator/utils"
)

type PostsController struct {
	DB *database.Queries
}

func (pc *PostsController) GetFollowedPosts(w http.ResponseWriter, r *http.Request, user database.User) {

	Limit, err := utils.ParseInt32(utils.UrlParamsParser(r, "Limit"))
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to parse params")
		return
	}

	Offset, err := utils.ParseInt32(utils.UrlParamsParser(r, "Offset"))
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to parse params")
		return
	}
	posts, err := pc.DB.GetFollowedPosts(context.Background(), database.GetFollowedPostsParams{Limit: Limit, Offset: Offset, UserID: user.ID})
	if err != nil {
		utils.RespondWithError(w, 404, "Posts not found")
		return
	}
	utils.RespondWithArrayJSON(w, 200, posts)
}
func PostRouter(DB *database.Queries) *chi.Mux {
	userController := PostsController{DB}
	middlewares, router := InitializeDependencies(DB)
	router.Get("/", middlewares.Auth(userController.GetFollowedPosts))
	return router

}
