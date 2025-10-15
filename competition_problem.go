package online_judge_model

import "time"

type CompetitionProblem struct {
	ID            uint64                    `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                                 // 比赛题目 ID
	CompetitionID uint64                    `gorm:"column:competition_id;type:bigint unsigned;uniqueIndex:uk_competition_problem" json:"competition_id"` // 比赛 ID
	ProblemID     uint64                    `gorm:"column:problem_id;type:bigint unsigned;uniqueIndex:uk_competition_problem" json:"problem_id"`         // 题目 ID
	Status        *CompetitionProblemStatus `gorm:"column:status;type:tinyint" json:"status"`                                                            // 比赛题目状态 ( 0: 禁用, 1: 启用 )
	CreatedAt     time.Time                 `gorm:"column:created_at;type:datetime;autoCreateTime:milli" json:"created_at"`                              // 创建时间
	UpdatedAt     time.Time                 `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at"`                              // 更新时间
}

type CompetitionProblemStatus int8

const (
	CompetitionProblemStatusDisabled CompetitionProblemStatus = iota // 禁用
	CompetitionProblemStatusEnabled                                  // 启用
)

func (s *CompetitionProblemStatus) Int8() int8 {
	if s == nil {
		return int8(CompetitionProblemStatusDisabled)
	}
	return int8(*s)
}
