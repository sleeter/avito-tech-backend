package http_server

import (
	"avito-tech-backend/internal/core"
	"avito-tech-backend/internal/http-server/handlers"
	"avito-tech-backend/internal/pkg/web"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
)

type App struct {
	Server     web.Server
	Router     *gin.Engine
	Repository *core.Repository
}

func New(repository *core.Repository) *App {
	app := &App{
		Repository: repository,
	}
	app.initRoutes()
	app.Server = web.NewServer(repository.Config.Server, app.Router)
	return app
}

func (app *App) Start(ctx context.Context) error {
	return app.Server.Run(ctx)
}

func (app *App) initRoutes() {
	app.Router = gin.Default()

	app.Router.GET("/user_banner", authMiddleware(), app.mappedHandler(handlers.GetUserBanner))

	admin := app.Router.Group("/")
	admin.Use(authAdminMiddleware())
	{
		app.Router.GET("/banner", app.mappedHandler(handlers.GetBanners))
		app.Router.POST("/banner", app.mappedHandler(handlers.CreateBanner))
		app.Router.PATCH("/banner/:id", app.mappedHandler(handlers.UpdateBanner))
		app.Router.DELETE("/banner/:id", app.mappedHandler(handlers.DeleteBanner))
	}
}

func (app *App) mappedHandler(handler func(*gin.Context, *core.Repository) error) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		if err := handler(ctx, app.Repository); err != nil {
			_ = ctx.AbortWithError(http.StatusInternalServerError, err)
		}
	}
}

func authAdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if token != "admin_token" {
			ctx.AbortWithStatus(http.StatusForbidden)
		}
		ctx.Next()
	}
}
func authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		if token != "user_token" && token != "admin_token" {
			ctx.AbortWithStatus(http.StatusForbidden)
		}
		ctx.Next()
	}
}
