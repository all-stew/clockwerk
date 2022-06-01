drop table if exists `accounts`;
create table `accounts`
(
    `id`            bigint       not null auto_increment comment 'id',
    `account`       varchar(128) not null default '' comment '账号',
    `account_type`  tinyint      not null default 0 comment '账号类型',
    `parent_id`     bigint       not null default 0 comment '父账号id',
    `account_param` text comment '账号参数',
    `user_id`       bigint       not null comment 'user_id',
    `status`        tinyint      not null default 0 comment '状态',
    `created_by`     bigint                default 0 comment '创建者id',
    `created_at`     datetime              default CURRENT_TIMESTAMP comment '创建时间',
    `updated_by`     bigint                default 0 comment '更新者id',
    `updated_at`     datetime              default CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`     datetime              default CURRENT_TIMESTAMP comment '删除时间',
    primary key (`id`)
) engine = innodb
  auto_increment = 100 comment = '账户信息表';

-- ----------------------------
-- 用户信息表
-- ----------------------------
drop table if exists `users`;
create table `users`
(
    `id`              bigint       not null auto_increment comment '用户ID',
    `username`        varchar(64)  not null comment '用户账号',
    `nickname`        varchar(32)  not null comment '用户昵称',
    `email`           varchar(128) not null default '' comment '用户邮箱',
    `phone`           varchar(64)  not null default '' comment '手机号码',
    `status`          tinyint      not null default 0 comment '帐号状态',
    `password`        varchar(256) not null default '' comment '账户密码',
    `created_by`       bigint                default 0 comment '创建者id',
    `created_at`       datetime              default CURRENT_TIMESTAMP comment '创建时间',
    `updated_by`       bigint                default 0 comment '更新者id',
    `updated_at`       datetime              default CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`       datetime              default CURRENT_TIMESTAMP comment '删除时间',
    primary key (`id`)
) engine = innodb
  auto_increment = 100 comment = '用户信息表';

drop table if exists `user_notifications`;
create table `user_notifications`
(
    `id`                 bigint       not null auto_increment comment 'id',
    `notification_type`  tinyint      not null default 0 comment '通知类型',
    `notification_key`   varchar(256) not null default '' comment '通知key',
    `notification_param` text comment '通知参数',
    `user_id`            bigint       not null comment '用户id',
    `created_by`          bigint                default 0 comment '创建者id',
    `created_at`          datetime              default CURRENT_TIMESTAMP comment '创建时间',
    `updated_by`          bigint                default 0 comment '更新者id',
    `updated_at`          datetime              default CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`          datetime              default CURRENT_TIMESTAMP comment '删除时间',
    primary key (`id`)
) engine = innodb
  auto_increment = 100 comment = '用户通知表';