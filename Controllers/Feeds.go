package Controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	Middlewares "github.com/younesabouali/rss-aggregator/internal/auth"
	"github.com/younesabouali/rss-aggregator/internal/database"
	"github.com/younesabouali/rss-aggregator/utils"
)

type FeedController struct {
	DB *database.Queries
}

func (fc FeedController) createFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type FeedParams struct {
		Name string
		Url  string
	}
	feedParams, err := utils.BodyParser(r, FeedParams{})
	if err != nil {
		utils.RespondWithError(w, 400, err.Error())
		return
	}
	dbParams := database.CreateFeedParams{}
	defaultParams := utils.GetDefaultParams()
	dbParams.ID = defaultParams.Id
	dbParams.Createdat = defaultParams.CreatedAt
	dbParams.Updatedat = defaultParams.UpdatedAt
	dbParams.Name = feedParams.Name
	dbParams.Url = feedParams.Url
	dbParams.UserID = user.ID

	createdFeed, err := fc.DB.CreateFeed(context.Background(), dbParams)
	if err != nil {
		fmt.Printf(err.Error())
		utils.RespondWithError(w, 400, "Couldn't create Feed")
		return
	}
	// fmt.Println(feedParams.Link, feedParams.Name)
	utils.RespondWithJSON(w, 200, createdFeed)
}
func (fc FeedController) getAllFeeds(w http.ResponseWriter, r *http.Request) {
	type URLParams struct {
		Limit  int
		Offset int
	}
	dbParams := database.GetFeedParams{}
	e, err := utils.ParseInt32(utils.UrlParamsParser(r, "Limit"))
	if err != nil {
		utils.RespondWithError(w, 400, "Couldn't Search Feed")
		return
	}
	dbParams.Limit = e
	e, err = utils.ParseInt32(utils.UrlParamsParser(r, "Offset"))
	if err != nil {
		utils.RespondWithError(w, 400, "Unable to parse params")
		return
	}
	dbParams.Offset = e
	feed, err := fc.DB.GetFeed(context.Background(), dbParams)
	if err != nil {
		fmt.Println(err.Error(), dbParams)
		utils.RespondWithError(w, 404, "Unable to perform search")
		return
	}

	if feed != nil {
		utils.RespondWithJSON(w, 200, feed)
		return
	}
	var emptyFeed [0]int
	utils.RespondWithJSON(w, 200, emptyFeed)
}

func FeedRouter(DB *database.Queries) *chi.Mux {
	middlewares := Middlewares.Middlewares{DB: DB}
	router := chi.NewRouter()
	feedController := FeedController{DB: DB}
	router.Post("/", middlewares.Auth(feedController.createFeed))
	router.Get("/", feedController.getAllFeeds)
	return router

}
