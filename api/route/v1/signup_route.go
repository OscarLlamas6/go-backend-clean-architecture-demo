package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oscarllamas6/go-backend-clean-architecture/api/controller"
	"github.com/oscarllamas6/go-backend-clean-architecture/bootstrap"
	"github.com/oscarllamas6/go-backend-clean-architecture/domain"
	"github.com/oscarllamas6/go-backend-clean-architecture/mongo"
	"github.com/oscarllamas6/go-backend-clean-architecture/repository"
	"github.com/oscarllamas6/go-backend-clean-architecture/usecase"
)

func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	sc := controller.SignupController{
		SignupUsecase: usecase.NewSignupUsecase(ur, timeout),
		Env:           env,
	}
	group.POST("/signup", sc.Signup)
}
