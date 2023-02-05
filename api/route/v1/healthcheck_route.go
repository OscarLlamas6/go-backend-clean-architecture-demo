package route

import (
	"github.com/gin-gonic/gin"
	hc "github.com/oscarllamas6/go-backend-clean-architecture/api/controller"
)

func NewHealthCheckRouter(group *gin.RouterGroup) {

	// Asignamos el controller al path y este path al group del route
	group.GET("/health-check", hc.HealthCheck)
}
