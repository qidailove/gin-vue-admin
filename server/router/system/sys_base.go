package system

import (
	"github.com/gin-gonic/gin"
)

type BaseRouter struct{}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	baseRouter := Router.Group("/")
	{
		baseRouter.POST("login", baseApi.Login)
		baseRouter.GET("captchaImage", baseApi.Captcha)
	}
	return baseRouter
}
