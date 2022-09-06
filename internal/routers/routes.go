package routers

import (
	"github.com/gin-gonic/gin"
)

//type Router gin.Engine
//
//func (*Router) AddSubRoutes() {
//
//}

func SetetupRouter() *gin.Engine {
	router := gin.Default()
	apiV1Group := router.Group("api/v1")
	// 注册登陆路由到 api v1 group
	loadAuthRouter(apiV1Group)
	loadHostTRouter(apiV1Group)
	apiV2Group := router.Group("api/v2")
	loadHostTRouter(apiV2Group)
	return router
}
