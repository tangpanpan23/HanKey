package hanbao

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// LevelService 关卡服务
type LevelService struct {
	roots         []CharacterRoot
	vocabularies  []Vocabulary
	dialectExamples []DialectExample
	rng           *rand.Rand
}

// NewLevelService 创建关卡服务
func NewLevelService() *LevelService {
	return &LevelService{
		roots:           CharacterRootsData,
		vocabularies:    VocabularyData,
		dialectExamples: DialectExamplesData,
		rng:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateLevel 生成关卡
func (s *LevelService) GenerateLevel(levelType string, rootID int64, difficulty int) (*Level, error) {
	switch levelType {
	case "pronunciation":
		return s.generatePronunciationLevel(rootID, difficulty)
	case "listening":
		return s.generateListeningLevel(rootID, difficulty)
	case "dialect":
		return s.generateDialectLevel(rootID, difficulty)
	default:
		return nil, fmt.Errorf("不支持的关卡类型: %s", levelType)
	}
}

// generatePronunciationLevel 生成音读破译室关卡
func (s *LevelService) generatePronunciationLevel(rootID int64, difficulty int) (*Level, error) {
	root := s.findRootByID(rootID)
	if root == nil {
		return nil, fmt.Errorf("字根不存在: %d", rootID)
	}

	// 获取相关的日语词汇
	jaVocabs := s.getVocabulariesByRootAndLanguage(rootID, "ja")
	if len(jaVocabs) < 2 {
		return nil, fmt.Errorf("字根 %s 没有足够的日语词汇数据", root.Root)
	}

	// 随机选择两个词汇进行比较
	vocab1 := jaVocabs[s.rng.Intn(len(jaVocabs))]
	var vocab2 Vocabulary
	for {
		vocab2 = jaVocabs[s.rng.Intn(len(jaVocabs))]
		if vocab2.ID != vocab1.ID {
			break
		}
	}

	questions := []Question{
		{
			ID:   fmt.Sprintf("pron_%d_1", time.Now().Unix()),
			Type: "multiple_choice",
			Content: fmt.Sprintf("这两个日语词中相同的\"%s\"，读音有何规律？\n• %s（%s）\n• %s（%s）",
				root.Root, vocab1.Word, vocab1.Romaji, vocab2.Word, vocab2.Romaji),
			Options: []string{
				"模仿了古汉语的不同方言层次",
				"完全相同的发音",
				"现代汉语的标准发音",
				"随机的发音变化",
			},
			CorrectAnswer: "模仿了古汉语的不同方言层次",
			Hint:         fmt.Sprintf("中文\"%s\"在不同语境下的发音差异", root.Root),
			Explanation:  fmt.Sprintf("日语中的汉字词继承了中国古代汉语的读音层次，反映了历史上的语言演变"),
		},
	}

	level := &Level{
		ID:          fmt.Sprintf("pron_%d_%d", rootID, time.Now().Unix()),
		Type:        "pronunciation",
		Title:       "音读破译室 🔊",
		Description: fmt.Sprintf("探索\"%s\"在日语中的发音奥秘", root.Root),
		RootID:      rootID,
		Difficulty:  difficulty,
		TimeLimit:   180, // 3分钟
		Questions:   questions,
		Reward: Reward{
			Roots: []int64{rootID},
			Score: 100,
		},
		CreatedAt: time.Now(),
	}

	return level, nil
}

// generateListeningLevel 生成韩语听力侦探关卡
func (s *LevelService) generateListeningLevel(rootID int64, difficulty int) (*Level, error) {
	root := s.findRootByID(rootID)
	if root == nil {
		return nil, fmt.Errorf("字根不存在: %d", rootID)
	}

	// 获取相关的韩语词汇
	koVocabs := s.getVocabulariesByRootAndLanguage(rootID, "ko")
	if len(koVocabs) < 3 {
		return nil, fmt.Errorf("字根 %s 没有足够的韩语词汇数据", root.Root)
	}

	// 随机选择3个词汇
	selectedVocabs := make([]Vocabulary, 0, 3)
	usedIndices := make(map[int]bool)
	for len(selectedVocabs) < 3 && len(usedIndices) < len(koVocabs) {
		idx := s.rng.Intn(len(koVocabs))
		if !usedIndices[idx] {
			usedIndices[idx] = true
			selectedVocabs = append(selectedVocabs, koVocabs[idx])
		}
	}

	// 构建问题内容
	var vocabList strings.Builder
	for i, vocab := range selectedVocabs {
		vocabList.WriteString(fmt.Sprintf("%d. %s (%s)\n", i+1, vocab.Word, vocab.Pronunciation))
	}

	questions := []Question{
		{
			ID:   fmt.Sprintf("listen_%d_1", time.Now().Unix()),
			Type: "text_input",
			Content: fmt.Sprintf("请聆听这段韩语内容，圈出你听到的、像中文的词汇：\n\n%s\n\n你听到了几个像中文的词？",
				vocabList.String()),
			CorrectAnswer: fmt.Sprintf("%d", len(selectedVocabs)),
			Hint:         "韩语70%正式词汇是汉字词，听起来很熟悉",
			Explanation:  fmt.Sprintf("韩语中的汉字词直接借用汉字的音和义，%s相关的词汇都源于中文", root.Root),
		},
	}

	level := &Level{
		ID:          fmt.Sprintf("listen_%d_%d", rootID, time.Now().Unix()),
		Type:        "listening",
		Title:       "韩语听力侦探 🎧",
		Description: fmt.Sprintf("在韩语中寻找\"%s\"的身影", root.Root),
		RootID:      rootID,
		Difficulty:  difficulty,
		TimeLimit:   240, // 4分钟
		Questions:   questions,
		Reward: Reward{
			Roots: []int64{rootID},
			Score: 150,
		},
		CreatedAt: time.Now(),
	}

	return level, nil
}

// generateDialectLevel 生成方言连接彩蛋关卡
func (s *LevelService) generateDialectLevel(rootID int64, difficulty int) (*Level, error) {
	root := s.findRootByID(rootID)
	if root == nil {
		return nil, fmt.Errorf("字根不存在: %d", rootID)
	}

	// 查找相关的方言例子
	var dialectExample *DialectExample
	for _, example := range s.dialectExamples {
		if example.RootID == rootID {
			dialectExample = &example
			break
		}
	}

	if dialectExample == nil {
		return nil, fmt.Errorf("字根 %s 没有方言数据", root.Root)
	}

	questions := []Question{
		{
			ID:   fmt.Sprintf("dialect_%d_1", time.Now().Unix()),
			Type: "multiple_choice",
			Content: fmt.Sprintf("用你的方言说\"%s\"，会怎么说？\n\n标准汉语：%s\n%s方言：%s",
				dialectExample.Standard, dialectExample.Standard, dialectExample.DialectType, dialectExample.Dialect),
			Options: []string{
				fmt.Sprintf("与%s发音相似", dialectExample.Dialect),
				"完全不同",
				"标准汉语发音",
				"现代普通话发音",
			},
			CorrectAnswer: fmt.Sprintf("与%s发音相似", dialectExample.Dialect),
			Hint:         "汉字读音是一部活的迁徙史",
			Explanation:  dialectExample.Description,
		},
	}

	level := &Level{
		ID:          fmt.Sprintf("dialect_%d_%d", rootID, time.Now().Unix()),
		Type:        "dialect",
		Title:       "方言连接彩蛋 🗺️",
		Description: fmt.Sprintf("探索\"%s\"的方言奥秘", root.Root),
		RootID:      rootID,
		Difficulty:  difficulty,
		TimeLimit:   120, // 2分钟
		Questions:   questions,
		Reward: Reward{
			Roots: []int64{rootID},
			Score: 80,
		},
		CreatedAt: time.Now(),
	}

	return level, nil
}

// ValidateAnswer 验证答案
func (s *LevelService) ValidateAnswer(levelID string, questionID string, userAnswer string) (*AnswerResult, error) {
	// 这里简化处理，实际应该从存储中获取关卡数据
	// 现在直接返回正确的结果用于演示

	return &AnswerResult{
		Correct:     true,
		Score:       100,
		Explanation: "回答正确！汉字的语言演变史真是迷人",
		NextHint:    "继续探索更多汉字词根的奥秘",
	}, nil
}

// AnswerResult 答案验证结果
type AnswerResult struct {
	Correct     bool   `json:"correct"`
	Score       int    `json:"score"`
	Explanation string `json:"explanation"`
	NextHint    string `json:"next_hint,omitempty"`
}

// Helper methods
func (s *LevelService) findRootByID(rootID int64) *CharacterRoot {
	for _, root := range s.roots {
		if root.ID == rootID {
			return &root
		}
	}
	return nil
}

func (s *LevelService) getVocabulariesByRootAndLanguage(rootID int64, language string) []Vocabulary {
	var result []Vocabulary
	for _, vocab := range s.vocabularies {
		if vocab.RootID == rootID && vocab.Language == language {
			result = append(result, vocab)
		}
	}
	return result
}

// GenerateSessionLevels 为用户会话生成关卡序列
func (s *LevelService) GenerateSessionLevels(unlockedRoots []int64) ([]Level, error) {
	if len(unlockedRoots) == 0 {
		return nil, fmt.Errorf("没有可用的字根")
	}

	var levels []Level
	levelCount := min(5, len(unlockedRoots)*2) // 每个字根最多2个关卡

	// 随机选择字根和关卡类型
	levelTypes := []string{"pronunciation", "listening", "dialect"}

	for i := 0; i < levelCount; i++ {
		rootID := unlockedRoots[s.rng.Intn(len(unlockedRoots))]
		levelType := levelTypes[s.rng.Intn(len(levelTypes))]

		level, err := s.GenerateLevel(levelType, rootID, 1)
		if err != nil {
			continue // 跳过无法生成的关卡
		}

		levels = append(levels, *level)
	}

	return levels, nil
}
