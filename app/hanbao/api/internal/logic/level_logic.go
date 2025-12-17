package logic

import (
	"errors"
	"strconv"
	"strings"

	"github.com/zeromicro/go-zero/core/logx"
	"hanbao-engine/app/hanbao/api/internal/svc"
	"hanbao-engine/app/hanbao/api/internal/types"
	"hanbao-engine/pkg/hanbao"
)

// HanbaoGetLevelLogic 获取关卡逻辑
type HanbaoGetLevelLogic struct {
	logx.Logger
	ctx    *svc.ServiceContext
}

// NewHanbaoGetLevelLogic 创建关卡获取逻辑
func NewHanbaoGetLevelLogic(ctx *svc.ServiceContext) *HanbaoGetLevelLogic {
	return &HanbaoGetLevelLogic{
		Logger: logx.WithContext(nil),
		ctx:    ctx,
	}
}

// HanbaoGetLevel 获取关卡
func (l *HanbaoGetLevelLogic) HanbaoGetLevel(req *types.LevelRequest) (resp *types.Level, err error) {
	// 解析关卡参数，格式: type_rootId_difficulty
	// 示例: pron_1_1 (音读破译室，字根1，难度1)

	parts := strings.Split(req.LevelId, "_")
	if len(parts) != 3 {
		return nil, errors.New("无效的关卡ID格式")
	}

	levelType := parts[0]
	rootID, _ := strconv.ParseInt(parts[1], 10, 64)
	difficulty, _ := strconv.Atoi(parts[2])

	level, err := l.ctx.LevelService.GenerateLevel(levelType, rootID, difficulty)
	if err != nil {
		l.Error("生成关卡失败: ", err)
		return nil, err
	}

	resp = convertLevel(*level)
	return resp, nil
}

// HanbaoAnswerLevelLogic 关卡答题逻辑
type HanbaoAnswerLevelLogic struct {
	logx.Logger
	ctx    *svc.ServiceContext
}

// NewHanbaoAnswerLevelLogic 创建答题逻辑
func NewHanbaoAnswerLevelLogic(ctx *svc.ServiceContext) *HanbaoAnswerLevelLogic {
	return &HanbaoAnswerLevelLogic{
		Logger: logx.WithContext(nil),
		ctx:    ctx,
	}
}

// HanbaoAnswerLevel 提交答案
func (l *HanbaoAnswerLevelLogic) HanbaoAnswerLevel(req *types.AnswerRequest) (resp *types.AnswerResult, err error) {
	// 注意：这里的req是AnswerRequest，LevelId需要从路径参数获取
	// 暂时使用示例关卡ID
	levelId := "example"
	l.Info("关卡答题: ", levelId, " 问题: ", req.QuestionID)

	result, err := l.ctx.LevelService.ValidateAnswer(levelId, req.QuestionID, req.Answer)
	if err != nil {
		l.Error("答案验证失败: ", err)
		return nil, err
	}

	resp = &types.AnswerResult{
		Correct:     result.Correct,
		Score:       result.Score,
		Explanation: result.Explanation,
		NextHint:    result.NextHint,
	}

	return resp, nil
}

// convertLevel 转换关卡格式
func convertLevel(level hanbao.Level) *types.Level {
	return &types.Level{
		ID:          level.ID,
		Type:        level.Type,
		Title:       level.Title,
		Description: level.Description,
		RootID:      level.RootID,
		Difficulty:  level.Difficulty,
		TimeLimit:   level.TimeLimit,
		Questions:   convertQuestions(level.Questions),
		Reward:      convertReward(level.Reward),
	}
}

// convertQuestions 转换问题格式
func convertQuestions(questions []hanbao.Question) []types.Question {
	result := make([]types.Question, len(questions))
	for i, q := range questions {
		result[i] = types.Question{
			ID:           q.ID,
			Type:         q.Type,
			Content:      q.Content,
			Options:      q.Options,
			CorrectAnswer: q.CorrectAnswer,
			Hint:         q.Hint,
			Explanation:  q.Explanation,
		}
	}
	return result
}

// convertReward 转换奖励格式
func convertReward(reward hanbao.Reward) types.Reward {
	return types.Reward{
		Roots:      reward.Roots,
		Score:      reward.Score,
		Achievement: reward.Achievement,
	}
}
