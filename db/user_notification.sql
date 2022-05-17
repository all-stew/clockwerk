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