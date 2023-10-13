CREATE DATABASE IF NOT EXISTS `momo_user` /*!40100 DEFAULT CHARACTER SET utf8 COLLATE utf8_unicode_ci */;

USE `momo_user`;

DROP TABLE IF EXISTS `momo_user`;


CREATE TABLE `user` (
  `id`          bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `username`    varchar(32)  NOT NULL DEFAULT '' COMMENT '用户名',
  `avatar`      varchar(256) NOT NULL DEFAULT '' COMMENT '头像',
  `mobile`      varchar(128) NOT NULL DEFAULT '' COMMENT '手机号',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_update_time` (`update_time`),
  UNIQUE KEY `uk_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表';


-- 字段 `created_at`