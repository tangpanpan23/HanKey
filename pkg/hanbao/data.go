package hanbao

import "time"

// Predefined character roots with their vocabulary
var CharacterRootsData = []CharacterRoot{
	// Tier 1 - High priority roots
	{ID: 1, Root: "电", Pinyin: "diàn", Difficulty: 1, Tier: 1, Description: "电力、电子相关", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, Root: "话", Pinyin: "huà", Difficulty: 1, Tier: 1, Description: "言语、对话", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 3, Root: "学", Pinyin: "xué", Difficulty: 1, Tier: 1, Description: "学习、教育", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 4, Root: "生", Pinyin: "shēng", Difficulty: 1, Tier: 1, Description: "生命、生产", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 5, Root: "国", Pinyin: "guó", Difficulty: 1, Tier: 1, Description: "国家、国际", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 6, Root: "家", Pinyin: "jiā", Difficulty: 1, Tier: 1, Description: "家庭、家居", CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// Tier 2 - Medium priority roots
	{ID: 7, Root: "发", Pinyin: "fā", Difficulty: 2, Tier: 2, Description: "发出、发展", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 8, Root: "现", Pinyin: "xiàn", Difficulty: 2, Tier: 2, Description: "显现、现在", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 9, Root: "图", Pinyin: "tú", Difficulty: 2, Tier: 2, Description: "图画、地图", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 10, Root: "书", Pinyin: "shū", Difficulty: 2, Tier: 2, Description: "书籍、书写", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 11, Root: "馆", Pinyin: "guǎn", Difficulty: 2, Tier: 2, Description: "馆舍、场所", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 12, Root: "文", Pinyin: "wén", Difficulty: 2, Tier: 2, Description: "文字、文化", CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 13, Root: "化", Pinyin: "huà", Difficulty: 2, Tier: 2, Description: "变化、化学", CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

var VocabularyData = []Vocabulary{
	// 电 (diàn) - Japanese examples
	{ID: 1, RootID: 1, Language: "ja", Word: "電話", Romaji: "denwa", Pronunciation: "でんわ", Meaning: "telephone", ReadType: "on", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 2, RootID: 1, Language: "ja", Word: "電気", Romaji: "denki", Pronunciation: "でんき", Meaning: "electricity", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 3, RootID: 1, Language: "ja", Word: "電車", Romaji: "densha", Pronunciation: "でんしゃ", Meaning: "train", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 4, RootID: 1, Language: "ja", Word: "電池", Romaji: "denchi", Pronunciation: "でんち", Meaning: "battery", ReadType: "on", Difficulty: 1, ExampleCount: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 电 (diàn) - Korean examples
	{ID: 5, RootID: 1, Language: "ko", Word: "전화", Pronunciation: "jeon-hwa", Meaning: "telephone", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 6, RootID: 1, Language: "ko", Word: "전기", Pronunciation: "jeon-gi", Meaning: "electricity", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 7, RootID: 1, Language: "ko", Word: "전철", Pronunciation: "jeon-cheol", Meaning: "electric train", Difficulty: 1, ExampleCount: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 话 (huà) - Japanese examples
	{ID: 8, RootID: 2, Language: "ja", Word: "会話", Romaji: "kaiwa", Pronunciation: "かいわ", Meaning: "conversation", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 9, RootID: 2, Language: "ja", Word: "電話", Romaji: "denwa", Pronunciation: "でんわ", Meaning: "telephone", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 话 (huà) - Korean examples
	{ID: 10, RootID: 2, Language: "ko", Word: "대화", Pronunciation: "dae-hwa", Meaning: "conversation", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 11, RootID: 2, Language: "ko", Word: "전화", Pronunciation: "jeon-hwa", Meaning: "telephone", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 学 (xué) - Japanese examples
	{ID: 12, RootID: 3, Language: "ja", Word: "学生", Romaji: "gakusei", Pronunciation: "がくせい", Meaning: "student", ReadType: "on", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 13, RootID: 3, Language: "ja", Word: "学校", Romaji: "gakkou", Pronunciation: "がっこう", Meaning: "school", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 14, RootID: 3, Language: "ja", Word: "大学", Romaji: "daigaku", Pronunciation: "だいがく", Meaning: "university", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 15, RootID: 3, Language: "ja", Word: "学習", Romaji: "gakushuu", Pronunciation: "がくしゅう", Meaning: "study/learning", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 学 (xué) - Korean examples
	{ID: 16, RootID: 3, Language: "ko", Word: "학생", Pronunciation: "hak-saeng", Meaning: "student", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 17, RootID: 3, Language: "ko", Word: "학교", Pronunciation: "hak-gyo", Meaning: "school", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 18, RootID: 3, Language: "ko", Word: "대학", Pronunciation: "dae-hak", Meaning: "university", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 生 (shēng) - Japanese examples
	{ID: 19, RootID: 4, Language: "ja", Word: "学生", Romaji: "gakusei", Pronunciation: "がくせい", Meaning: "student", ReadType: "on", Difficulty: 1, ExampleCount: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 20, RootID: 4, Language: "ja", Word: "生活", Romaji: "seikatsu", Pronunciation: "せいかつ", Meaning: "life/lifestyle", ReadType: "on", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 21, RootID: 4, Language: "ja", Word: "生命", Romaji: "seimei", Pronunciation: "せいめい", Meaning: "life", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 22, RootID: 4, Language: "ja", Word: "生物", Romaji: "seibutsu", Pronunciation: "せいぶつ", Meaning: "living things", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 23, RootID: 4, Language: "ja", Word: "生鮮", Romaji: "seisen", Pronunciation: "せいせん", Meaning: "fresh food", ReadType: "on", Difficulty: 1, ExampleCount: 1, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 生 (shēng) - Korean examples
	{ID: 24, RootID: 4, Language: "ko", Word: "학생", Pronunciation: "hak-saeng", Meaning: "student", Difficulty: 1, ExampleCount: 4, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 25, RootID: 4, Language: "ko", Word: "생활", Pronunciation: "saeng-hwal", Meaning: "life/lifestyle", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 26, RootID: 4, Language: "ko", Word: "생명", Pronunciation: "saeng-myeong", Meaning: "life", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 27, RootID: 4, Language: "ko", Word: "생물", Pronunciation: "saeng-mul", Meaning: "living things", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 国 (guó) - Japanese examples
	{ID: 28, RootID: 5, Language: "ja", Word: "中国", Romaji: "chuugoku", Pronunciation: "ちゅうごく", Meaning: "China", ReadType: "on", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 29, RootID: 5, Language: "ja", Word: "外国", Romaji: "gaikoku", Pronunciation: "がいこく", Meaning: "foreign country", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 30, RootID: 5, Language: "ja", Word: "国際", Romaji: "kokusai", Pronunciation: "こくさい", Meaning: "international", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 国 (guó) - Korean examples
	{ID: 31, RootID: 5, Language: "ko", Word: "중국", Pronunciation: "jung-guk", Meaning: "China", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 32, RootID: 5, Language: "ko", Word: "외국", Pronunciation: "oe-guk", Meaning: "foreign country", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 33, RootID: 5, Language: "ko", Word: "국제", Pronunciation: "guk-je", Meaning: "international", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 家 (jiā) - Japanese examples
	{ID: 34, RootID: 6, Language: "ja", Word: "家庭", Romaji: "katei", Pronunciation: "かてい", Meaning: "family", ReadType: "on", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 35, RootID: 6, Language: "ja", Word: "家", Romaji: "ie", Pronunciation: "いえ", Meaning: "home/house", ReadType: "kun", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 36, RootID: 6, Language: "ja", Word: "家族", Romaji: "kazoku", Pronunciation: "かぞく", Meaning: "family", ReadType: "on", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},

	// 家 (jiā) - Korean examples
	{ID: 37, RootID: 6, Language: "ko", Word: "가족", Pronunciation: "ga-jok", Meaning: "family", Difficulty: 1, ExampleCount: 3, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 38, RootID: 6, Language: "ko", Word: "가정", Pronunciation: "ga-jeong", Meaning: "home/family", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
	{ID: 39, RootID: 6, Language: "ko", Word: "집", Pronunciation: "jip", Meaning: "home/house", Difficulty: 1, ExampleCount: 2, CreatedAt: time.Now(), UpdatedAt: time.Now()},
}

// Dialect examples for level 3 challenges
var DialectExamplesData = []DialectExample{
	{ID: 1, RootID: 3, Standard: "吃饭", Dialect: "食饭", DialectType: "cantonese", Description: "粤语中'吃'读作'食'", AudioURL: ""},
	{ID: 2, RootID: 3, Standard: "吃饭", Dialect: "食飯", DialectType: "minnan", Description: "闽南语中'吃'读作'食'", AudioURL: ""},
	{ID: 3, RootID: 4, Standard: "学生", Dialect: "學生", DialectType: "cantonese", Description: "粤语保留古汉语发音", AudioURL: ""},
}
