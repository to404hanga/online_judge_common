package model

import "time"

type User struct {
	ID        uint64      `gorm:"column:id;type:bigint unsigned;primaryKey" json:"id"`                      // 用户 ID
	Username  string      `gorm:"column:username;type:varchar(50);uniqueIndex:uk_username" json:"username"` // 用户名(学号)
	Realname  string      `gorm:"column:realname;type:varchar(50)" json:"realname"`                         // 真实姓名
	Password  string      `gorm:"column:password;type:varchar(255)" json:"password"`                        // 密码
	Role      *UserRole   `gorm:"column:role;type:tinyint" json:"role"`                                     // 角色(0:普通用户,1:管理员)
	Status    *UserStatus `gorm:"column:status;type:tinyint" json:"status"`                                 // 状态(0:正常,1:禁用)
	CreatedAt time.Time   `gorm:"column:created_at;type:datetime;autoCreateTime:milli" json:"created_at"`   // 创建时间
	UpdatedAt time.Time   `gorm:"column:updated_at;type:datetime;autoUpdateTime:milli" json:"updated_at"`   // 更新时间
}

func (User) TableName() string {
	return "user"
}

type UserRole int8

const (
	UserRoleNormal UserRole = iota // 普通用户
	UserRoleAdmin                  // 管理员
)

func (r *UserRole) Int8() int8 {
	if r == nil {
		return int8(UserRoleNormal)
	}
	return int8(*r)
}

func (r *UserRole) String() string {
	switch *r {
	case UserRoleNormal:
		return "普通用户"
	case UserRoleAdmin:
		return "管理员"
	default:
		return "未知角色"
	}
}

type UserStatus int8

const (
	UserStatusNormal   UserStatus = iota // 正常
	UserStatusDisabled                   // 禁用
)

func (s *UserStatus) Int8() int8 {
	if s == nil {
		return int8(UserStatusNormal)
	}
	return int8(*s)
}

func (s *UserStatus) String() string {
	switch *s {
	case UserStatusNormal:
		return "正常"
	case UserStatusDisabled:
		return "禁用"
	default:
		return "未知状态"
	}
}
