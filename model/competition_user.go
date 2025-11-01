package model

type CompetitionUser struct {
	ID            uint64                 `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                                                          // 比赛用户 ID
	CompetitionID uint64                 `gorm:"column:competition_id;type:bigint unsigned;not null;uniqueIndex:uk_competition_user_id" json:"competition_id"` // 比赛 ID
	UserID        uint64                 `gorm:"column:user_id;type:bigint unsigned;not null;uniqueIndex:uk_competition_user_id" json:"user_id"`               // 用户 ID
	Username      string                 `gorm:"column:username;type:varchar(50);uniqueIndex:uk_username" json:"username"`                                     // 用户名(学号)
	Realname      string                 `gorm:"column:realname;type:varchar(50)" json:"realname"`                                                             // 真实姓名
	Status        *CompetitionUserStatus `gorm:"column:status;type:tinyint;not null;default:0" json:"status"`                                                  // 状态 ( 0: 正常, 1: 禁用 )
	TotalTime     int                    `gorm:"column:total_time;type:int;not null;default:0" json:"total_time"`                                              // 总耗时 ( 单位: 分钟 )
	PassCount     int                    `gorm:"column:pass_count;type:int;not null;default:0" json:"pass_count"`                                              // 通过题目数
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
