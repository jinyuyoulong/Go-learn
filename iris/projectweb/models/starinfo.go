package models

type StarInfo struct {
	Id           int    `xorm:"not null pk autoincr comment('主键ID') INT(10)" form:"id"`
	NameZh       string `xorm:"not null comment('中文名') VARCHAR(50)" form:"name_zh"`
	NameEn       string `xorm:"not null comment('英文名') VARCHAR(50)" form:"name_en"`
	Avatar       string `xorm:"not null comment('头像') VARCHAR(255)" form:"avatar"`
	Birthday     string `xorm:"not null comment('出生日期') VARCHAR(50)" form:"birthday"`
	Height       int    `xorm:"not null default 0 comment('身高，单位cm') INT(10)" form:"height"`
	Weight       int    `xorm:"not null comment('体重，单位g') INT(10)" form:"weight"`
	Club         string `xorm:"not null comment('俱乐部') VARCHAR(50)" form:"club"`
	Jersy        string `xorm:"not null comment('球衣号码以及主打位置') VARCHAR(50)" form:"jersy"`
	Country      string `xorm:"not null comment('国籍') VARCHAR(50)" form:"country"`
	Birthaddress string `xorm:"not null comment('出生地') VARCHAR(255)" form:"birthaddress"`
	Feature      string `xorm:"not null comment('个人特点') VARCHAR(255)" form:"feature"`
	Moreinfo     string `xorm:"comment('更多介绍') TEXT" form:"moreinfo"`
	SysStatus    int    `xorm:"not null default 0 comment('状态，默认值 0 正常，1 删除') TINYINT(4)" form:"-"`
	SysCreated   int    `xorm:"not null default 0 comment('创建时间') INT(10)" form:"-"`
	SysUpdated   int    `xorm:"not null default 0 comment('最后修改时间') INT(10)" form:"-"`
}
