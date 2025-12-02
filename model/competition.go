package model

import "time"

type Competition struct {
	ID        uint64             `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                       // 比赛 ID
	Name      string             `gorm:"column:name;type:varchar(255)" json:"name"`                                 // 比赛名称
	StartTime time.Time          `gorm:"column:start_time;type:datetime(3);index:idx_start_time" json:"start_time"` // 比赛开始时间
	EndTime   time.Time          `gorm:"column:end_time;type:datetime(3);index:idx_end_time" json:"end_time"`       // 比赛结束时间
	Status    *CompetitionStatus `gorm:"column:status;type:tinyint" json:"status"`                                  // 比赛状态(0:未发布, 1:已发布, 2:已删除)
	CreatorID uint64             `gorm:"column:creator_id;type:bigint unsigned" json:"creator_id"`                  // 创建者 ID
	UpdaterID uint64             `gorm:"column:updater_id;type:bigint unsigned" json:"updater_id"`                  // 更新者 ID
	CreatedAt time.Time          `gorm:"column:created_at;type:datetime(3);autoCreateTime:milli" json:"created_at"` // 创建时间
	UpdatedAt time.Time          `gorm:"column:updated_at;type:datetime(3);autoUpdateTime:milli" json:"updated_at"` // 更新时间
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

func (s *CompetitionStatus) String() string {
	switch *s {
	case CompetitionStatusUnpublished:
		return "未发布"
	case CompetitionStatusPublished:
		return "已发布"
	case CompetitionStatusDeleted:
		return "已删除"
	default:
		return "未知状态"
	}
}
