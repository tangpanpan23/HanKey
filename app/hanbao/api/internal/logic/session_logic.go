package logic

import (
	"time"

	"github.com/google/uuid"
	"github.com/zeromicro/go-zero/core/logx"
	"hanbao-engine/app/hanbao/api/internal/svc"
	"hanbao-engine/app/hanbao/api/internal/types"
)

// HanbaoStartSessionLogic 会话开始逻辑
type HanbaoStartSessionLogic struct {
	logx.Logger
	ctx    *svc.ServiceContext
}

// NewHanbaoStartSessionLogic 创建会话逻辑
func NewHanbaoStartSessionLogic(ctx *svc.ServiceContext) *HanbaoStartSessionLogic {
	return &HanbaoStartSessionLogic{
		Logger: logx.WithContext(nil),
		ctx:    ctx,
	}
}

// HanbaoStartSession 开始新会话
func (l *HanbaoStartSessionLogic) HanbaoStartSession(req *types.StartSessionRequest) (resp *types.StartSessionResponse, err error) {
	sessionID := uuid.New().String()
	startTime := time.Now()

	l.Info("创建新会话: ", sessionID, " 用户: ", req.UserID)

	resp = &types.StartSessionResponse{
		SessionID: sessionID,
		StartTime: startTime.Format(time.RFC3339),
		Status:    "active",
		Message:   "汉字寻宝之旅开始！请先进行词根解锁仪式。",
	}

	return resp, nil
}
