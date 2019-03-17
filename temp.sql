create table user(
id INT(10) not null auto_increment comment '主键ID',
username varchar(50) not null comment '用户名',
age tinyint(1) not null default '0' comment '年龄',
email varchar(50) not null comment '邮箱',
pwd varchar(50) not null comment '密码',
primary key (id)
)engine=InnoDB default charset=utf8 comment '用户表';

create table starinfo(
id int(10) not null auto_increment comment '主键ID',
name_zh varchar(50) not null comment '中文名',
name_en varchar(50) not null comment '英文名',
avatar varchar(255) not null comment '头像',
birthday varchar(50) not null comment '出生日期',
height int(10) not null comment '身高，单位cm',
weight int(10) not null comment '体重，单位g',
club varchar(50) not null comment '俱乐部',
jersy varchar(50) not null comment '球衣号码以及主打位置',
country varchar(50) not null comment '国籍',
birthaddress varchar(255) not null comment '出生地',
feature varchar(255) not null comment '个人特点',
moreinfo TEXT not null comment '更多介绍',
SysStatus tinyint not null default 0 comment '状态，默认值 0 正常，1 删除',
SysCreated int(10) not null default 0 comment '创建时间',
SysUpdated int(10) not null default 0 comment '最后修改时间',
primary key (id)
)engine=InnoDB default charset=utf8 comment '球星信息';