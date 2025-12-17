package hanbao

import "time"

// CharacterRoot 汉字字根
type CharacterRoot struct {
	ID          int64     `json:"id" db:"id"`
	Root        string    `json:"root" db:"root"`               // 字根汉字，如 "电"
	Pinyin      string    `json:"pinyin" db:"pinyin"`           // 拼音，如 "diàn"
	Difficulty  int       `json:"difficulty" db:"difficulty"`   // 难度等级 1-3
	Tier        int       `json:"tier" db:"tier"`               // 优先级层级 1-3
	Description string    `json:"description" db:"description"` // 字根描述
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// Vocabulary 词汇信息
type Vocabulary struct {
	ID             int64         `json:"id" db:"id"`
	RootID         int64         `json:"root_id" db:"root_id"`               // 关联的字根ID
	Language       string        `json:"language" db:"language"`             // 语言: "ja" 或 "ko"
	Word           string        `json:"word" db:"word"`                     // 词汇，如 "電話"
	Romaji         string        `json:"romaji,omitempty" db:"romaji"`       // 日语罗马字，如 "denwa"
	Pronunciation  string        `json:"pronunciation" db:"pronunciation"`   // 发音，如 "でんわ"
	Meaning        string        `json:"meaning" db:"meaning"`               // 含义，如 "telephone"
	ReadType       string        `json:"read_type,omitempty" db:"read_type"` // 读音类型: "on" 或 "kun" (日语)
	Difficulty     int           `json:"difficulty" db:"difficulty"`         // 难度等级
	ExampleCount   int           `json:"example_count" db:"example_count"`   // 示例数量
	CreatedAt      time.Time     `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at" db:"updated_at"`
}

// UserSession 用户会话
type UserSession struct {
	ID            string    `json:"id" db:"id"`
	UserID        string    `json:"user_id" db:"user_id"`                 // 用户标识（可匿名）
	UnlockedRoots []int64   `json:"unlocked_roots" db:"unlocked_roots"`   // 已解锁的字根ID列表
	CompletedLevels []string `json:"completed_levels" db:"completed_levels"` // 已完成的关卡ID
	Score         int       `json:"score" db:"score"`                     // 总得分
	Accuracy      float64   `json:"accuracy" db:"accuracy"`               // 准确率
	StartTime     time.Time `json:"start_time" db:"start_time"`
	LastActive    time.Time `json:"last_active" db:"last_active"`
	Status        string    `json:"status" db:"status"`                   // 会话状态: "active", "completed"
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

// Level 关卡
type Level struct {
	ID          string `json:"id" db:"id"`
	Type        string `json:"type" db:"type"`               // 关卡类型: "pronunciation", "listening", "dialect"
	Title       string `json:"title" db:"title"`             // 关卡标题
	Description string `json:"description" db:"description"` // 关卡描述
	RootID      int64  `json:"root_id" db:"root_id"`         // 关联字根ID
	Difficulty  int    `json:"difficulty" db:"difficulty"`   // 难度等级
	TimeLimit   int    `json:"time_limit" db:"time_limit"`   // 时间限制（秒）
	Questions   []Question `json:"questions" db:"questions"` // 问题列表
	Reward      Reward  `json:"reward" db:"reward"`          // 奖励
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// Question 问题
type Question struct {
	ID          string   `json:"id"`
	Type        string   `json:"type"`        // 问题类型: "multiple_choice", "text_input", "audio_match"
	Content     string   `json:"content"`     // 问题内容
	Options     []string `json:"options,omitempty"` // 选项（选择题）
	CorrectAnswer string `json:"correct_answer"` // 正确答案
	Hint        string   `json:"hint,omitempty"` // 提示
	Explanation string   `json:"explanation"` // 解释
}

// Reward 奖励
type Reward struct {
	Roots   []int64 `json:"roots"`   // 解锁的字根
	Score   int     `json:"score"`   // 得分
	Achievement string `json:"achievement,omitempty"` // 成就
}

// Achievement 成就
type Achievement struct {
	ID          string `json:"id" db:"id"`
	Name        string `json:"name" db:"name"`
	Description string `json:"description" db:"description"`
	Icon        string `json:"icon" db:"icon"`
	Condition   string `json:"condition" db:"condition"` // 达成条件
	Reward      string `json:"reward" db:"reward"`       // 奖励描述
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
}

// DialectExample 方言示例
type DialectExample struct {
	ID          int64  `json:"id" db:"id"`
	RootID      int64  `json:"root_id" db:"root_id"`
	Standard    string `json:"standard" db:"standard"`       // 标准汉语，如 "吃饭"
	Dialect     string `json:"dialect" db:"dialect"`         // 方言，如 "食饭" (粤语)
	DialectType string `json:"dialect_type" db:"dialect_type"` // 方言类型，如 "cantonese", "minnan"
	Description string `json:"description" db:"description"` // 文化解释
	AudioURL    string `json:"audio_url,omitempty" db:"audio_url"` // 音频链接
}

// TreasureMap 藏宝图
type TreasureMap struct {
	UserID       string                    `json:"user_id"`
	SessionID    string                    `json:"session_id"`
	Roots        []CharacterRoot           `json:"roots"`         // 已解锁的字根
	Vocabularies map[string][]Vocabulary  `json:"vocabularies"` // 按字根分组的词汇
	Connections  []Connection             `json:"connections"`   // 字根连接关系
	Achievements []Achievement            `json:"achievements"` // 获得的成就
	Stats        SessionStats             `json:"stats"`        // 会话统计
}

// Connection 字根连接关系
type Connection struct {
	FromRootID int64  `json:"from_root_id"`
	ToRootID   int64  `json:"to_root_id"`
	Type       string `json:"type"` // 连接类型: "derivative", "similar", "compound"
	Description string `json:"description"`
}

// SessionStats 会话统计
type SessionStats struct {
	TotalRoots     int     `json:"total_roots"`      // 总字根数
	UnlockedRoots  int     `json:"unlocked_roots"`   // 已解锁字根数
	TotalWords     int     `json:"total_words"`      // 总词汇数
	LearnedWords   int     `json:"learned_words"`    // 已学习词汇数
	Accuracy       float64 `json:"accuracy"`         // 准确率
	AverageTime    int     `json:"average_time"`     // 平均用时（秒）
	CompletionRate float64 `json:"completion_rate"`  // 完成率
}
