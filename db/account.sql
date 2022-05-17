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