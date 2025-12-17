package logic

import (
	"github.com/zeromicro/go-zero/core/logx"
	"hanbao-engine/app/hanbao/api/internal/svc"
	"hanbao-engine/app/hanbao/api/internal/types"
	"hanbao-engine/pkg/hanbao"
)

// HanbaoUnlockLogic 词根解锁仪式逻辑
type HanbaoUnlockLogic struct {
	logx.Logger
	ctx    *svc.ServiceContext
}

// NewHanbaoUnlockLogic 创建解锁逻辑
func NewHanbaoUnlockLogic(ctx *svc.ServiceContext) *HanbaoUnlockLogic {
	return &HanbaoUnlockLogic{
		Logger: logx.WithContext(nil),
		ctx:    ctx,
	}
}

// HanbaoUnlock 词根解锁仪式
func (l *HanbaoUnlockLogic) HanbaoUnlock(req *types.UnlockRequest) (resp *types.UnlockResult, err error) {
	l.Info("词根解锁请求: ", req.Words)

	// 调用解锁服务
	result, err := l.ctx.UnlockService.AnalyzeWords(hanbao.UnlockRequest{Words: req.Words})
	if err != nil {
		l.Error("解锁分析失败: ", err)
		return nil, err
	}

	// 转换为API响应格式
	resp = &types.UnlockResult{
		InputWords:     result.InputWords,
		DetectedRoots:  convertCharacterRoots(result.DetectedRoots),
		RootCount:      result.RootCount,
		UnlockableWords: result.UnlockableWords,
		WordBreakdown:  result.WordBreakdown,
		Insights:       result.Insights,
	}

	l.Info("词根解锁成功，发现 ", result.RootCount, " 个字根")
	return resp, nil
}

// convertCharacterRoots 转换字根格式
func convertCharacterRoots(roots []hanbao.CharacterRoot) []types.CharacterRoot {
	result := make([]types.CharacterRoot, len(roots))
	for i, root := range roots {
		result[i] = types.CharacterRoot{
			ID:          root.ID,
			Root:        root.Root,
			Pinyin:      root.Pinyin,
			Difficulty:  root.Difficulty,
			Tier:        root.Tier,
			Description: root.Description,
		}
	}
	return result
}
