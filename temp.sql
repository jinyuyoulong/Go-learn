
# 数据库建表遵循 全小写原则

create table user(
id INT(10) not null auto_increment comment '主键ID',
username varchar(50) not null comment '用户名',
age tinyint(1) not null default '0' comment '年龄',
email varchar(50) not null comment '邮箱',
pwd varchar(50) not null comment '密码',
primary key (id)
)engine=InnoDB default charset=utf8 comment '用户表';

create table star_info(
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
sys_status tinyint not null default 0 comment '状态，默认值 0 正常，1 删除',
sys_created int(10) not null default 0 comment '创建时间',
sys_updated int(10) not null default 0 comment '最后修改时间',
primary key (id)
)engine=InnoDB default charset=utf8 comment '球星信息';

# Navicat 连接 mysql8 的解决方案有三：
# 方案一 重新创建一个账号，设此账号的加密方式为 mysql_native_password，使用这个账号
create user 'fanjinlong'@'%' identified with mysql_native_password by '333';
create user 'your username'@'%' identified with caching_sha2_password by 'your password';

# 方案二 修改root 的密码加密方式，刷新
ALTER USER 'root'@'localhost' IDENTIFIED BY 'password' PASSWORD EXPIRE NEVER; # 更改root 的加密方式 账户永不过期
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY 'password';# 重置密码
FLUSH PRIVILEGES; #刷新数据库

-- 方案三 更换MySQL8 客户端