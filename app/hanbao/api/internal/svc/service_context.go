package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"hanbao-engine/pkg/hanbao"
)

// ServiceContext 服务上下文
type ServiceContext struct {
	Config         rest.RestConf
	UnlockService  *hanbao.UnlockCeremonyService
	LevelService   *hanbao.LevelService
	TreasureMapService *hanbao.TreasureMapService
}

// NewServiceContext 创建服务上下文
func NewServiceContext(c rest.RestConf) *ServiceContext {
	return &ServiceContext{
		Config:            c,
		UnlockService:     hanbao.NewUnlockCeremonyService(),
		LevelService:      hanbao.NewLevelService(),
		TreasureMapService: hanbao.NewTreasureMapService(),
	}
}
