package plugins

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"go.uber.org/zap"
	"langgo/bootstrap"
	"langgo/config"
	"sync"
)

var lgMinio = new(LangGoMinio)

// LangGoMinio .
type LangGoMinio struct {
	Once        *sync.Once
	MinioClient *minio.Client
}

// NewMinio .
func (lg *LangGoMinio) NewMinio() *minio.Client {
	if lgMinio.MinioClient != nil {
		return lgMinio.MinioClient
	} else {
		return lg.New().(*minio.Client)
	}
}

func newLangGoMinio() *LangGoMinio {
	return &LangGoMinio{
		MinioClient: &minio.Client{},
		Once:        &sync.Once{},
	}
}

// Name .
func (lg *LangGoMinio) Name() string {
	return "Minio"
}

// New .
func (lg *LangGoMinio) New() interface{} {
	lgMinio = newLangGoMinio()
	lgMinio.initMinio(bootstrap.GlobalConfig())
	return lg.MinioClient
}

// Health .
func (lg *LangGoMinio) Health() {
	_, err := lgMinio.MinioClient.ListBuckets(context.Background())
	if err != nil {
		bootstrap.NewLogger().Logger.Error("Minio connect failed, err:", zap.Any("err", err))
		panic("failed to connect minio")
	}
}

// Close .
func (lg *LangGoMinio) Close() {}

func init() {
	p := &LangGoMinio{}
	RegisteredPlugin(p)
}

func (lg *LangGoMinio) initMinio(conf *config.Configuration) {
	lg.Once.Do(func() {
		if !conf.IsMinioEnable() {
			return
		}
		endpoint := conf.Minio.EndPoint
		accessKeyID := conf.Minio.AccessKeyID
		secretAccessKey := conf.Minio.SecretAccessKey
		useSSL := conf.Minio.UseSSL
		client, err := minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
			Secure: useSSL,
		})

		if err != nil {
			bootstrap.NewLogger().Logger.Error("minio连接错误: ", zap.Any("err", err))
			panic(err)
		} else {
			lgMinio.MinioClient = client
		}
	})
}
