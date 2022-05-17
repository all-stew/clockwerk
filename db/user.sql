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
    `created_by`       bigint                default 0 comment '创建者id',
    `created_at`       datetime              default CURRENT_TIMESTAMP comment '创建时间',
    `updated_by`       bigint                default 0 comment '更新者id',
    `updated_at`       datetime              default CURRENT_TIMESTAMP comment '更新时间',
    `deleted_at`       datetime              default CURRENT_TIMESTAMP comment '删除时间',
    primary key (`id`)
) engine = innodb
  auto_increment = 100 comment = '用户信息表';