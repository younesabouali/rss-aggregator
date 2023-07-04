package Controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	Middlewares "github.com/younesabouali/rss-aggregator/internal/auth"
	"github.com/younesabouali/rss-aggregator/internal/database"
	"github.com/younesabouali/rss-aggregator/utils"
	jsonformatter "github.com/younesabouali/rss-aggregator/utils"
)

type UserController struct {
	DB *database.Queries
}

func (c UserController) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	type userParams struct {
		Name string
	}
	result, err := jsonformatter.BodyParser(r, userParams{})
	if err != nil {
		jsonformatter.RespondWithError(w, 400, "couldn't parse user")
		return
	}
	createdUser, err := c.DB.CreateUser(context.Background(), database.CreateUserParams{ID: uuid.New(), Name: result.Name, Updatedat: time.Now(), Createdat: time.Now()})
	if err != nil {
		utils.RespondWithError(w, 400, "Couldn't create user")
	}
	utils.RespondWithJSON(w, 200, createdUser)
}
func (c UserController) GetByApiKey(w http.ResponseWriter, r *http.Request, user database.User) {
	jsonformatter.RespondWithJSON(w, 200, user)
}
func UserRouter(DB *database.Queries) *chi.Mux {
	middlewares := Middlewares.Middlewares{DB: DB}
	router := chi.NewRouter()
	userController := UserController{DB: DB}
	router.Post("/", userController.CreateUserHandler)
	router.Get("/", middlewares.Auth(userController.GetByApiKey))
	return router

}
