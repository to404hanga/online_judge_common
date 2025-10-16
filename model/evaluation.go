package model

import "time"

type Evaluation struct {
	ID           uint64                    `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                                          // 判题 ID
	SubmissionID uint64                    `gorm:"column:submission_id;type:bigint unsigned;uniqueIndex:uk_evaluation_submission_testcase" json:"submission_id"` // 提交 ID
	TestcaseID   uint64                    `gorm:"column:testcase_id;type:bigint unsigned;uniqueIndex:uk_evaluation_submission_testcase" json:"testcase_id"`     // 测试用例 ID
	TimeUsed     int                       `gorm:"column:time_used;type:int;default:-1" json:"time_used"`                                                        // 运行时间 ( 单位: 毫秒, -1 表示未判题 )
	MemoryUsed   int                       `gorm:"column:memory_used;type:int;default:-1" json:"memory_used"`                                                    // 运行内存 ( 单位: KB, -1 表示未判题 )
	Accepted     *EvaluationAcceptedStatus `gorm:"column:accepted;type:tinyint;default:0" json:"accepted"`                                                       // 是否通过 ( 0: 未判题, 1: 已通过, 2: 未通过 )
	CreatedAt    time.Time                 `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`                                  // 创建时间
	UpdatedAt    time.Time                 `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`      // 更新时间
}

func (Evaluation) TableName() string {
	return "evaluation"
}

type EvaluationAcceptedStatus int8

const (
	EvaluationAcceptedStatusUnjudged   EvaluationAcceptedStatus = iota // 未判题
	EvaluationAcceptedStatusAccepted                                   // 已通过
	EvaluationAcceptedStatusUnaccepted                                 // 未通过
)

func (e *EvaluationAcceptedStatus) Int8() int8 {
	if e == nil {
		return int8(EvaluationAcceptedStatusUnjudged)
	}
	return int8(*e)
}

func (e *EvaluationAcceptedStatus) String() string {
	switch *e {
	case EvaluationAcceptedStatusUnjudged:
		return "未判题"
	case EvaluationAcceptedStatusAccepted:
		return "已通过"
	case EvaluationAcceptedStatusUnaccepted:
		return "未通过"
	default:
		return "未知状态"
	}
}
