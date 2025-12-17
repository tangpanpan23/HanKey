package logic

import (
	"github.com/zeromicro/go-zero/core/logx"
	"hanbao-engine/app/hanbao/api/internal/svc"
	"hanbao-engine/app/hanbao/api/internal/types"
	"hanbao-engine/pkg/hanbao"
)

// HanbaoGetTreasureMapLogic 获取藏宝图逻辑
type HanbaoGetTreasureMapLogic struct {
	logx.Logger
	ctx    *svc.ServiceContext
}

// NewHanbaoGetTreasureMapLogic 创建藏宝图逻辑
func NewHanbaoGetTreasureMapLogic(ctx *svc.ServiceContext) *HanbaoGetTreasureMapLogic {
	return &HanbaoGetTreasureMapLogic{
		Logger: logx.WithContext(nil),
		ctx:    ctx,
	}
}

// HanbaoGetTreasureMap 获取藏宝图
func (l *HanbaoGetTreasureMapLogic) HanbaoGetTreasureMap(req *types.TreasureMapRequest) (resp *types.TreasureMap, err error) {
	l.Info("获取藏宝图: ", req.SessionID)

	// 示例：假设用户已解锁前5个字根
	unlockedRoots := []int64{1, 2, 3, 4, 5}

	treasureMap, err := l.ctx.TreasureMapService.GenerateTreasureMap(req.SessionID, unlockedRoots)
	if err != nil {
		l.Error("生成藏宝图失败: ", err)
		return nil, err
	}

	resp = convertTreasureMap(treasureMap)
	return resp, nil
}

// HanbaoGetRecommendationsLogic 获取推荐逻辑
type HanbaoGetRecommendationsLogic struct {
	logx.Logger
	ctx    *svc.ServiceContext
}

// NewHanbaoGetRecommendationsLogic 创建推荐逻辑
func NewHanbaoGetRecommendationsLogic(ctx *svc.ServiceContext) *HanbaoGetRecommendationsLogic {
	return &HanbaoGetRecommendationsLogic{
		Logger: logx.WithContext(nil),
		ctx:    ctx,
	}
}

// HanbaoGetRecommendations 获取推荐
func (l *HanbaoGetRecommendationsLogic) HanbaoGetRecommendations(req *types.RecommendationsRequest) (resp *types.RecommendationsResponse, err error) {
	l.Info("获取推荐: ", req.SessionID)

	// 示例：当前已解锁的字根
	currentRoots := []int64{1, 2, 3, 4, 5}

	recommendations := l.ctx.TreasureMapService.GetNextRecommendations(currentRoots)

	resp = &types.RecommendationsResponse{
		RecommendedRoots: convertCharacterRootsForRecommendations(recommendations),
		Reason:          "基于你已掌握的字根，推荐学习相似难度的新字根",
		NextGoals: []string{
			"解锁更多生活相关的字根",
			"挑战更高难度的词汇",
			"探索不同语言的发音规律",
		},
	}

	return resp, nil
}

// convertTreasureMap 转换藏宝图格式
func convertTreasureMap(tm *hanbao.TreasureMap) *types.TreasureMap {
	return &types.TreasureMap{
		UserID:       tm.UserID,
		SessionID:    tm.SessionID,
		Roots:        convertCharacterRoots(tm.Roots),
		Vocabularies: convertVocabularies(tm.Vocabularies),
		Connections:  convertConnections(tm.Connections),
		Achievements: convertAchievements(tm.Achievements),
		Stats:        convertStats(tm.Stats),
	}
}

// convertVocabularies 转换词汇格式
func convertVocabularies(vocabMap map[string][]hanbao.Vocabulary) map[string][]types.Vocabulary {
	result := make(map[string][]types.Vocabulary)
	for key, vocabs := range vocabMap {
		result[key] = make([]types.Vocabulary, len(vocabs))
		for i, vocab := range vocabs {
			result[key][i] = types.Vocabulary{
				ID:            vocab.ID,
				RootID:        vocab.RootID,
				Language:      vocab.Language,
				Word:          vocab.Word,
				Romaji:        vocab.Romaji,
				Pronunciation: vocab.Pronunciation,
				Meaning:       vocab.Meaning,
				ReadType:      vocab.ReadType,
				Difficulty:    vocab.Difficulty,
				ExampleCount:  vocab.ExampleCount,
			}
		}
	}
	return result
}

// convertConnections 转换连接格式
func convertConnections(conns []hanbao.Connection) []types.Connection {
	result := make([]types.Connection, len(conns))
	for i, conn := range conns {
		result[i] = types.Connection{
			FromRootID:  conn.FromRootID,
			ToRootID:    conn.ToRootID,
			Type:        conn.Type,
			Description: conn.Description,
		}
	}
	return result
}

// convertAchievements 转换成就格式
func convertAchievements(achievements []hanbao.Achievement) []types.Achievement {
	result := make([]types.Achievement, len(achievements))
	for i, achievement := range achievements {
		result[i] = types.Achievement{
			ID:          achievement.ID,
			Name:        achievement.Name,
			Description: achievement.Description,
			Icon:        achievement.Icon,
			Condition:   achievement.Condition,
			Reward:      achievement.Reward,
		}
	}
	return result
}

// convertStats 转换统计格式
func convertStats(stats hanbao.SessionStats) types.SessionStats {
	return types.SessionStats{
		TotalRoots:     stats.TotalRoots,
		UnlockedRoots:  stats.UnlockedRoots,
		TotalWords:     stats.TotalWords,
		LearnedWords:   stats.LearnedWords,
		Accuracy:       stats.Accuracy,
		AverageTime:    stats.AverageTime,
		CompletionRate: stats.CompletionRate,
	}
}

// convertCharacterRootsForRecommendations 转换推荐字根格式
func convertCharacterRootsForRecommendations(roots []hanbao.CharacterRoot) []types.CharacterRoot {
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
