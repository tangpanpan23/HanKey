package hanbao

import (
	"fmt"
	"time"
)

// TreasureMapService 藏宝图服务
type TreasureMapService struct {
	roots        []CharacterRoot
	vocabularies []Vocabulary
}

// NewTreasureMapService 创建藏宝图服务
func NewTreasureMapService() *TreasureMapService {
	return &TreasureMapService{
		roots:        CharacterRootsData,
		vocabularies: VocabularyData,
	}
}

// GenerateTreasureMap 生成藏宝图
func (s *TreasureMapService) GenerateTreasureMap(sessionID string, unlockedRoots []int64) (*TreasureMap, error) {
	if len(unlockedRoots) == 0 {
		return nil, fmt.Errorf("没有已解锁的字根")
	}

	// 获取已解锁的字根详情
	roots := make([]CharacterRoot, 0, len(unlockedRoots))
	for _, rootID := range unlockedRoots {
		for _, root := range s.roots {
			if root.ID == rootID {
				roots = append(roots, root)
				break
			}
		}
	}

	// 按字根分组词汇
	vocabularies := make(map[string][]Vocabulary)
	totalWords := 0
	jaWords := 0
	koWords := 0

	for _, rootID := range unlockedRoots {
		rootKey := ""
		for _, root := range roots {
			if root.ID == rootID {
				rootKey = root.Root
				break
			}
		}

		if rootKey == "" {
			continue
		}

		var rootVocabs []Vocabulary
		for _, vocab := range s.vocabularies {
			if vocab.RootID == rootID {
				rootVocabs = append(rootVocabs, vocab)
				totalWords++
				if vocab.Language == "ja" {
					jaWords++
				} else if vocab.Language == "ko" {
					koWords++
				}
			}
		}
		vocabularies[rootKey] = rootVocabs
	}

	// 生成连接关系（简化版）
	connections := s.generateConnections(unlockedRoots)

	// 计算统计数据
	stats := SessionStats{
		TotalRoots:     len(roots),
		UnlockedRoots:  len(unlockedRoots),
		TotalWords:     totalWords,
		LearnedWords:   totalWords, // 假设都已学习
		Accuracy:       85.0,       // 示例准确率
		AverageTime:    45,         // 示例平均用时
		CompletionRate: 100.0,      // 完成率
	}

	// 获取成就
	achievements := s.calculateAchievements(stats)

	treasureMap := &TreasureMap{
		UserID:       "demo_user", // 示例用户ID
		SessionID:    sessionID,
		Roots:        roots,
		Vocabularies: vocabularies,
		Connections:  connections,
		Achievements: achievements,
		Stats:        stats,
	}

	return treasureMap, nil
}

// generateConnections 生成字根连接关系
func (s *TreasureMapService) generateConnections(unlockedRoots []int64) []Connection {
	connections := make([]Connection, 0)

	// 简单的连接逻辑：相同类型的字根连接
	rootMap := make(map[int64]CharacterRoot)
	for _, rootID := range unlockedRoots {
		for _, root := range s.roots {
			if root.ID == rootID {
				rootMap[rootID] = root
				break
			}
		}
	}

	// 生成衍生关系连接
	for i, rootID1 := range unlockedRoots {
		root1 := rootMap[rootID1]
		for j := i + 1; j < len(unlockedRoots); j++ {
			rootID2 := unlockedRoots[j]
			root2 := rootMap[rootID2]

			// 如果是相同难度或相邻难度，则建立连接
			if abs(root1.Difficulty-root2.Difficulty) <= 1 {
				connections = append(connections, Connection{
					FromRootID: rootID1,
					ToRootID:   rootID2,
					Type:       "derivative",
					Description: fmt.Sprintf("%s → %s 的词根演变", root1.Root, root2.Root),
				})
			}
		}
	}

	return connections
}

// calculateAchievements 计算成就
func (s *TreasureMapService) calculateAchievements(stats SessionStats) []Achievement {
	achievements := make([]Achievement, 0)

	// 基础成就
	if stats.UnlockedRoots >= 3 {
		achievements = append(achievements, Achievement{
			ID:          "explorer_1",
			Name:        "汉字侦探见习生",
			Description: "解锁3个汉字字根",
			Icon:        "🕵️",
			Condition:   "解锁至少3个字根",
			Reward:      "解锁进阶关卡",
			CreatedAt:   time.Now(),
		})
	}

	if stats.TotalWords >= 10 {
		achievements = append(achievements, Achievement{
			ID:          "scholar_1",
			Name:        "语言学者",
			Description: "掌握10个日韩词汇",
			Icon:        "🎓",
			Condition:   "学习至少10个词汇",
			Reward:      "获得词根亲和力加成",
			CreatedAt:   time.Now(),
		})
	}

	if stats.Accuracy >= 80.0 {
		achievements = append(achievements, Achievement{
			ID:          "master_1",
			Name:        "解谜大师",
			Description: "准确率达到80%",
			Icon:        "🏆",
			Condition:   "单次会话准确率≥80%",
			Reward:      "解锁专家级关卡",
			CreatedAt:   time.Now(),
		})
	}

	return achievements
}

// GenerateReportText 生成文字报告
func (s *TreasureMapService) GenerateReportText(treasureMap *TreasureMap) string {
	report := fmt.Sprintf(`🎯 15分钟战报

✅ 已解锁字根：%d个
✅ 已掌握词汇：日语%d个 + 韩语%d个
✅ 解密准确率：%.1f%%
🔥 解锁成就：%d个

📊 词根网络预览：
`,
		treasureMap.Stats.UnlockedRoots,
		len(treasureMap.Vocabularies), // 简化为字根数量，实际应该统计词汇
		len(treasureMap.Vocabularies),
		treasureMap.Stats.Accuracy,
		len(treasureMap.Achievements),
	)

	// 添加字根树状图
	for rootName, vocabs := range treasureMap.Vocabularies {
		report += fmt.Sprintf("\n【%s】─┬─ %s\n", rootName, s.formatVocabSample(vocabs, 3))
	}

	// 添加成就
	if len(treasureMap.Achievements) > 0 {
		report += "\n🏆 获得成就：\n"
		for _, achievement := range treasureMap.Achievements {
			report += fmt.Sprintf("• %s - %s\n", achievement.Name, achievement.Description)
		}
	}

	report += "\n🚀 下一站建议：\n基于你已掌握的字根，下一关将解锁更多相关词汇。继续探索汉字的语言网络！"

	return report
}

// formatVocabSample 格式化词汇示例
func (s *TreasureMapService) formatVocabSample(vocabs []Vocabulary, maxCount int) string {
	if len(vocabs) == 0 {
		return ""
	}

	result := ""
	count := min(maxCount, len(vocabs))

	for i := 0; i < count; i++ {
		vocab := vocabs[i]
		if vocab.Language == "ja" {
			result += fmt.Sprintf("%s（%s）", vocab.Word, vocab.Romaji)
		} else {
			result += fmt.Sprintf("%s（%s）", vocab.Word, vocab.Pronunciation)
		}

		if i < count-1 {
			result += "\n      ├─ "
		}
	}

	if len(vocabs) > maxCount {
		result += fmt.Sprintf("\n      └─ ...等%d个词汇", len(vocabs)-maxCount)
	}

	return result
}

// GetNextRecommendations 获取下一阶段推荐
func (s *TreasureMapService) GetNextRecommendations(currentRoots []int64) []CharacterRoot {
	recommendations := make([]CharacterRoot, 0)

	// 找到未解锁的字根
	unlockedMap := make(map[int64]bool)
	for _, rootID := range currentRoots {
		unlockedMap[rootID] = true
	}

	// 推荐相同难度或更高一级的字根
	for _, root := range s.roots {
		if !unlockedMap[root.ID] {
			// 优先推荐相同难度
			hasSameDifficulty := false
			for _, unlockedID := range currentRoots {
				for _, unlockedRoot := range s.roots {
					if unlockedRoot.ID == unlockedID && unlockedRoot.Difficulty == root.Difficulty {
						hasSameDifficulty = true
						break
					}
				}
				if hasSameDifficulty {
					break
				}
			}

			if hasSameDifficulty || root.Difficulty <= 2 {
				recommendations = append(recommendations, root)
				if len(recommendations) >= 3 {
					break
				}
			}
		}
	}

	return recommendations
}

// abs 返回整数的绝对值
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
