package model

import "time"

type CompetitionUser struct {
	ID            uint64                 `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                                                                                // 比赛用户 ID
	CompetitionID uint64                 `gorm:"column:competition_id;type:bigint unsigned;not null;uniqueIndex:uk_competition_user_id;index:idx_ranking,priority:1,sort:asc" json:"competition_id"` // 比赛 ID
	UserID        uint64                 `gorm:"column:user_id;type:bigint unsigned;not null;uniqueIndex:uk_competition_user_id" json:"user_id"`                                                     // 用户 ID
	Username      string                 `gorm:"column:username;type:varchar(50);uniqueIndex:uk_username" json:"username"`                                                                           // 用户名(学号)
	Realname      string                 `gorm:"column:realname;type:varchar(50)" json:"realname"`                                                                                                   // 真实姓名
	Status        *CompetitionUserStatus `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`                                                                                        // 状态 ( 0: 正常, 1: 禁用 )
	PassCount     int                    `gorm:"column:pass_count;type:int;not null;default:0;index:idx_ranking,priority:2,sort:desc" json:"pass_count"`                                             // 通过题目数
	TotalTime     int64                  `gorm:"column:total_time;type:bigint;not null;default:0;index:idx_ranking,priority:3,sort:asc" json:"total_time"`                                           // 总耗时 ( 单位: 毫秒 )
	RetryCount    int                    `gorm:"column:retry_count;type:int;not null;default:0" json:"retry_count"`                                                                                  // 重试次数
	StartTime     time.Time              `gorm:"column:start_time;type:datetime(3);not null" json:"start_time"`                                                                                      // 比赛开始时间
}

type CompetitionUserStatus int8

const (
	CompetitionUserStatusNormal   CompetitionUserStatus = iota // 正常
	CompetitionUserStatusDisabled                              // 禁用
)

func (c *CompetitionUserStatus) String() string {
	switch *c {
	case CompetitionUserStatusNormal:
		return "normal"
	case CompetitionUserStatusDisabled:
		return "disabled"
	default:
		return "unknown"
	}
}

func (c *CompetitionUserStatus) Int8() int8 {
	if c == nil {
		return 0
	}
	return int8(*c)
}
