package hanbao

import (
	"fmt"
	"unicode/utf8"
)

// UnlockCeremonyService 词根解锁仪式服务
type UnlockCeremonyService struct {
	roots      []CharacterRoot
	vocabularies []Vocabulary
}

// NewUnlockCeremonyService 创建解锁仪式服务
func NewUnlockCeremonyService() *UnlockCeremonyService {
	return &UnlockCeremonyService{
		roots:        CharacterRootsData,
		vocabularies: VocabularyData,
	}
}

// UnlockRequest 解锁请求
type UnlockRequest struct {
	Words []string `json:"words"` // 用户输入的词语，如 ["电话", "发现", "图书馆"]
}

// UnlockResult 解锁结果
type UnlockResult struct {
	InputWords     []string           `json:"input_words"`
	DetectedRoots  []CharacterRoot    `json:"detected_roots"`
	RootCount      int                `json:"root_count"`
	UnlockableWords int               `json:"unlockable_words"` // 可解锁的词汇总数
	WordBreakdown  map[string]int     `json:"word_breakdown"`   // 按语言分组的词汇数
	Insights       []string           `json:"insights"`         // AI洞察
}

// AnalyzeWords 分析用户输入的词语
func (s *UnlockCeremonyService) AnalyzeWords(req UnlockRequest) (*UnlockResult, error) {
	if len(req.Words) == 0 {
		return nil, fmt.Errorf("至少需要输入一个词语")
	}

	if len(req.Words) > 5 {
		return nil, fmt.Errorf("最多只能输入5个词语")
	}

	// 提取所有字根
	detectedRoots := make(map[int64]CharacterRoot)
	allChars := make([]string, 0)

	// 收集所有输入词语中的汉字
	for _, word := range req.Words {
		chars := s.splitIntoChars(word)
		allChars = append(allChars, chars...)
	}

	// 匹配字根
	for _, char := range allChars {
		for _, root := range s.roots {
			if root.Root == char {
				detectedRoots[root.ID] = root
				break
			}
		}
	}

	// 转换为切片
	roots := make([]CharacterRoot, 0, len(detectedRoots))
	for _, root := range detectedRoots {
		roots = append(roots, root)
	}

	// 计算可解锁的词汇
	wordBreakdown := make(map[string]int)
	unlockableWords := 0

	for _, vocab := range s.vocabularies {
		if _, exists := detectedRoots[vocab.RootID]; exists {
			wordBreakdown[vocab.Language]++
			unlockableWords++
		}
	}

	// 生成AI洞察
	insights := s.generateInsights(roots, wordBreakdown)

	return &UnlockResult{
		InputWords:     req.Words,
		DetectedRoots:  roots,
		RootCount:      len(roots),
		UnlockableWords: unlockableWords,
		WordBreakdown:  wordBreakdown,
		Insights:       insights,
	}, nil
}

// splitIntoChars 将中文词语拆分为单个汉字
func (s *UnlockCeremonyService) splitIntoChars(word string) []string {
	var chars []string
	for _, r := range word {
		if utf8.RuneLen(r) > 1 { // 中文字符
			chars = append(chars, string(r))
		}
	}
	return chars
}

// generateInsights 生成AI洞察
func (s *UnlockCeremonyService) generateInsights(roots []CharacterRoot, wordBreakdown map[string]int) []string {
	insights := make([]string, 0)

	// 基础洞察
	if len(roots) > 0 {
		insights = append(insights, fmt.Sprintf("你输入的%d个词中，有%d个汉字字根！", len(roots), len(roots)))
		insights = append(insights, fmt.Sprintf("这%d个字根能帮你解锁至少%d个日韩语词汇", len(roots), wordBreakdown["ja"]+wordBreakdown["ko"]))
	}

	// 语言分布洞察
	if jaCount, ok := wordBreakdown["ja"]; ok && jaCount > 0 {
		insights = append(insights, fmt.Sprintf("日语词汇：%d个（包括音读和训读）", jaCount))
	}
	if koCount, ok := wordBreakdown["ko"]; ok && koCount > 0 {
		insights = append(insights, fmt.Sprintf("韩语词汇：%d个（汉字词）", koCount))
	}

	// 难度分析
	easyRoots := 0
	for _, root := range roots {
		if root.Difficulty == 1 {
			easyRoots++
		}
	}
	if easyRoots > 0 {
		insights = append(insights, fmt.Sprintf("其中%d个是高频字根，特别适合入门学习", easyRoots))
	}

	// 文化洞察
	if len(roots) >= 3 {
		insights = append(insights, "汉字字根是一部活的语言迁徙史，日韩语中的汉字词都源于中国古代汉语")
	}

	// 任务建议
	totalWords := wordBreakdown["ja"] + wordBreakdown["ko"]
	if totalWords > 0 {
		insights = append(insights, fmt.Sprintf("今日任务：通过解谜，解锁其中%d个词汇", min(20, totalWords)))
	}

	return insights
}

// min 返回两个整数中的较小值
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// GetRootByID 根据ID获取字根
func (s *UnlockCeremonyService) GetRootByID(rootID int64) *CharacterRoot {
	for _, root := range s.roots {
		if root.ID == rootID {
			return &root
		}
	}
	return nil
}

// GetVocabulariesByRoot 获取指定字根的所有词汇
func (s *UnlockCeremonyService) GetVocabulariesByRoot(rootID int64) []Vocabulary {
	var result []Vocabulary
	for _, vocab := range s.vocabularies {
		if vocab.RootID == rootID {
			result = append(result, vocab)
		}
	}
	return result
}

// GetVocabulariesByLanguage 获取指定语言的所有词汇
func (s *UnlockCeremonyService) GetVocabulariesByLanguage(language string) []Vocabulary {
	var result []Vocabulary
	for _, vocab := range s.vocabularies {
		if vocab.Language == language {
			result = append(result, vocab)
		}
	}
	return result
}
