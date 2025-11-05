package model

import "time"

type Evaluation struct {
	ID             uint64         `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                                          // 判题 ID
	SubmissionID   uint64         `gorm:"column:submission_id;type:bigint unsigned;uniqueIndex:uk_evaluation_submission_testcase" json:"submission_id"` // 提交 ID
	TestcaseNumber int            `gorm:"column:testcase_number;type:int;uniqueIndex:uk_evaluation_submission_testcase" json:"testcase_number"`         // 测试用例编号
	Status         *ProblemStatus `gorm:"column:status;type:tinyint;default:0" json:"status"`                                                           // 评测结果 ( 0: 未评测, 1: Accepted, 2: Wrong Answer, 3: Compile Error, 4: Runtime Error, 5: Time Limit Exceeded, 6: Memory Limit Exceeded, 7: Output Limit Exceeded )
	TimeUsed       int            `gorm:"column:time_used;type:int;default:-1" json:"time_used"`                                                        // 运行时间 ( 单位: 毫秒, -1 表示未判题 )
	MemoryUsed     int            `gorm:"column:memory_used;type:int;default:-1" json:"memory_used"`                                                    // 运行内存 ( 单位: KB, -1 表示未判题 )
	CreatedAt      time.Time      `gorm:"column:created_at;type:datetime;default:CURRENT_TIMESTAMP" json:"created_at"`                                  // 创建时间
	UpdatedAt      time.Time      `gorm:"column:updated_at;type:datetime;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" json:"updated_at"`      // 更新时间
}

func (Evaluation) TableName() string {
	return "evaluation"
}
