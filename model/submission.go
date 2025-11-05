package model

import "time"

type Submission struct {
	ID            uint64              `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                                                               // 提交 ID
	CompetitionID uint64              `gorm:"column:competition_id;type:bigint unsigned;uniqueIndex:uk_submission_competition_problem_user_code_language" json:"competition_id"` // 比赛 ID
	UserID        uint64              `gorm:"column:user_id;type:bigint unsigned;uniqueIndex:uk_submission_competition_problem_user_code_language" json:"user_id"`               // 用户 ID
	ProblemID     uint64              `gorm:"column:problem_id;type:bigint unsigned;uniqueIndex:uk_submission_competition_problem_user_code_language" json:"problem_id"`         // 题目 ID
	Code          string              `gorm:"column:code;type:text" json:"code"`                                                                                                 // 提交代码
	CodeHash      string              `gorm:"column:code_hash;type:varchar(255);uniqueIndex:uk_submission_competition_problem_user_code_language" json:"code_hash"`              // 提交代码哈希值
	Language      *SubmissionLanguage `gorm:"column:language;type:tinyint;uniqueIndex:uk_submission_competition_problem_user_code_language" json:"language"`                     // 提交语言 ( 0: C/C++, 1: Python, 2: Java, 3: Go )
	Status        *SubmissionStatus   `gorm:"column:status;type:tinyint" json:"status"`                                                                                          // 提交状态 ( 0: 待判题, 1: 判题中, 2: 已判题 )
	Result        *SubmissionResult   `gorm:"column:result;type:tinyint" json:"result"`                                                                                          // 判题结果 ( 0: 未判题, 1: Accepted, 2: Wrong Answer, 3: Compile Error, 4: Runtime Error, 5: Time Limit Exceeded, 6: Memory Limit Exceeded, 7: Output Limit Exceeded )
	TimeUsed      *int                `gorm:"column:time_used;type:int" json:"time_used"`                                                                                        // 判题时间 ( 单位: 毫秒 )
	MemoryUsed    *int                `gorm:"column:memory_used;type:int" json:"memory_used"`                                                                                    // 判题内存 ( 单位: KB )
	CreatedAt     time.Time           `gorm:"column:created_at;type:datetime;autoCreateTime:milli" json:"created_at"`                                                            // 创建时间
	UpdatedAt     time.Time           `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at"`                                                            // 更新时间
}

func (Submission) TableName() string {
	return "submission"
}

type SubmissionLanguage int8

const (
	SubmissionLanguageC      SubmissionLanguage = iota // C/C++
	SubmissionLanguageCPP                              // C++
	SubmissionLanguagePython                           // Python
	SubmissionLanguageJava                             // Java
	SubmissionLanguageGo                               // Go
)

func (l *SubmissionLanguage) Int8() int8 {
	if l == nil {
		return int8(SubmissionLanguageC)
	}
	return int8(*l)
}

func (l *SubmissionLanguage) String() string {
	switch *l {
	case SubmissionLanguageC:
		return "C/C++"
	case SubmissionLanguageCPP:
		return "C++"
	case SubmissionLanguagePython:
		return "Python"
	case SubmissionLanguageJava:
		return "Java"
	case SubmissionLanguageGo:
		return "Go"
	default:
		return "未知语言"
	}
}

type SubmissionStatus int8

const (
	SubmissionStatusPending SubmissionStatus = iota // 待判题
	SubmissionStatusJudging                         // 判题中
	SubmissionStatusJudged                          // 已判题
)

func (s *SubmissionStatus) Int8() int8 {
	if s == nil {
		return int8(SubmissionStatusPending)
	}
	return int8(*s)
}

func (s *SubmissionStatus) String() string {
	switch s.Int8() {
	case int8(SubmissionStatusPending):
		return "待判题"
	case int8(SubmissionStatusJudging):
		return "判题中"
	case int8(SubmissionStatusJudged):
		return "已判题"
	default:
		return "未知"
	}
}

type SubmissionResult int8

const (
	SubmissionResultUnjudged            SubmissionResult = iota // 未判题
	SubmissionResultAccepted                                    // Accepted
	SubmissionResultWrongAnswer                                 // Wrong Answer
	SubmissionResultCompileError                                // Compile Error
	SubmissionResultRuntimeError                                // Runtime Error
	SubmissionResultTimeLimitExceeded                           // Time Limit Exceeded
	SubmissionResultMemoryLimitExceeded                         // Memory Limit Exceeded
	SubmissionResultOutputLimitExceeded                         // Output Limit Exceeded
)

func (r *SubmissionResult) Int8() int8 {
	if r == nil {
		return int8(SubmissionResultUnjudged)
	}
	return int8(*r)
}

func (r *SubmissionResult) String() string {
	switch r.Int8() {
	case int8(SubmissionResultUnjudged):
		return "Unjudged"
	case int8(SubmissionResultAccepted):
		return "Accepted"
	case int8(SubmissionResultWrongAnswer):
		return "Wrong Answer"
	case int8(SubmissionResultCompileError):
		return "Compile Error"
	case int8(SubmissionResultRuntimeError):
		return "Runtime Error"
	case int8(SubmissionResultTimeLimitExceeded):
		return "Time Limit Exceeded"
	case int8(SubmissionResultMemoryLimitExceeded):
		return "Memory Limit Exceeded"
	case int8(SubmissionResultOutputLimitExceeded):
		return "Output Limit Exceeded"
	default:
		return "Unknown"
	}
}
