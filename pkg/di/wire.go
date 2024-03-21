// +build wireinject

package di


import (
	"service/pkg/api"
	"service/pkg/api/service"
	"service/pkg/config"
	"service/pkg/db"
	"service/pkg/repo"
	

	"github.com/google/wire"
)

func InitializeServer(c *config.Config)(*api.Server,error) {
	wire.Build(db.InitDb,repo.NewUserRepo,service.NewUserServer,api.NewGrpcServe,utils.init)
	return &api.Server{},nil
}