package bootstrap

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"langgo/config"
	"langgo/utils"
	"path/filepath"
	"sync"
)

var (
	configPath          string
	rootPath            = utils.RootPath()
	lgConfig            = new(LangGoConfig)
	defaultConfFilePath = ""
)

// LangGoConfig 自定义Log
type LangGoConfig struct {
	Conf *config.Configuration
	Once *sync.Once
}

// newLangGoConfig .
func newLangGoConfig() *LangGoConfig {
	return &LangGoConfig{
		Conf: &config.Configuration{},
		Once: &sync.Once{},
	}
}

func GlobalConfig() *config.Configuration {
	// main 函数已经调用过NewConfig方法
	return lgConfig.Conf
}

// NewConfig 初始化配置对象
func NewConfig(confFile string) {
	if lgConfig.Conf != nil {
		return
	} else {
		lgConfig = newLangGoConfig()
		if confFile == "" {
			lgConfig.initLangGoConfig(defaultConfFilePath)
		} else {
			lgConfig.initLangGoConfig(confFile)
		}
		return
	}
}

// InitLangGoConfig 初始化日志
func (lg *LangGoConfig) initLangGoConfig(confFile string) {
	lg.Once.Do(
		func() {
			initConfig(lg.Conf, confFile)
		},
	)
}

func initConfig(conf *config.Configuration, confFile string) {
	pflag.StringVarP(&configPath, "conf", "", filepath.Join(rootPath, confFile),
		"config path, eg: --conf config.yaml")
	if !filepath.IsAbs(configPath) {
		configPath = filepath.Join(rootPath, configPath)
	}

	//lgLogger.Logger.Info("load config:" + configPath)
	fmt.Println("Load Config: " + configPath)

	v := viper.New()
	v.SetConfigFile(configPath)
	v.SetConfigType("yaml")
	if err := v.ReadInConfig(); err != nil {
		//lgLogger.Logger.Error("read config failed: ", zap.String("err", err.Error()))
		fmt.Println("read config failed: ", zap.String("err", err.Error()))
		panic(err)
	}

	if err := v.Unmarshal(&conf); err != nil {
		//lgLogger.Logger.Error("config parse failed: ", zap.String("err", err.Error()))
		fmt.Println("config parse failed: ", zap.String("err", err.Error()))
	}

	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		//lgLogger.Logger.Info("", zap.String("config file changed:", in.Name))
		fmt.Println("", zap.String("config file changed:", in.Name))
		defer func() {
			if err := recover(); err != nil {
				//lgLogger.Logger.Error("config file changed err:", zap.Any("err", err))
				fmt.Println("config file changed err:", zap.Any("err", err))
			}
		}()
		if err := v.Unmarshal(&conf); err != nil {
			//lgLogger.Logger.Error("config parse failed: ", zap.String("err", err.Error()))
			fmt.Println("config parse failed: ", zap.String("err", err.Error()))
		}
	})
	lgConfig.Conf = conf
}
