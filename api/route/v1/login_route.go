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

func NewLoginRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	// Creamos el repositorio para el caso de uso
	// Al repositorio le mandamos la BD y la colleccion
	ur := repository.NewUserRepository(db, domain.CollectionUser)

	// Creamos el controlle creando un nuevo caso de uso con el repositorio
	lc := &controller.LoginController{
		LoginUsecase: usecase.NewLoginUsecase(ur, timeout),
		Env:          env,
	}

	// Asignamos el controller al path y este path al group del route
	group.POST("/login", lc.Login)
}
