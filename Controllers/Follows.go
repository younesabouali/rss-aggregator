package Controllers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	Middlewares "github.com/younesabouali/rss-aggregator/internal/auth"
	"github.com/younesabouali/rss-aggregator/internal/database"
	"github.com/younesabouali/rss-aggregator/utils"
)

type FollowController struct {
	DB *database.Queries
}

func (fc FollowController) getFollowedFeed(w http.ResponseWriter, r *http.Request, user database.User) {

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
	followedFeed, err := fc.DB.GetFollowedFeed(context.Background(), database.GetFollowedFeedParams{UserID: user.ID, Limit: Limit, Offset: Offset})
	utils.RespondWithArrayJSON(w, 200, followedFeed)
}
func (fc FollowController) followFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type BodyParams struct {
		FeedId uuid.UUID
	}
	bodyParams, err := utils.BodyParser(r, BodyParams{})
	if err != nil {
		utils.RespondWithError(w, 400, "Coundn't parse params")
		return
	}
	defaultParams := utils.GetDefaultParams()
	payload := database.FollowFeedParams{}
	payload.FeedID = bodyParams.FeedId
	payload.UserID = user.ID
	payload.ID = defaultParams.Id
	payload.CreatedAt = defaultParams.CreatedAt
	payload.UpdatedAt = defaultParams.UpdatedAt
	createdFeedFollow, err := fc.DB.FollowFeed(context.Background(), payload)
	if err != nil {
		utils.RespondWithError(w, 400, "bad request")
		return
	}
	utils.RespondWithJSON(w, 200, createdFeedFollow)
}
func (fc FollowController) deleteFollowFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	feedID := utils.UrlParamsParser(r, "feedId")
	feed_uuid, err := uuid.Parse(feedID)
	if err != nil {
		utils.RespondWithError(w, 400, "Invalid id")
		return
	}
	deletedFollow, err := fc.DB.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{UserID: user.ID, FeedID: feed_uuid})
	if err != nil {
		utils.RespondWithError(w, 404, "Invalid id")
		return
	}
	utils.RespondWithJSON(w, 200, deletedFollow)
	return
}
func FollowRouter(DB *database.Queries) *chi.Mux {
	middlewares := Middlewares.Middlewares{DB: DB}
	router := chi.NewRouter()
	followController := FollowController{DB: DB}
	router.Post("/", middlewares.Auth(followController.followFeed))
	router.Delete("/", middlewares.Auth(followController.deleteFollowFeed))
	router.Get("/", middlewares.Auth(followController.getFollowedFeed))
	return router

}
