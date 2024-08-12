package routes

import (
	"github.com/RivGames/my-knowledge-base/internal/controllers"
	"github.com/RivGames/my-knowledge-base/internal/middleware"
	"github.com/RivGames/my-knowledge-base/pkg/app"
	"github.com/RivGames/my-knowledge-base/pkg/validation"
	"github.com/labstack/echo/v4"
)

type Router struct {
	*echo.Echo
	Middleware *middleware.Middleware
	App        *app.App
}

func NewRouter(app *app.App) *Router {
	echo := echo.New()
	echo.Validator = validation.NewCustomValidator()

	middleware := middleware.New(echo, app)

	return &Router{echo, middleware, app}
}

func (r *Router) ListenAndServe() {
	r.setupRoutes()

	r.listenAndServe()
}

// Add Swagger
// Add user relation to question
// Add "data" wrapper to every response, except responses that contain only one message

func (r *Router) setupRoutes() {
	apiV1 := r.Echo.Group("/api/v1")

	apiV1.GET("/up", controllers.Up)

	// Auth
	apiV1.POST("/register", controllers.Register)
	apiV1.POST("/login", controllers.Login)

	// Questions
	apiV1.GET("/questions", controllers.GetQuestions)
	apiV1.POST("/questions", controllers.CreateQuestion, middleware.WithAuthentication)
	apiV1.GET("/questions/:id", controllers.GetQuestion)
	apiV1.PUT("/questions/:id", controllers.UpdateQuestion, middleware.WithAuthentication)
	apiV1.DELETE("/questions/:id", controllers.DeleteQuestion, middleware.WithAuthentication)

	// Answers
	apiV1.GET("/answers", controllers.GetAnswers)
	apiV1.POST("/answers", controllers.CreateAnswer, middleware.WithAuthentication)
	apiV1.GET("/answers/:id", controllers.GetAnswer)
	apiV1.GET("/question/:id/answers", controllers.GetQuestionAnswers)
	apiV1.PUT("/answers/:id", controllers.UpdateAnswer, middleware.WithAuthentication)
	apiV1.DELETE("/answers/:id", controllers.DeleteAnswer, middleware.WithAuthentication)

}

func (r *Router) listenAndServe() {
	r.Echo.Logger.Fatal(r.Echo.Start(":80"))
}
