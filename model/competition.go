package model

import "time"

type Competition struct {
	ID        uint64              `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                    // 比赛 ID
	Name      string              `gorm:"column:name;type:varchar(255)" json:"name"`                              // 比赛名称
	StartTime time.Time           `gorm:"column:start_time;type:datetime;index:idx_start_time" json:"start_time"` // 比赛开始时间
	EndTime   time.Time           `gorm:"column:end_time;type:datetime;index:idx_end_time" json:"end_time"`       // 比赛结束时间
	Status    *CompetitionStatus  `gorm:"column:status;type:tinyint" json:"status"`                               // 比赛状态(0:未发布, 1:已发布, 2:已删除)
	Visible   *CompetitionVisible `gorm:"column:visible;type:tinyint" json:"visible"`                             // 是否仅比赛时可见(0:否, 1:是)
	CreatorID uint64              `gorm:"column:creator_id;type:bigint unsigned" json:"creator_id"`               // 创建者 ID
	UpdaterID uint64              `gorm:"column:updater_id;type:bigint unsigned" json:"updater_id"`               // 更新者 ID
	CreatedAt time.Time           `gorm:"column:created_at;type:datetime;autoCreateTime:milli" json:"created_at"` // 创建时间
	UpdatedAt time.Time           `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at"` // 更新时间
}

func (Competition) TableName() string {
	return "competition"
}

type CompetitionStatus int8

const (
	CompetitionStatusUnpublished CompetitionStatus = iota // 未发布
	CompetitionStatusPublished                            // 已发布
	CompetitionStatusDeleted                              // 已删除
)

func (s *CompetitionStatus) Int8() int8 {
	if s == nil {
		return int8(CompetitionStatusUnpublished)
	}
	return int8(*s)
}

type CompetitionVisible int8

const (
	CompetitionVisibleFalse CompetitionVisible = iota // 不可见
	CompetitionVisibleTrue                            // 可见
)

func (v *CompetitionVisible) Int8() int8 {
	if v == nil {
		return int8(CompetitionVisibleFalse)
	}
	return int8(*v)
}
