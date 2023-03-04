package app

import (
	"context"
	"go.uber.org/zap"
	"langgo/api"
	"langgo/bootstrap"
	"langgo/config"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// App 应用结构体
type App struct {
	conf    *config.Configuration
	logger  *zap.Logger
	httpSrv *http.Server
}

// NewHttpServer 创建http app engine
func NewHttpServer() *http.Server {
	return &http.Server{
		Addr:    ":" + bootstrap.GlobalConfig().App.Port,
		Handler: api.GetCoreRouter().Engine,
	}
}

// NewApp 创建新应用
func NewApp(
	logger *zap.Logger,
	httpSrv *http.Server,
) *App {
	return &App{
		conf:    bootstrap.GlobalConfig(),
		logger:  logger,
		httpSrv: httpSrv,
	}
}

// RunServer 启动服务
func (a *App) RunServer() {
	// 启动应用
	a.logger.Info("start app ...")
	if err := a.run(); err != nil {
		panic(err)
	}

	// 等待中断信号以优雅地关闭应用
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Printf("shutdown app ...")

	// 设置 5 秒的超时时间
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 关闭应用
	if err := a.stop(ctx); err != nil {
		panic(err)
	}
}

// run 启动服务
func (a *App) run() error {
	// 启动 http app
	go func() {
		a.logger.Info("http app started")
		if err := a.httpSrv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()
	return nil
}

// stop 停止服务
func (a *App) stop(ctx context.Context) error {
	// 关闭 http app
	a.logger.Info("http app has been stop")
	if err := a.httpSrv.Shutdown(ctx); err != nil {
		return err
	}
	return nil
}
