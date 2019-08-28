create database if not EXISTS db_doublejwt DEFAULT CHARSET=utf8;

use db_doublejwt;

CREATE TABLE `tb_user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '添加时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `uuid` varchar(50) NOT NULL COMMENT '用户uuid',
  `username` char(50) NOT NULL COMMENT '用户帐号',
  `password` char(100) NOT NULL COMMENT '用户密码',
  PRIMARY KEY (`id`),
  KEY `idx_tb_user_updated_at` (`updated_at`),
  KEY `idx_uuid` (`uuid`),
  KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户表';


CREATE TABLE `tb_comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL COMMENT '添加时间',
  `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
  `topic_id` int(10) unsigned NOT NULL COMMENT '主题id',
  `topic_type` varchar(50) NOT NULL COMMENT '主题类型',
  `content` text NOT NULL COMMENT '评论内容',
  `from_uid` int(10) unsigned NOT NULL COMMENT '评论用户id',
  PRIMARY KEY (`id`),
  KEY `idx_tb_comment_updated_at` (`updated_at`),
  KEY `idx_topic_id` (`topic_id`),
  KEY `idx_from_uid` (`from_uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论表';