package model

import "time"

type ProblemTestcase struct {
	ID        uint64                 `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                      // 测试用例 ID
	ProblemID uint64                 `gorm:"column:problem_id;type:bigint unsigned;uniqueIndex:uk_problem_testcase" json:"problem_id"` // 题目 ID
	InputURL  string                 `gorm:"column:input_url;type:varchar(255);uniqueIndex:uk_problem_testcase" json:"input_url"`      // 输入用例 URL
	OutputURL string                 `gorm:"column:output_url;type:varchar(255);uniqueIndex:uk_problem_testcase" json:"output_url"`    // 输出用例 URL
	Status    *ProblemTestcaseStatus `gorm:"column:status;type:tinyint;default:1" json:"status"`                                       // 测试用例状态 ( 0: 禁用, 1: 启用 )
	CreatedAt time.Time              `gorm:"column:created_at;type:datetime;autoCreateTime:milli" json:"created_at"`                   // 创建时间
	UpdatedAt time.Time              `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at"`                   // 更新时间
}

func (ProblemTestcase) TableName() string {
	return "problem_testcase"
}

type ProblemTestcaseStatus int8

const (
	ProblemTestcaseStatusDisabled ProblemTestcaseStatus = iota // 禁用
	ProblemTestcaseStatusEnabled                               // 启用
)

func (s *ProblemTestcaseStatus) Int8() int8 {
	if s == nil {
		return int8(ProblemTestcaseStatusEnabled)
	}
	return int8(*s)
}
