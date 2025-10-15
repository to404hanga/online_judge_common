package model

import "time"

type Problem struct {
	ID          uint64          `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                    // 题目 ID
	Title       string          `gorm:"column:title;type:varchar(255)" json:"title"`                            // 题目标题
	URL         string          `gorm:"column:url;type:varchar(255)" json:"url"`                                // 题目 url
	Status      *ProblemStatus  `gorm:"column:status;type:tinyint" json:"status"`                               // 题目状态(0:未发布, 1:已发布, 2:已删除)
	Visible     *ProblemVisible `gorm:"column:visible;type:tinyint" json:"visible"`                             // 题目可见性(0:不可见, 1:可见)
	TimeLimit   int             `gorm:"column:time_limit;type:int" json:"time_limit"`                           // 题目时间限制(单位:毫秒)
	MemoryLimit int             `gorm:"column:memory_limit;type:int" json:"memory_limit"`                       // 题目内存限制(单位:KB)
	CreatorID   uint64          `gorm:"column:creator_id;type:bigint unsigned" json:"creator_id"`               // 题目创建者 ID
	UpdaterID   uint64          `gorm:"column:updater_id;type:bigint unsigned" json:"updater_id"`               // 题目更新者 ID
	CreatedAt   time.Time       `gorm:"column:created_at;type:datetime;autoCreateTime:milli" json:"created_at"` // 创建时间
	UpdatedAt   time.Time       `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at"` // 更新时间
}

func (Problem) TableName() string {
	return "problem"
}

type ProblemStatus int8

const (
	ProblemStatusUnpublished ProblemStatus = iota // 未发布
	ProblemStatusPublished                        // 已发布
	ProblemStatusDeleted                          // 已删除
)

func (s *ProblemStatus) Int8() int8 {
	if s == nil {
		return int8(ProblemStatusUnpublished)
	}
	return int8(*s)
}

type ProblemVisible int8

const (
	ProblemVisibleFalse ProblemVisible = iota // 不可见
	ProblemVisibleTrue                        // 可见
)

func (v *ProblemVisible) Int8() int8 {
	if v == nil {
		return int8(ProblemVisibleFalse)
	}
	return int8(*v)
}
