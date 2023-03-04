package middleware

import (
	"github.com/gin-gonic/gin"
	"langgo/bootstrap"
	"sync"
)

//支持自定义中间件,可以通过Remove和Add调整顺序
type GlobalMiddleware struct {
	lock           sync.Mutex
	middlewareList []gin.HandlerFunc
}

var globalMiddleware *GlobalMiddleware

func GetMiddleWareList() []gin.HandlerFunc {
	return globalMiddleware.middlewareList
}

func Init() {
	globalMiddleware = &GlobalMiddleware{
		lock:           sync.Mutex{},
		middlewareList: []gin.HandlerFunc{},
	}
	lgLogger := bootstrap.NewLogger()

	globalMiddleware.AddAfterLast(NewCors().Handler())
	globalMiddleware.AddAfterLast(NewTrace(lgLogger).Handler())
	globalMiddleware.AddAfterLast(NewRequestLog(lgLogger).Handler())
}

func (g *GlobalMiddleware) AddAfterLast(handlerFunc gin.HandlerFunc) {
	defer func() {
		g.lock.Unlock()
	}()
	g.lock.Lock()
	g.middlewareList = append(g.middlewareList, handlerFunc)
}

func (g *GlobalMiddleware) RemoveLastMiddleware() gin.HandlerFunc {
	defer func() {
		g.lock.Unlock()
	}()
	g.lock.Lock()
	if len(g.middlewareList) == 0 {
		return nil
	}
	length := len(g.middlewareList)
	res := g.middlewareList[length-1]
	g.middlewareList = g.middlewareList[:length-1]
	return res
}

func (g *GlobalMiddleware) AddBeforeFirst(handlerFunc gin.HandlerFunc) {
	defer func() {
		g.lock.Unlock()
	}()
	g.lock.Lock()
	g.middlewareList = append([]gin.HandlerFunc{handlerFunc}, g.middlewareList...)
}
