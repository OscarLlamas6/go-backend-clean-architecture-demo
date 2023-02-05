package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oscarllamas6/go-backend-clean-architecture/api/middleware"
	"github.com/oscarllamas6/go-backend-clean-architecture/bootstrap"
	"github.com/oscarllamas6/go-backend-clean-architecture/mongo"
)

func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, routerV1 *gin.RouterGroup) {
	publicRouterV1 := routerV1.Group("")
	// All Public APIs
	NewSignupRouter(env, timeout, db, publicRouterV1)
	NewHealthCheckRouter(publicRouterV1)
	// El router llama a uno o varios routes
	NewLoginRouter(env, timeout, db, publicRouterV1)
	NewRefreshTokenRouter(env, timeout, db, publicRouterV1)

	protectedRouterV1 := routerV1.Group("")
	// Middleware to verify AccessToken
	protectedRouterV1.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	NewProfileRouter(env, timeout, db, protectedRouterV1)
	NewTaskRouter(env, timeout, db, protectedRouterV1)
}
