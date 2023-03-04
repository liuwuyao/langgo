package api

import (
	"github.com/gin-gonic/gin"
	gs "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"langgo/app/middleware"
	"langgo/bootstrap"
	"langgo/docs"
	_ "langgo/docs"
	"strings"
	"sync"
)

type CoreRouter struct {
	lock            sync.Mutex
	Engine          *gin.Engine
	routerGroupList []func(router *gin.Engine)
}

var coreRouter *CoreRouter

func GetCoreRouter() *CoreRouter {
	return coreRouter
}

// 注意循环依赖问题，先init coreRouter，再在业务层使用coreRouter
func InitCoreRouter() {
	if coreRouter != nil {
		return
	}
	conf := bootstrap.GlobalConfig()
	if strings.ToLower(conf.App.Env) == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	coreRouter = &CoreRouter{
		Engine:          gin.New(),
		routerGroupList: []func(router *gin.Engine){},
	}
	// 跨域 trace-id 日志
	coreRouter.Engine.Use(middleware.GetMiddleWareList()...)

	// 静态资源
	coreRouter.Engine.StaticFile("/assets", "../../static/image/back.png")
	// 注册路由组
	coreRouter.ApplyRouterGroup()

	// swag docs
	docs.SwaggerInfo.BasePath = "/"
	coreRouter.Engine.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))
}

func (core *CoreRouter) AddRouterGroup(f func(router *gin.Engine)) {
	defer func() {
		core.lock.Unlock()
	}()
	core.lock.Lock()
	core.routerGroupList = append(coreRouter.routerGroupList, f)
}

func (core *CoreRouter) ApplyRouterGroup() {
	for _, f := range core.routerGroupList {
		f(core.Engine)
	}
}
