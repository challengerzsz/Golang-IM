package model

import "time"

const (
	SEX_WOMEN   = "WOMEN"
	SEX_MAN     = "MAN"
	SEX_UNKNOWN = "UNKNOWN"
)

type User struct {
	// 用户ID
	Id     int64  `xorm:"pk autoincr bigint(20)" form:"id" json:"id"`
	Mobile string `xorm:"varchar(20)" form:"mobile" json:"mobile"`
	Passwd string `xorm:"varchar(40)" form:"passwd" json:"-"`
	// 角色
	Avatar   string `xorm:"varchar(150)" form:"avatar" json:"avatar"`
	Sex      string `xorm:"varchar(2)" form:"sex" json:"sex"`
	Nickname string `xorm:"varchar(20)" form:"nickname" json:"nickname"`
	// 加盐随机字符串6
	Salt   string `xorm:"varchar(10)" form:"salt" json:"-"`
	Online int    `xorm:"int(10)" form:"online" json:"online"`
	// 前端鉴权因子
	Token    string    `xorm:"varchar(40)" form:"token" json:"token"`
	Memo     string    `xorm:"varchar(140)" form:"memo" json:"memo"`
	Createat time.Time `xorm:"datetime" form:"createat" json:"createat"`
}
